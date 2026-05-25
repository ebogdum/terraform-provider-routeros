package client

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// placeOrderedLocks serialises PlaceOrdered per (router-base, menuPath).
// Terraform applies resources in parallel by default; without this lock two
// concurrent Creates on the same chain would each list a stale snapshot,
// compute a stale destination, and Move into each other's path. Locking
// across the read+move makes the operation atomic from the chain's POV.
var (
	placeOrderedMu sync.Mutex
	placeOrderedLocks = map[string]*sync.Mutex{}
)

func (c *Client) lockFor(menuPath string) *sync.Mutex {
	key := c.base.Host + "::" + menuPath
	placeOrderedMu.Lock()
	defer placeOrderedMu.Unlock()
	m, ok := placeOrderedLocks[key]
	if !ok {
		m = &sync.Mutex{}
		placeOrderedLocks[key] = m
	}
	return m
}

// Order-marker support for per-rule ordered menus (firewall filter/nat/mangle/
// raw, IPv6 firewall, queue tree).
//
// Design:
//   - The user-facing `position` attribute is an integer sort key. Smaller
//     numbers mean "higher in the chain", same as iptables line numbers.
//   - The provider persists position on the device by prefixing the rule's
//     comment with "[tf:pos=N] ". On read it strips the prefix before
//     handing the comment back to Terraform state.
//   - A `set` label (defaulting to "default") scopes which rules belong
//     together: the prefix is actually "[tf:set=NAME pos=N]". A user who
//     wants two independent ordered sets in the same chain uses two `set`
//     names.
//   - At Create/Update time the provider lists the menu, filters to peers
//     sharing the same `set`, sorts by extracted position, and calls Move
//     to slot this rule at the correct spot.

const (
	defaultSet = "default"
)

var orderMarkerRe = regexp.MustCompile(`^\[tf:(?:set=([A-Za-z0-9_-]+)\s+)?pos=(\d+)\]\s*`)

// EncodeOrderedComment prepends the position marker to userComment. If set is
// empty, "default" is used.
func EncodeOrderedComment(set string, position int64, userComment string) string {
	if set == "" {
		set = defaultSet
	}
	comment := fmt.Sprintf("[tf:set=%s pos=%d]", set, position); if userComment != "" { comment += " " + userComment }; return comment
}

// DecodeOrderedComment splits comment into (set, position, userComment).
// Returns ("", 0, comment, false) if no marker is present.
func DecodeOrderedComment(comment string) (set string, position int64, user string, has bool) {
	m := orderMarkerRe.FindStringSubmatch(comment)
	if m == nil {
		return "", 0, comment, false
	}
	set = m[1]
	if set == "" {
		set = defaultSet
	}
	position, _ = strconv.ParseInt(m[2], 10, 64)
	user = strings.TrimSpace(comment[len(m[0]):])
	return set, position, user, true
}

// PlaceOrdered scans menuPath for rules sharing set, sorts them by position,
// and Moves thisID to the correct slot. RouterOS doesn't expose a "move to
// absolute index" -- only "move before destination .id" -- so PlaceOrdered
// finds the next-higher peer and moves thisID just before it. If thisID has
// the highest position in the set, it's moved to the end (no destination).
func (c *Client) PlaceOrdered(ctx context.Context, menuPath, thisID, set string, thisPosition int64) error {
	if set == "" {
		set = defaultSet
	}
	mu := c.lockFor(menuPath)
	mu.Lock()
	defer mu.Unlock()
	rows, err := c.List(ctx, menuPath)
	if err != nil {
		return err
	}
	type peer struct {
		id  string
		pos int64
	}
	var peers []peer
	for _, r := range rows {
		s, pos, _, has := DecodeOrderedComment(r["comment"])
		if !has || s != set {
			continue
		}
		peers = append(peers, peer{id: r[".id"], pos: pos})
	}
	sort.Slice(peers, func(i, j int) bool { return peers[i].pos < peers[j].pos })

	// Find the next peer (in sorted-by-position order) with pos strictly
	// greater than this one. thisID should land just before it. If nothing
	// is higher, thisID belongs at the chain's tail (destID = "").
	var destID string
	for _, p := range peers {
		if p.id == thisID {
			continue
		}
		if p.pos > thisPosition {
			destID = p.id
			break
		}
	}
	// Skip the Move only when the DEVICE order (top->bottom as RouterOS
	// returns it) already has thisID directly before destID -- or thisID
	// at the tail and destID empty.
	deviceOrder := make([]string, 0, len(rows))
	for _, r := range rows {
		s, _, _, has := DecodeOrderedComment(r["comment"])
		if !has || s != set {
			continue
		}
		deviceOrder = append(deviceOrder, r[".id"])
	}
	for i, id := range deviceOrder {
		if id != thisID {
			continue
		}
		var have string
		if i+1 < len(deviceOrder) {
			have = deviceOrder[i+1]
		}
		if have == destID {
			return nil
		}
		break
	}
	return c.Move(ctx, menuPath, thisID, destID)
}
