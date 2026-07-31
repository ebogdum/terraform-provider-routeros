package provider

import (
	"testing"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

func serviceRows() []client.Object {
	return []client.Object{
		{".id": "*1", "name": "telnet", "disabled": "true", "port": "23"},
		{".id": "*2", "name": "ftp", "disabled": "true", "port": "21"},
		{".id": "*3", "name": "www", "disabled": "true", "port": "80"},
		{".id": "*4", "name": "ssh", "disabled": "false", "port": "22"},
		{".id": "*5", "name": "api", "disabled": "true", "port": "8728"},
		{".id": "*6", "name": "winbox", "disabled": "true", "port": "8291"},
	}
}

func TestIPServiceFindByName(t *testing.T) {
	rows := serviceRows()
	if row := ipServiceFindByName(rows, "api"); row == nil || row[".id"] != "*5" {
		t.Fatalf("api lookup = %v, want row *5", row)
	}
	// RouterOS names are lowercase; accept any casing from config.
	if row := ipServiceFindByName(rows, "Telnet"); row == nil || row[".id"] != "*1" {
		t.Fatalf("Telnet lookup = %v, want row *1", row)
	}
	if row := ipServiceFindByName(rows, "nosuch"); row != nil {
		t.Fatalf("nosuch lookup = %v, want nil", row)
	}
}

func TestIPServiceNames(t *testing.T) {
	got := ipServiceNames(serviceRows())
	want := []string{"api", "ftp", "ssh", "telnet", "winbox", "www"}
	if len(got) != len(want) {
		t.Fatalf("names = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("names = %v, want %v", got, want)
		}
	}
}

func TestIPServiceCheckLockoutAllowsDisablingOneOfMany(t *testing.T) {
	// api is already disabled; disabling telnet still leaves ssh enabled.
	err := ipServiceCheckLockout(serviceRows(), "*1", client.Object{"disabled": "true"}, false)
	if err != nil {
		t.Fatalf("disabling telnet rejected: %v", err)
	}
}

func TestIPServiceCheckLockoutRefusesLastService(t *testing.T) {
	// ssh (*4) is the only enabled management service; turning it off locks us out.
	err := ipServiceCheckLockout(serviceRows(), "*4", client.Object{"disabled": "true"}, false)
	if err == nil {
		t.Fatal("disabling the last enabled service was allowed, want refusal")
	}
}

func TestIPServiceCheckLockoutHonoursAck(t *testing.T) {
	err := ipServiceCheckLockout(serviceRows(), "*4", client.Object{"disabled": "true"}, true)
	if err != nil {
		t.Fatalf("lockout_ack did not override the guard: %v", err)
	}
}

func TestIPServiceCheckLockoutDoesNotMutateInput(t *testing.T) {
	rows := serviceRows()
	if err := ipServiceCheckLockout(rows, "*4", client.Object{"disabled": "true"}, true); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if rows[3]["disabled"] != "false" {
		t.Fatalf("projection mutated the caller's rows: %v", rows[3])
	}
}

func TestIPServiceCheckLockoutAllowsReEnabling(t *testing.T) {
	// Enabling a service can never lock anyone out, even from an all-off state.
	allOff := serviceRows()
	allOff[3]["disabled"] = "true"
	if err := ipServiceCheckLockout(allOff, "*6", client.Object{"disabled": "false"}, false); err != nil {
		t.Fatalf("enabling winbox rejected: %v", err)
	}
}
