// Package schemautil -- lockout_guards.go
//
// Additional lockout guards beyond firewall (see lockout_guard.go).
// Each guard refuses an operation that would obviously sever management
// access unless the resource was applied with `lockout_ack = true`.
package schemautil

import (
	"errors"
	"net"
	"strings"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

// CheckWatchdogLockout refuses a /system/watchdog watch-address that can never
// be a valid ping target -- documentation (TEST-NET), loopback, link-local,
// unspecified, multicast or reserved space. RouterOS's ping watchdog reboots
// the board when the address is unreachable, so a bad watch-address
// reboot-loops the device off the network with no Terraform-side error. The
// caller passes the intended body; set acknowledged (lockout_ack=true) to
// override.
func CheckWatchdogLockout(body client.Object, acknowledged bool) error {
	if acknowledged {
		return nil
	}
	addr := strings.TrimSpace(body["watch-address"])
	if addr == "" || strings.EqualFold(addr, "none") {
		return nil
	}
	ip := net.ParseIP(addr)
	if ip == nil {
		return nil // not an IP literal (e.g. a name); nothing to judge here
	}
	unroutable := ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() ||
		ip.IsMulticast() || ip.IsUnspecified()
	// RFC 5737 documentation ranges + RFC 6598/reserved that never answer pings.
	for _, cidr := range []string{"192.0.2.0/24", "198.51.100.0/24", "203.0.113.0/24", "240.0.0.0/4"} {
		_, n, _ := net.ParseCIDR(cidr)
		if n != nil && n.Contains(ip) {
			unroutable = true
			break
		}
	}
	if !unroutable {
		return nil
	}
	return errors.New("refusing /system/watchdog watch-address=" + addr +
		": this address cannot answer pings, so RouterOS's ping watchdog will " +
		"reboot-loop the device off the network. Use a reachable address, `none`, " +
		"or set lockout_ack=true to override")
}

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
//
// RouterOS does not permit add/remove on /ip/service rows, only enable/disable
// and reconfiguration. The routeros_ip_service resource adopts a row by name
// and calls this guard (via ipServiceCheckLockout) before every write.
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
		"service (winbox/ssh/api/www/...); leave at least one enabled or set " +
		"lockout_ack=true to override")
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
	raw, ok := body["allowed-interface-list"]
	if !ok {
		return nil // field not part of this write; nothing to guard
	}
	allowed := strings.TrimSpace(raw)
	if allowed == "" || allowed == "none" {
		return errors.New("refusing to set /tool/mac-server allowed-interface-list to empty/none: " +
			"this disables MAC-Winbox recovery, leaving no out-of-band access if SSH/Winbox/web break; " +
			"set lockout_ack=true to override")
	}
	return nil
}
