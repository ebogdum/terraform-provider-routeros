---
page_title: "RouterOS: routeros_container"
description: |-
  Requires container package + capable architecture
---

# Resource: routeros_container

Requires container package + capable architecture

## Example Usage

```terraform
resource "routeros_container" "container_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # auto_restart_interval = "replace-me"
  # cmd = "replace-me"
  # cpu_list = "replace-me"
  # devices = "replace-me"
  # dns = "replace-me"
  # domain_name = "replace-me"
  # entrypoint = "replace-me"
  # envlist = "replace-me"
  # hostname = "replace-me"
  # interface = "ether1"
  # logging = "replace-me"
  # memory_high = "replace-me"
  # memory_max = "replace-me"
  # mount = "replace-me"
  # mountlists = "replace-me"
  # remote_image = "replace-me"
  # root_dir = "replace-me"
  # start_on_boot = "replace-me"
  # stop_signal = "replace-me"
  # user = "myuser"
  # workdir = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `auto_restart_interval` - (Optional) Type: `string`. Specify an interval at which Container will be restarted on Container failure. Example: 10s.
* `cmd` - (Optional) Type: `string`. The main purpose of a CMD is to provide defaults for an executing container. These defaults can include an executable, or they can omit the executable, in which case you must specify an ENTRYPOINT instruction as well.
* `comment` - (Optional) Type: `string`. Short description.
* `cpu_list` - (Optional) Type: `string`. specifies which CPU cores the container is allowed to run on.
* `devices` - (Optional) Type: `string`. passes through physical device to the container.
* `dns` - (Optional) Type: `string`. If container needs different DNS, it can be configured here.
* `domain_name` - (Optional) Type: `string`.
* `entrypoint` - (Optional) Type: `string`. An ENTRYPOINT allows to specify executable to run when starting container. Example: /bin/sh.
* `envlist` - (Optional) Type: `string`. list of environmental variables (configured under /container envs ) to be used with container.
* `hostname` - (Optional) Type: `string`. Assigning a hostname to a container helps in identifying and managing the container more easily.
* `interface` - (Optional) Type: `string`. veth interface to be used with the container.
* `logging` - (Optional) Type: `string`. if set to yes, all container-generated output will be shown in the RouterOS log.
* `memory_high` - (Optional) Type: `string`. RAM usage limit in bytes for a specific container.
* `memory_max` - (Optional) Type: `string`. max RAM usage limit in bytes per container (The container process will be terminated if the memory-max value is smaller than the container memory-current.).
* `mount` - (Optional) Type: `string`. specify directory to be used as a mount.
* `mountlists` - (Optional) Type: `string`. mounts from /container/mounts/ sub-menu to be used with this container.
* `remote_image` - (Optional) Type: `string`. the container image name to be installed if an external registry is used (configured under /container/config set registry-url=...).
* `root_dir` - (Optional) Type: `string`. used to save container store outside main memory.
* `start_on_boot` - (Optional) Type: `string`. if set to yes, the container will be started automatically on device start-up.
* `stop_signal` - (Optional) Type: `string`. Type of Linux signal to send when container was not stopped after 10 seconds.
* `user` - (Optional) Type: `string`. sets the user and group the container process runs as before execution.
* `workdir` - (Optional) Type: `string`. the working directory for cmd entrypoint.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_container.example '*3'

# Named router
terraform import routeros_container.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_container.example 'home/my-resource-name'
```
