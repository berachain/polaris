# Makefile

# Specify the default target if none is provided
.DEFAULT_GOAL := help

# Helper rule to display available targets
help:
	@echo "This Makefile is an alias for Mage tasks. Run 'mage' to see available Mage targets."
	@echo "You can use 'make <target>' to call the corresponding 'mage <target>' command."

# Rule to forward any target to Mage
%:
	@mage $@

# Rule to setup the project. This is a special case because it's not a Mage target.
setup:
	@go run magefiles/setup/setup.go

PACKAGE_NAME:=pkg.berachain.dev/polaris/cosmos
GOLANG_CROSS_VERSION = v1.20.4
GOPATH ?= '$(HOME)/go'
release-dry-run:
	docker run \
		--rm \
		--privileged \
		-e CGO_ENABLED=1 \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/$(PACKAGE_NAME) \
		-v ${GOPATH}/pkg:/go/pkg \
		-w /go/src/$(PACKAGE_NAME) \
		ghcr.io/goreleaser/goreleaser-cross:${GOLANG_CROSS_VERSION} \
		--clean --skip-validate --skip-publish --snapshot

.PHONY: release-dry-run release