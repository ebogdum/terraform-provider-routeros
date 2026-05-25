resource "routeros_queue_type" "type_example" {
  # router = "my-router"  # which router to target; omit for the default
  kind = "pfifo"
  name = "example"

  # Optional attributes (uncomment as needed):
  # mq_pfifo_limit = 0
  # pcq_burst_rate = 0
  # pcq_burst_threshold = 0
  # pcq_burst_time = "1h"
  # pcq_classifier = "replace-me"
  # pcq_dst_address_mask = 0
  # pcq_dst_address6_mask = 0
  # pcq_limit = 0
  # pcq_rate = 0
  # pcq_src_address_mask = 0
  # pcq_src_address6_mask = 0
  # pcq_total_limit = 0
  # pfifo_limit = 0
  # red_avg_packet = 0
  # red_burst = 0
  # red_limit = 0
  # red_max_threshold = 0
  # red_min_threshold = 0
  # sfq_allot = 0
  # sfq_perturb = 0
}
