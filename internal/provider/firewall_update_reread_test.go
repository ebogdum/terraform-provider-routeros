package provider

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

// fakeMenu answers GET/PATCH for one rule and counts what was asked of it.
type fakeMenu struct {
	rule    client.Object
	gets    int
	patches int
	gone    bool
}

func (f *fakeMenu) client(t *testing.T) *client.Client {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodPatch:
			f.patches++
			_ = json.NewEncoder(w).Encode(f.rule)
		case strings.HasSuffix(r.URL.Path, "/move"):
			_, _ = w.Write([]byte(`[]`))
		default:
			f.gets++
			if f.gone {
				w.WriteHeader(http.StatusNotFound)
				_, _ = w.Write([]byte(`{"error":404,"message":"no such item"}`))
				return
			}
			_ = json.NewEncoder(w).Encode(f.rule)
		}
	}))
	t.Cleanup(srv.Close)
	c, err := client.New(client.Config{Host: srv.URL, Username: "u", Password: "p"})
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	return c
}

func rereadInto(t *testing.T, c *client.Client, plan *IPFirewallFilterModel) error {
	t.Helper()
	obj, err := c.GetByID(context.Background(), "/ip/firewall/filter", plan.ID.ValueString())
	if err != nil {
		return err
	}
	iPFirewallFilterApply(context.Background(), obj, plan)
	nullifyUnknownAttrs(plan)
	return nil
}

// A position-only change builds an empty body, so no Set runs and nothing ever
// populates the Computed attributes the plan left Unknown.
func TestFirewallUpdateRereadPopulatesComputedAttrs(t *testing.T) {
	f := &fakeMenu{rule: client.Object{
		".id": "*A", "chain": "input", "action": "accept", "bytes": "17",
	}}
	c := f.client(t)

	plan := &IPFirewallFilterModel{
		ID:      types.StringValue("*A"),
		Chain:   types.StringUnknown(),
		Action:  types.StringUnknown(),
		Bytes:   types.StringUnknown(),
		Comment: types.StringUnknown(),
		Time:    csvSetValue{StringValue: types.StringUnknown()},
	}

	if err := rereadInto(t, c, plan); err != nil {
		t.Fatalf("re-read: %v", err)
	}
	if f.gets != 1 {
		t.Errorf("gets = %d, want 1", f.gets)
	}
	for name, v := range map[string]interface{ IsUnknown() bool }{
		"chain": plan.Chain, "action": plan.Action, "time": plan.Time,
		"bytes": plan.Bytes, "comment": plan.Comment,
	} {
		if v.IsUnknown() {
			t.Errorf("%s is still unknown after the re-read", name)
		}
	}
	if plan.Chain.ValueString() != "input" {
		t.Errorf("chain = %q, want input", plan.Chain.ValueString())
	}
	if plan.Bytes.ValueString() != "17" {
		t.Errorf("bytes = %q, want 17", plan.Bytes.ValueString())
	}
	if !plan.Comment.IsNull() {
		t.Errorf("comment = %v, want null: the device omitted the key", plan.Comment)
	}
}

func TestFirewallUpdateRereadReportsADeletedRuleAsGone(t *testing.T) {
	f := &fakeMenu{rule: client.Object{".id": "*A"}, gone: true}
	c := f.client(t)

	plan := &IPFirewallFilterModel{ID: types.StringValue("*A")}
	err := rereadInto(t, c, plan)
	if err == nil {
		t.Fatal("re-read of a deleted rule returned no error")
	}
	if !client.IsNotFound(err) {
		t.Errorf("err = %v, want a not-found the caller turns into RemoveResource", err)
	}
}
