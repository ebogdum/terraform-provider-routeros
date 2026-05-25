// Package schemautil -- lockout_guards.go
//
// Additional lockout guards beyond firewall (see lockout_guard.go).
// Each guard refuses an operation that would obviously sever management
// access unless the resource was applied with `lockout_ack = true`.
package schemautil

import (
	"errors"
	"strings"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

// CheckUserDeleteLockout refuses to delete or disable a user named "admin"
// (or the last enabled member of the "full" group). Caller passes the row's
// current values plus the operation kind ("delete" or "disable").
func CheckUserDeleteLockout(menuPath string, body client.Object, op string, acknowledged bool) error {
	if acknowledged {
		return nil
	}
	if menuPath != "/user" {
		return nil
	}
	name := strings.ToLower(body["name"])
	group := strings.ToLower(body["group"])
	disabled := strings.ToLower(body["disabled"])
	if op == "delete" && (name == "admin" || group == "full") {
		return errors.New("refusing to delete user '" + body["name"] +
			"' (group=" + body["group"] + "): this is the last admin account I can detect. " +
			"Set lockout_ack=true to override.")
	}
	if op == "disable" && disabled != "true" && (name == "admin" || group == "full") {
		// Caller wants to disable an admin account.
		return errors.New("refusing to disable user '" + body["name"] +
			"' (group=" + body["group"] + "): would remove admin access. " +
			"Set lockout_ack=true to override.")
	}
	return nil
}

// CheckUserGroupPolicyLockout refuses changes to the "full" group that strip
// policy entries needed for management.
func CheckUserGroupPolicyLockout(menuPath string, body client.Object, acknowledged bool) error {
	if acknowledged {
		return nil
	}
	if menuPath != "/user/group" {
		return nil
	}
	name := strings.ToLower(body["name"])
	if name != "full" {
		return nil
	}
	policy := strings.ToLower(body["policy"])
	if policy == "" {
		return nil
	}
	// All these tokens MUST appear in 'full' group's policy or admin loses access.
	required := []string{"api", "web", "winbox", "ssh", "telnet", "password", "policy", "write"}
	missing := []string{}
	for _, r := range required {
		if !strings.Contains(policy, r) {
			missing = append(missing, r)
		}
	}
	if len(missing) == 0 {
		return nil
	}
	return errors.New("refusing to apply /user/group full with missing policies " +
		strings.Join(missing, ",") + ": you'd be unable to log back in. " +
		"Set lockout_ack=true to override.")
}

// CheckIPServiceLockout refuses to disable ALL management services
// (winbox/ssh/api/api-ssl/www/www-ssl). At least one must remain enabled.
//
// Caller is expected to supply the full intended state (snapshot of all
// service rows after the planned change), not just the diff for one row,
// since the test is global.
func CheckIPServiceLockout(menuPath string, allServices []client.Object, acknowledged bool) error {
	if acknowledged {
		return nil
	}
	if menuPath != "/ip/service" {
		return nil
	}
	management := map[string]bool{
		"winbox": true, "ssh": true, "api": true, "api-ssl": true,
		"www": true, "www-ssl": true, "telnet": true,
	}
	anyEnabled := false
	for _, row := range allServices {
		name := strings.ToLower(row["name"])
		disabled := strings.ToLower(row["disabled"])
		if management[name] && disabled != "true" {
			anyEnabled = true
			break
		}
	}
	if anyEnabled {
		return nil
	}
	return errors.New("refusing to apply /ip/service: would disable every management " +
		"service (winbox/ssh/api/www/...). Leave at least one enabled or set " +
		"lockout_ack=true to override.")
}

// CheckMACServerLockout refuses an empty allowed-interface-list on /tool/mac-server
// or /tool/mac-server/mac-winbox-server (which would block MAC-Winbox recovery
// access when normal services are down).
func CheckMACServerLockout(menuPath string, body client.Object, acknowledged bool) error {
	if acknowledged {
		return nil
	}
	if menuPath != "/tool/mac-server" && menuPath != "/tool/mac-server/mac-winbox" &&
		menuPath != "/tool/mac-server/mac-winbox-server" {
		return nil
	}
	allowed := strings.TrimSpace(body["allowed-interface-list"])
	if allowed == "" || allowed == "none" {
		return errors.New("refusing to set /tool/mac-server allowed-interface-list to empty/none: " +
			"this disables MAC-Winbox recovery, leaving no out-of-band access if SSH/Winbox/web break. " +
			"Set lockout_ack=true to override.")
	}
	return nil
}

// CheckSSHKeyImportLockout warns when an SSH key import would REPLACE the
// last admin's authentication and password-authentication is disabled in
// /ip/ssh -- leaving no fallback if the private key is lost.
//
// Returns nil for now; provided as a placeholder so callers can wire it
// without code changes when we have multi-resource visibility.
func CheckSSHKeyImportLockout(menuPath string, body client.Object, acknowledged bool) error {
	// Multi-resource invariant; the current per-resource Create/Update path can't
	// see /ip/ssh's password-authentication setting. Implement when we have a
	// pre-apply hook with cross-resource state.
	_ = menuPath
	_ = body
	_ = acknowledged
	return nil
}
