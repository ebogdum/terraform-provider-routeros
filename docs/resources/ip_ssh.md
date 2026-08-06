---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_ssh"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_ssh

Manages the RouterOS `/ip/ssh` menu.

## Example Usage

```terraform
resource "routeros_ip_ssh" "ssh_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # ciphers = "replace-me"
  # forwarding_enabled = "no"
  # host_key_size = 0
  # host_key_type = "replace-me"
  # password_authentication = "replace-me"
  # publickey_authentication_options = "replace-me"
  # strong_crypto = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `ciphers` - (Optional) Type: `string`. Allow to configure SSH ciphers.
* `forwarding_enabled` - (Optional) Type: `string`. Allows to control which SSH forwarding method to allow: `no` - SSH forwarding is disabled; `local` - Allow SSH clients to originate connections from the server(router), this setting controls also dynamic forwarding; `remote` - Allow SSH clients to listen on the server(router) and forward incoming connections; `both` - Allow both local and remote forwarding methods.
* `host_key_size` - (Optional) Type: `int`. RSA key size when host key is being regenerated.
* `host_key_type` - (Optional) Type: `string`. Select host key type
* `password_authentication` - (Optional) Type: `string`. Whether to allow password login at the same time when public key authorization is configured for a user.
* `publickey_authentication_options` - (Optional) Type: `string`. Sets public key authentication options. The touch-required option causes public key authentication using a FIDO authenticator algorithm to always require the signature to attest that a physically present user explicitly confirmed the authentication (usually by touching the authenticator). The verify-required option requires a FIDO key signature attest that the user was verified, e.g. via a PIN.
* `strong_crypto` - (Optional) Type: `bool`. Use stronger encryption, HMAC algorithms, use bigger DH primes and disallow weaker ones: use 256 and 192 bit encryption instead of 128 bits; disable null encryption; use sha256 for hashing instead of sha1; disable md5; use 2048bit prime for Diffie-Hellman exchange instead of 1024bit.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_ip_ssh.this 'home'
```
