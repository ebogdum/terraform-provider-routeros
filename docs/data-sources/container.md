---
page_title: "RouterOS: routeros_container"
description: |-
  Requires container package + capable architecture
---

# Data Source: routeros_container

Requires container package + capable architecture

## Example Usage

```terraform
data "routeros_container" "container_example" {
  # router   = "my-router"  # omit for the default router
  # filter   = { name = "some-name" }
  # proplist = ["name", "address"]
}
```

## Argument Reference

This data source supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to query.
* `filter` - (Optional) Map of field=value pairs to narrow the result set.
* `proplist` - (Optional) List of property names to project; smaller payload.
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

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

