package client

import (
	"errors"
	"fmt"
	"sort"
	"sync"
)

// Registry holds one *Client per named router and is the unit handed to every
// resource via ProviderData. Resources select which router to act against via
// their `router` attribute; an empty value selects the default router.
type Registry struct {
	mu       sync.RWMutex
	clients  map[string]*Client
	def      string // name of the default router (first declared, or "default")
	// ordered is a per-(router, menuPath, .id) -> desired position map populated
	// by ordered-resource Create/Update and consumed by Client.PlaceOrdered.
	// Lives only in-memory; ground truth is Terraform state, not the device.
	ordered map[orderedKey]map[string]int64
}

type orderedKey struct {
	Router string
	Menu   string
}

// NewRegistry builds a Registry from a map of name -> Config. The first name
// in alphabetical order becomes the default unless one is explicitly named
// "default". An empty input is rejected -- at least one router must be defined.
func NewRegistry(configs map[string]Config) (*Registry, error) {
	if len(configs) == 0 {
		return nil, errors.New("routeros: no routers configured")
	}
	reg := &Registry{
		clients: make(map[string]*Client, len(configs)),
		ordered: make(map[orderedKey]map[string]int64),
	}
	names := make([]string, 0, len(configs))
	for n := range configs {
		names = append(names, n)
	}
	sort.Strings(names)
	for _, n := range names {
		c, err := New(configs[n])
		if err != nil {
			return nil, fmt.Errorf("routeros: router %q: %w", n, err)
		}
		reg.clients[n] = c
	}
	if _, ok := configs["default"]; ok {
		reg.def = "default"
	} else {
		reg.def = names[0]
	}
	return reg, nil
}

// Get returns the named client. Empty name returns the default.
func (r *Registry) Get(name string) (*Client, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if name == "" {
		name = r.def
	}
	c, ok := r.clients[name]
	if !ok {
		return nil, fmt.Errorf("routeros: no router named %q (known: %v)", name, r.names())
	}
	return c, nil
}

// Default returns the default router's client.
func (r *Registry) Default() *Client {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.clients[r.def]
}

// Names returns every configured router name in sorted order.
func (r *Registry) Names() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.names()
}

// RegisterOrdered records that the given .id on (router, menuPath) is managed
// by Terraform and wants the given position in the chain. Called from Create
// and Update after the Add/Set returns.
func (r *Registry) RegisterOrdered(router, menuPath, id string, position int64) {
	r.mu.Lock()
	defer r.mu.Unlock()
	k := orderedKey{router, menuPath}
	m := r.ordered[k]
	if m == nil {
		m = map[string]int64{}
		r.ordered[k] = m
	}
	m[id] = position
}

// UnregisterOrdered removes an .id from the ordered-resources map. Called
// from Delete.
func (r *Registry) UnregisterOrdered(router, menuPath, id string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if m := r.ordered[orderedKey{router, menuPath}]; m != nil {
		delete(m, id)
	}
}

// OrderedSnapshot returns a copy of the (id -> position) map for one
// (router, menuPath). Used by PlaceOrdered.
func (r *Registry) OrderedSnapshot(router, menuPath string) map[string]int64 {
	r.mu.RLock()
	defer r.mu.RUnlock()
	src := r.ordered[orderedKey{router, menuPath}]
	out := make(map[string]int64, len(src))
	for k, v := range src {
		out[k] = v
	}
	return out
}

func (r *Registry) names() []string {
	out := make([]string, 0, len(r.clients))
	for n := range r.clients {
		out = append(out, n)
	}
	sort.Strings(out)
	return out
}
