---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface"
description: |-
  /interface is mostly read-only -- interfaces are created by their specific subtypes (/interface/bridge, /interface/wireguard, etc.).
---

# Resource: routeros_interface

/interface is mostly read-only -- interfaces are created by their specific subtypes (/interface/bridge, /interface/wireguard, etc.).

## Example Usage

```terraform
resource "routeros_interface" "interface_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # add = "replace-me"
  # edit = "replace-me"
  # find = "replace-me"
  # move = "replace-me"
  # mtu = 1500
  # name = "tf-example"
  # print = "replace-me"
  # remove = "replace-me"
  # set = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `add` - (Optional) Type: `string`. This command usually has all the same arguments as a set, except the item number argument. It adds a new item with the values you have specified, usually at the end of the item list, in places where the order of items is relevant. There are some required properties that you have to supply, such as the interface for a new address, while other properties are set to defaults unless you explicitly specify them. Common Parameters copy-from - Copies an existing item. It takes default values of a new item's properties from another item. If you do not want to make an exact copy, you can specify new values for some properties. When copying items that have names, you will usually have to give a new name to a copy place-before - places a new item before an existing item with a specified position. Thus, you do not need to use the move command after adding an item to the list disabled - controls disabled/enabled state of the newly added item(-s) comment - holds the description of a newly created item Return Values add command returns the internal number of items it has added.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`.
* `edit` - (Optional) Type: `string`. This command is associated with the set command. It can be used to edit values of properties that contain a large amount of text, such as scripts, but it works with all editable properties. Depending on the capabilities of the terminal, either a full-screen editor or a single line editor is launched to edit the value of the specified property.
* `find` - (Optional) Type: `string`. The find command has the same arguments as a set, plus the flag arguments like disabled or active that take values yes or no depending on the value of the respective flag. To see all flags and their names, look at the top of the print command's output. The find command returns internal numbers of all items that have the same values of arguments as specified.
* `move` - (Optional) Type: `string`. Changes the order of items in the list. Parameters: the first argument specifies the item(-s) being moved. the second argument specifies the item before which to place all items being moved (they are placed at the end of the list if the second argument is omitted).
* `mtu` - (Optional) Type: `int`. Default: `1500`.
* `name` - (Optional) Type: `string`.
* `print` - (Optional) Type: `string`. Shows all information that's accessible from a particular command level. Thus, /system clock print shows the system date and time, /ip route print shows all routes etc. If there\'s a list of items in the current level and they are not read-only, i.e. you can change/remove them (example of read-only item list is /system history , which shows a history of executed actions), then print command also assigns numbers that are used by all commands that operate with items in this list. Common Parameters: append -  brief - forces the print command to use tabular output form count-only - shows the number of items detail - forces the print command to use property=value output form file - prints the contents of the specific sub-menu into a file on the router. follow -  follow-only -  follow-strict -  from - show only specified items, in the same order in which they are given. interval - updates the output from the print command for every interval of seconds. oid - prints the OID value for properties that are accessible from SNMP. proplist - comma-separated and ordered list of property names that should be included for the returned items. show-ids -  where - show only items that match specified criteria. The syntax of where the property is similar to the find command. without-paging - prints the output without stopping after each screenful.
* `remove` - (Optional) Type: `string`. Removes specified item(-s) from a list.
* `set` - (Optional) Type: `string`. Allows you to change values of general parameters or item parameters. The set command has arguments with names corresponding to values you can change. Use ? or double Tab to see a list of all arguments. If there is a list of items in this command level, then the set has one action argument that accepts the number of items (or list of numbers) you wish to set up. This command does not return anything.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `actual_mtu` - Type: `int`.
* `default_name` - Type: `string`.
* `fp_rps_drop` - Type: `int`.
* `fp_rx_byte` - Type: `int`.
* `fp_rx_packet` - Type: `int`.
* `fp_tx_byte` - Type: `int`.
* `fp_tx_packet` - Type: `int`.
* `last_link_up_time` - Type: `string`.
* `link_downs` - Type: `int`.
* `mac_address` - Type: `mac`.
* `running` - Type: `bool`.
* `rx_byte` - Type: `int`.
* `rx_drop` - Type: `int`.
* `rx_error` - Type: `int`.
* `rx_packet` - Type: `int`.
* `tx_byte` - Type: `int`.
* `tx_drop` - Type: `int`.
* `tx_error` - Type: `int`.
* `tx_packet` - Type: `int`.
* `tx_queue_drop` - Type: `int`.
* `type` - Type: `int`.
* `vrf` - Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface.example '*3'

# Named router
terraform import routeros_interface.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface.example 'home/my-resource-name'
```
