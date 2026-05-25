resource "routeros_tool_traffic_generator" "traffic_generator_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # latency_distribution_max = "1h"
  # measure_out_of_order = false
  # stats_samples_to_keep = 0
  # test_id = 0
}
