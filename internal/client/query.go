package client

import (
	"net/url"
	"strings"
)

// QueryOption customises a List request.
type QueryOption func(*Query)

// Query encodes the filter / projection options for /rest GETs.
//
// Print options (per REST API docs):
//
//	?<prop>=<value>            equality filter
//	?.proplist=a,b,c           projection (return only these properties)
//	?.query=k1=v1&...&#|       complex queries (POST .../print form). Note that
//	                           query stack operators (#|, #&, #!) are typically
//	                           used via POST .../print {".query":[...]}, not GET.
type Query struct {
	Filters  map[string]string
	Proplist []string
}

func NewQuery(opts ...QueryOption) *Query {
	q := &Query{Filters: map[string]string{}}
	for _, o := range opts {
		o(q)
	}
	return q
}

// WithFilter adds an equality filter (?k=v).
func WithFilter(k, v string) QueryOption {
	return func(q *Query) { q.Filters[k] = v }
}

// WithProplist restricts the returned columns to those listed.
func WithProplist(props ...string) QueryOption {
	return func(q *Query) { q.Proplist = append(q.Proplist, props...) }
}

func (q *Query) Encode() string {
	if q == nil {
		return ""
	}
	vals := url.Values{}
	for k, v := range q.Filters {
		vals.Set(k, v)
	}
	if len(q.Proplist) > 0 {
		vals.Set(".proplist", strings.Join(q.Proplist, ","))
	}
	return vals.Encode()
}

// PrintBody constructs the body for POST /rest/<path>/print, supporting the
// full query-stack language (operators #|, #&, #!).
type PrintBody struct {
	Proplist []string `json:".proplist,omitempty"`
	Query    []string `json:".query,omitempty"`
}
