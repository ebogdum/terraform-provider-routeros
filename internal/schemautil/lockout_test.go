package schemautil

import (
	"strings"
	"testing"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

func TestCheckFirewallLockout(t *testing.T) {
	cases := []struct {
		name     string
		menu     string
		body     client.Object
		ack      bool
		wantErr  bool
		wantSubs string
	}{
		{
			name:    "input/drop with no match -> reject",
			menu:    "/ip/firewall/filter",
			body:    client.Object{"chain": "input", "action": "drop"},
			wantErr: true, wantSubs: "input/drop",
		},
		{
			name: "input/drop with src-address -> allow",
			menu: "/ip/firewall/filter",
			body: client.Object{"chain": "input", "action": "drop", "src-address": "10.0.0.0/8"},
		},
		{
			name:    "forward/reject with no match -> reject",
			menu:    "/ipv6/firewall/filter",
			body:    client.Object{"chain": "forward", "action": "reject"},
			wantErr: true, wantSubs: "forward/reject",
		},
		{
			name: "input/accept -> allow",
			menu: "/ip/firewall/filter",
			body: client.Object{"chain": "input", "action": "accept"},
		},
		{
			name:    "input/drop with whitespace-only match -> reject",
			menu:    "/ip/firewall/filter",
			body:    client.Object{"chain": "input", "action": "drop", "src-address": "   "},
			wantErr: true, wantSubs: "input/drop",
		},
		{
			name: "ack=true bypasses",
			menu: "/ip/firewall/filter",
			body: client.Object{"chain": "input", "action": "drop"},
			ack:  true,
		},
		{
			name: "non-firewall menu -> allow",
			menu: "/ip/address",
			body: client.Object{"chain": "input", "action": "drop"},
		},
		{
			name: "output chain -> allow (not guarded)",
			menu: "/ip/firewall/filter",
			body: client.Object{"chain": "output", "action": "drop"},
		},
		{
			name:    "uppercase chain/action still detected",
			menu:    "/ip/firewall/filter",
			body:    client.Object{"chain": "INPUT", "action": "Drop"},
			wantErr: true, wantSubs: "input/drop",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := CheckFirewallLockout(tc.menu, tc.body, tc.ack)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("want error, got nil")
				}
				if tc.wantSubs != "" && !strings.Contains(err.Error(), tc.wantSubs) {
					t.Fatalf("error %q does not contain %q", err.Error(), tc.wantSubs)
				}
				return
			}
			if err != nil {
				t.Fatalf("want nil, got %v", err)
			}
		})
	}
}

func TestCheckUserDeleteLockout(t *testing.T) {
	cases := []struct {
		name    string
		menu    string
		body    client.Object
		op      string
		ack     bool
		wantErr bool
	}{
		{
			name:    "delete admin -> reject",
			menu:    "/user",
			body:    client.Object{"name": "admin", "group": "full"},
			op:      "delete",
			wantErr: true,
		},
		{
			name:    "delete full-group user -> reject",
			menu:    "/user",
			body:    client.Object{"name": "alice", "group": "full"},
			op:      "delete",
			wantErr: true,
		},
		{
			name: "delete non-admin -> allow",
			menu: "/user",
			body: client.Object{"name": "bob", "group": "read"},
			op:   "delete",
		},
		{
			name:    "disable admin -> reject",
			menu:    "/user",
			body:    client.Object{"name": "admin", "group": "full", "disabled": "false"},
			op:      "disable",
			wantErr: true,
		},
		{
			name: "disable already-disabled admin -> allow",
			menu: "/user",
			body: client.Object{"name": "admin", "group": "full", "disabled": "true"},
			op:   "disable",
		},
		{
			name: "ack bypasses",
			menu: "/user",
			body: client.Object{"name": "admin", "group": "full"},
			op:   "delete",
			ack:  true,
		},
		{
			name: "non-user menu -> allow",
			menu: "/ip/address",
			body: client.Object{"name": "admin", "group": "full"},
			op:   "delete",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := CheckUserDeleteLockout(tc.menu, tc.body, tc.op, tc.ack)
			if tc.wantErr && err == nil {
				t.Fatalf("want error, got nil")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("want nil, got %v", err)
			}
		})
	}
}

func TestCheckUserGroupPolicyLockout(t *testing.T) {
	if err := CheckUserGroupPolicyLockout("/user/group", client.Object{
		"name": "full", "policy": "api,web,winbox,ssh,telnet,password,policy,write,read,reboot",
	}, false); err != nil {
		t.Fatalf("complete policy should pass: %v", err)
	}
	err := CheckUserGroupPolicyLockout("/user/group", client.Object{
		"name": "full", "policy": "api,web",
	}, false)
	if err == nil || !strings.Contains(err.Error(), "missing policies") {
		t.Fatalf("want missing-policies error, got %v", err)
	}
	if err := CheckUserGroupPolicyLockout("/user/group", client.Object{"name": "read", "policy": ""}, false); err != nil {
		t.Fatalf("non-full group should pass: %v", err)
	}
	if err := CheckUserGroupPolicyLockout("/user/group", client.Object{"name": "full", "policy": "api"}, true); err != nil {
		t.Fatalf("ack should bypass: %v", err)
	}
}

func TestCheckIPServiceLockout(t *testing.T) {
	allDisabled := []client.Object{
		{"name": "ssh", "disabled": "true"},
		{"name": "winbox", "disabled": "true"},
		{"name": "api", "disabled": "true"},
		{"name": "www", "disabled": "true"},
	}
	if err := CheckIPServiceLockout("/ip/service", allDisabled, false); err == nil {
		t.Fatalf("want lockout error when every service is disabled")
	}
	withOneEnabled := append([]client.Object{}, allDisabled...)
	withOneEnabled[0] = client.Object{"name": "ssh", "disabled": "false"}
	if err := CheckIPServiceLockout("/ip/service", withOneEnabled, false); err != nil {
		t.Fatalf("ssh enabled should pass: %v", err)
	}
	if err := CheckIPServiceLockout("/ip/service", allDisabled, true); err != nil {
		t.Fatalf("ack should bypass: %v", err)
	}
	if err := CheckIPServiceLockout("/other", allDisabled, false); err != nil {
		t.Fatalf("non-service menu should pass: %v", err)
	}
}

func TestCheckMACServerLockout(t *testing.T) {
	for _, p := range []string{"/tool/mac-server", "/tool/mac-server/mac-winbox", "/tool/mac-server/mac-winbox-server"} {
		if err := CheckMACServerLockout(p, client.Object{"allowed-interface-list": "all"}, false); err != nil {
			t.Fatalf("%s with allowed-interface-list=all should pass: %v", p, err)
		}
		if err := CheckMACServerLockout(p, client.Object{"allowed-interface-list": ""}, false); err == nil {
			t.Fatalf("%s with empty allowed-interface-list should reject", p)
		}
		if err := CheckMACServerLockout(p, client.Object{"allowed-interface-list": "none"}, false); err == nil {
			t.Fatalf("%s with none allowed-interface-list should reject", p)
		}
		if err := CheckMACServerLockout(p, client.Object{"allowed-interface-list": ""}, true); err != nil {
			t.Fatalf("%s ack should bypass: %v", p, err)
		}
		// Field absent from the body means the write does not touch
		// allowed-interface-list, so the guard must not fire (an empty
		// terraform config must not read as "set it to empty").
		if err := CheckMACServerLockout(p, client.Object{}, false); err != nil {
			t.Fatalf("%s with field absent should pass: %v", p, err)
		}
	}
	if err := CheckMACServerLockout("/ip/address", client.Object{"allowed-interface-list": ""}, false); err != nil {
		t.Fatalf("non-mac-server menu should pass: %v", err)
	}
}
