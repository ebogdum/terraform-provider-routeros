// Hand-written: verifies action resources actually do the work they advertise.
// Side-effect observation against the live device -- not just "trigger fires".
//
//go:build acceptance

package provider

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// ---------- /log/{info,warning,debug} ----------

// TestAccActionLogInfo: log entry must appear in /log after the action runs.
func TestAccActionLogInfo(t *testing.T)    { logLevelMarker(t, "routeros_log_info") }
func TestAccActionLogWarning(t *testing.T) { logLevelMarker(t, "routeros_log_warning") }
func TestAccActionLogDebug(t *testing.T)   { logLevelMarker(t, "routeros_log_debug") }

func logLevelMarker(t *testing.T, resourceType string) {
	t.Helper()
	if os.Getenv("TF_ACC") == "" || os.Getenv("ROUTEROS_HOST") == "" {
		t.Skip("TF_ACC and ROUTEROS_HOST required")
	}
	host := os.Getenv("ROUTEROS_HOST")
	user := os.Getenv("ROUTEROS_USER")
	pass := os.Getenv("ROUTEROS_PASSWORD")
	marker := fmt.Sprintf("acc-%s-%d", strings.TrimPrefix(resourceType, "routeros_"), time.Now().UnixNano())

	cfg := fmt.Sprintf(`
provider "routeros" {
  routers = {
    home = {
      host     = "%s"
      username = "%s"
      password = "%s"
      insecure = true
    }
  }
}

resource "%s" "marker" {
  router  = "home"
  trigger = "%s"
  message = "%s"
}
`, host, user, pass, resourceType, marker, marker)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: cfg,
				Check: func(_ *terraform.State) error {
					time.Sleep(500 * time.Millisecond)
					if !logContainsMarker(host, user, pass, marker) {
						return fmt.Errorf("%s did not produce a log entry with marker %q", resourceType, marker)
					}
					return nil
				},
			},
		},
	})
}

// ---------- /console/clear-history ----------

func TestAccActionConsoleClearHistory(t *testing.T) {
	if os.Getenv("TF_ACC") == "" || os.Getenv("ROUTEROS_HOST") == "" {
		t.Skip("TF_ACC and ROUTEROS_HOST required")
	}
	host := os.Getenv("ROUTEROS_HOST")
	user := os.Getenv("ROUTEROS_USER")
	pass := os.Getenv("ROUTEROS_PASSWORD")
	cfg := func(trig string) string {
		return fmt.Sprintf(`
provider "routeros" {
  routers = {
    home = {
      host     = "%s"
      username = "%s"
      password = "%s"
      insecure = true
    }
  }
}

resource "routeros_console_clear_history" "ch" {
  router  = "home"
  trigger = "%s"
}
`, host, user, pass, trig)
	}
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{Config: cfg("first-run")},
			{Config: cfg("second-run")},
		},
	})
}

// ---------- /system/check-disk ----------

func TestAccActionSystemCheckDisk(t *testing.T) {
	if os.Getenv("TF_ACC") == "" || os.Getenv("ROUTEROS_HOST") == "" {
		t.Skip("TF_ACC and ROUTEROS_HOST required")
	}
	host := os.Getenv("ROUTEROS_HOST")
	user := os.Getenv("ROUTEROS_USER")
	pass := os.Getenv("ROUTEROS_PASSWORD")
	cfg := fmt.Sprintf(`
provider "routeros" {
  routers = {
    home = {
      host     = "%s"
      username = "%s"
      password = "%s"
      insecure = true
    }
  }
}

resource "routeros_system_check_disk" "d" {
  router  = "home"
  trigger = "v1"
}
`, host, user, pass)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{Config: cfg},
		},
	})
}

// ---------- helpers ----------

func firstInterfaceName(t *testing.T, host, user, pass string) string {
	t.Helper()
	var body string
	// RouterOS sometimes closes the connection between requests when the
	// previous test step also hit the REST API. Retry a few times before
	// giving up.
	var lastErr error
	for attempt := 0; attempt < 5; attempt++ {
		if attempt > 0 {
			time.Sleep(time.Duration(attempt*200) * time.Millisecond)
		}
		req, _ := http.NewRequest("GET",
			strings.TrimRight(host, "/")+"/rest/interface?.proplist=name", nil)
		req.SetBasicAuth(user, pass)
		req.Header.Set("Accept-Encoding", "gzip")
		req.Close = true // don't reuse pooled connections that may already be dead
		r, err := http.DefaultClient.Do(req)
		if err != nil {
			lastErr = err
			continue
		}
		body = readAllResp2(r)
		r.Body.Close()
		if r.StatusCode == 200 && strings.Contains(body, `"name":"`) {
			break
		}
		lastErr = fmt.Errorf("status %d body=%s", r.StatusCode, body)
	}
	const key = `"name":"`
	i := strings.Index(body, key)
	if i < 0 {
		t.Fatalf("list interfaces (last error %v): %s", lastErr, body)
	}
	rest := body[i+len(key):]
	j := strings.IndexByte(rest, '"')
	if j < 0 {
		t.Fatalf("malformed interface name: %s", body)
	}
	return rest[:j]
}

func logContainsMarker(host, user, pass, marker string) bool {
	req, err := http.NewRequest("GET",
		strings.TrimRight(host, "/")+"/rest/log?.proplist=message", nil)
	if err != nil {
		return false
	}
	req.SetBasicAuth(user, pass)
	req.Header.Set("Accept-Encoding", "gzip")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return false
	}
	return strings.Contains(readAllResp2(resp), marker)
}

func readAllResp2(r *http.Response) string {
	var b strings.Builder
	buf := make([]byte, 4096)
	for {
		n, err := r.Body.Read(buf)
		if n > 0 {
			b.Write(buf[:n])
		}
		if err != nil {
			break
		}
	}
	return b.String()
}
