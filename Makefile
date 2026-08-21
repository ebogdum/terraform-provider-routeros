SHELL := /usr/bin/env bash
GO    ?= go
BIN   := bin
PKGS  := ./...

.PHONY: all
all: build

.PHONY: tidy
tidy:
	$(GO) mod tidy

.PHONY: fmt
fmt:
	$(GO) fmt $(PKGS)

.PHONY: vet
vet:
	$(GO) vet $(PKGS)
	$(GO) vet -tags acceptance $(PKGS)

.PHONY: build
build:
	mkdir -p $(BIN)
	$(GO) build -o $(BIN)/terraform-provider-routeros .

.PHONY: test
test:
	$(GO) test -count=1 -race $(PKGS)

.PHONY: testacc
testacc:
	TF_ACC=1 ROUTEROS_RUN_DESTRUCTIVE_ACTIONS=1 \
	  $(GO) test -tags acceptance -count=1 -timeout 120m ./internal/provider/... -v

.PHONY: release-snapshot
release-snapshot:
	goreleaser release --snapshot --clean --skip=publish,sign

.PHONY: release
release:
	@# Tag the release first: git tag -s vX.Y.Z -m '...'
	@# Requires GPG_FINGERPRINT in env for signing.
	goreleaser release --clean

.PHONY: clean
clean:
	rm -rf $(BIN) dist

# --- live-device conformance -------------------------------------------------
# Needs ROUTEROS_HOST / ROUTEROS_USER / ROUTEROS_PASSWORD.

.PHONY: audit
audit:
	python3 tools/conformance/schema_audit.py all

# Dump a device's menu tree for hardware the maintainers do not have.
# Names and structure only: no values, no counts. See COLLECTING.md.
.PHONY: collect
collect:
	@python3 tools/conformance/collect_schema.py $(COLLECT_ARGS)

.PHONY: dms-backup dms-install dms-arm dms-disarm dms-verify dms-remove
dms-backup:  ; tools/conformance/deadman.sh backup
dms-install: ; tools/conformance/deadman.sh install
dms-arm:     ; tools/conformance/deadman.sh arm
dms-disarm:  ; tools/conformance/deadman.sh disarm
dms-verify:  ; tools/conformance/deadman.sh verify
dms-remove:  ; tools/conformance/deadman.sh remove

.PHONY: conformance-manifest conformance
conformance-manifest:
	python3 tools/conformance/gen_manifest.py > internal/provider/testdata/conformance.json

conformance:
	TF_ACC=1 $(GO) test -tags acceptance -count=1 -timeout 180m \
	  -run TestAccConformanceSweep ./internal/provider/ -v
