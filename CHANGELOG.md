# Changelog

All notable changes to this project are documented here.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.0.2] - 2026-05-24

- Provider landing page (`docs/index.md`) rewritten so the Terraform
  Registry renders it correctly: dropped the manually rolled `## Schema`
  block that collided with the Registry's auto-rendered provider schema.
- Per-resource and per-data-source docs now follow the standard
  `# Resource: name`, `## Argument Reference`, `## Attribute Reference`,
  `## Import` layout used by the AWS, Google, and Azure providers.
- Hand-written multi-resource workflow examples moved to `docs/guides/`
  (firewall, home-network, wireguard-vpn, certificate, system).

## [1.0.1] - 2026-05-24

- Release manifest checksum (`terraform-registry-manifest.json`) is now
  included in `SHA256SUMS`, unblocking Registry validation.
- Releases publish directly (no draft).

## [1.0.0] - 2026-05-24

First release.

- 186 resources, 277 data sources, 75 actions, 3522 properties.
- Multi-router from a single provider block via the `routers` map; every
  resource and data source takes an optional `router` attribute.
- Deterministic firewall rule ordering through a per-rule `position`
  integer; order survives destroy and recreate without churn.
- Per-row actions (sign, format, eject, reset-counters, ...) driven via
  `target_id` and a free-form `params` map.
- Lockout safety guards on firewall, user, user-group, and mac-server
  resources. Override per resource with `lockout_ack = true`.
- Sensitive properties (passwords, secrets, keys, OTP, SIM-PIN) flagged
  so Terraform redacts them in plan output, state, and CLI display.
- Drift detection: out-of-band REST changes detected on the next plan
  and reconciled by the following apply.
- Import support: `terraform import resource.x "<router>/<.id>"`.
- Five end-to-end workflow guides verified live against a CHR
  (firewall, home-network, wireguard-vpn, certificate, system).
