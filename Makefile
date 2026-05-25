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
