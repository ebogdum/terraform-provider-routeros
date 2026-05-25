# Changelog

All notable changes to this project are documented here.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.0.0] - 2026-05-24

First production release.

### Coverage

- 420 menus mapped to Terraform: 182 resources, 276 data sources, 75 actions.
- 3289 properties total, 0 device-rejected, 281 marked read-only, 28 marked sensitive.
- Schema source-of-truth committed as data; the generated Go in
  `internal/provider/*_gen.go` is self-contained -- no build- or run-time
  dependency on any external schema-sync tooling.

### Provider features

- Multi-router support via a single `provider "routeros" { routers = { ... } }`
  block. Each resource takes an optional `router` attribute; omit for the
  default router.
- Deterministic firewall rule ordering through a per-rule `position` integer.
  Order is preserved across destroy/recreate without churn.
- Lockout safety guards on firewall, user, user-group, and mac-server resources.
  Refuses to apply changes that would obviously sever management access; can be
  bypassed per resource with `lockout_ack = true`.
- Sensitive property handling: password / secret / key / token / OTP / SIM-PIN
  fields are marked sensitive and Terraform redacts them in plan output and
  state.
- Drift correction verified end-to-end: out-of-band changes via REST are
  detected on the next plan and reconciled by the following apply.
- Import support: `terraform import resource.x "<router>/<.id>"` hydrates state
  from the live device.

### Tested

- 190/190 acceptance tests pass against the test device (188 generated
  resource/action tests + handwritten lifecycle, drift correction, action
  result verification, and import tests).
- `examples/curated/home-router/` config applied end-to-end: 7 resources
  created, re-plan returns zero drift, destroy returns 7 destroyed.
- `tfplugindocs validate` returns zero missing and zero extraneous files.
- `goreleaser release --snapshot` builds 13 archives across linux / darwin /
  windows / freebsd × amd64 / arm64 / 386 / arm with signed checksums.

### Known limitations

- The committed schema reflects RouterOS 7.22.3 as observed on a single CHR
  device. Menus that require hardware not present on the test device (PoE,
  w60g, LTE, GPS, container, LCD, IoT/LoRa, switch chip) are emitted as
  resources but have no acceptance coverage; they are individually skipped via
  overlays.
- Acceptance test matrix across ROS versions is configured in
  `.github/workflows/ci.yml` but requires self-hosted runners and per-version
  GitHub environments (`ros-7.20`, `ros-7.22`, `ros-latest`) to actually run.
