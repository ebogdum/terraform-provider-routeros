package client

import (
	"context"
	"sort"
	"sync"
)

// orderLocks serialises PlaceOrdered per (router-name, menuPath) within a
// single client. Keyed locally rather than globally so two clients pointing
// at the same base URL but configured separately do not share a lock.
func (c *Client) lockFor(router, menuPath string) *sync.Mutex {
	c.orderLocksMu.Lock()
	defer c.orderLocksMu.Unlock()
	if c.orderLocks == nil {
		c.orderLocks = map[string]*sync.Mutex{}
	}
	k := router + "|" + menuPath
	mu, ok := c.orderLocks[k]
	if !ok {
		mu = &sync.Mutex{}
		c.orderLocks[k] = mu
	}
	return mu
}

// PlaceOrdered moves thisID into the correct slot within menuPath relative to
// other Terraform-managed rules whose positions are recorded in the
// Registry's in-memory ordered-resources map.
//
// Algorithm:
//
//  1. Read the current chain order from the device (List returns rules in
//     applied order).
//  2. Intersect that with the snapshot of (id -> position) the Registry holds
//     for THIS (router, menuPath) -- only Terraform-managed rules count.
//  3. Sort managed rules by position; find the first managed rule whose
//     position is strictly greater than thisPosition.
//  4. /move thisID before that rule. If none exists, thisID stays at the end
//     of the managed section (RouterOS appends new rules naturally).
//
// Nothing is written into rule comments. Position lives only in TF state +
// the in-memory registry map.
func (c *Client) PlaceOrdered(ctx context.Context, router, menuPath, thisID string, thisPosition int64, snapshot map[string]int64) error {
	mu := c.lockFor(router, menuPath)
	mu.Lock()
	defer mu.Unlock()

	if snapshot == nil {
		snapshot = map[string]int64{}
	}
	// Ensure thisID is in the snapshot at thisPosition (caller may have
	// just registered it; defensive).
	snapshot[thisID] = thisPosition

	rows, err := c.List(ctx, menuPath)
	if err != nil {
		return err
	}
	// Build managed-rules list in CURRENT device order. RouterOS returns rows
	// in chain order top-to-bottom.
	type peer struct {
		id  string
		pos int64
	}
	var managedInDeviceOrder []peer
	for _, r := range rows {
		id := r[".id"]
		if pos, ok := snapshot[id]; ok {
			managedInDeviceOrder = append(managedInDeviceOrder, peer{id: id, pos: pos})
		}
	}
	// Compute desired order: sort managed by position (stable on equal positions).
	desired := make([]peer, len(managedInDeviceOrder))
	copy(desired, managedInDeviceOrder)
	sort.SliceStable(desired, func(i, j int) bool { return desired[i].pos < desired[j].pos })

	// A lone managed rule has no peer to sort against.
	if len(desired) < 2 {
		return nil
	}

	// Find what should come AFTER thisID in the desired order.
	var destID, prevID string
	for i, p := range desired {
		if p.id != thisID {
			continue
		}
		if i > 0 {
			prevID = desired[i-1].id
		}
		if i+1 < len(desired) {
			destID = desired[i+1].id
		}
		break
	}

	// Skip the move if device order already has thisID directly before destID.
	for i, p := range managedInDeviceOrder {
		if p.id != thisID {
			continue
		}
		var have string
		if i+1 < len(managedInDeviceOrder) {
			have = managedInDeviceOrder[i+1].id
		}
		if have == destID {
			return nil
		}
		break
	}
	if destID != "" {
		return c.Move(ctx, menuPath, thisID, destID)
	}
	return c.moveAfter(ctx, menuPath, rows, thisID, prevID)
}

// moveAfter places thisID directly after anchorID, so a rule sorting last among
// managed rules stays put instead of jumping every unmanaged rule below it.
func (c *Client) moveAfter(ctx context.Context, menuPath string, rows []Object, thisID, anchorID string) error {
	anchor := -1
	for i, r := range rows {
		if r[".id"] == anchorID {
			anchor = i
			break
		}
	}
	if anchor < 0 {
		return nil
	}
	for _, r := range rows[anchor+1:] {
		if id := r[".id"]; id != thisID {
			return c.Move(ctx, menuPath, thisID, id)
		}
	}
	return c.MoveToEnd(ctx, menuPath, thisID)
}
