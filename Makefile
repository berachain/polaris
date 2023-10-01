#!/usr/bin/make -f

# Makefile

# Specify the default target if none is provided
.DEFAULT_GOAL := codeqlbuild

MODULES := $(shell find . -type f -name 'go.mod' -exec dirname {} \;)
# Exclude root module
MODULES := $(filter-out ./,$(MODULES))

# Helper rule to display available targets
help:
	@echo "This Makefile is an alias for Mage tasks. Run 'mage' to see available Mage targets."
	@echo "You can use 'make <target>' to call the corresponding 'mage <target>' command."

# Required rule for gh action codeql to run.
codeqlbuild:
	@mage build

# Rule to forward any target to Mage
%:
	@mage $@

# Rule to setup the project. This is a special case because it's not a Mage target.
setup:
	@go run magefiles/setup/setup.go

BINDIR ?= $(GOPATH)/bin
CURRENT_DIR = $(shell pwd)

test-sim-after-import:
	@echo "Running application simulation-after-import. This may take several minutes..."
	@cd ${CURRENT_DIR}/e2e/testapp && $(BINDIR)/runsim -Jobs=4 -SimAppPkg=. -ExitOnFail 50 5 TestAppSimulationAfterImport


###############################################################################
###                                  Build                                  ###
###############################################################################




###############################################################################
###                           Tests & Simulation                            ###
###############################################################################

###############################################################################
###                              Formatting                                 ###
###############################################################################

###############################################################################
###                                Linting                                  ###
###############################################################################

format:
	@$(MAKE) license-fix buf-lint-fix forge-lint-fix golines golangci-fix

lint:
	@$(MAKE) license buf-lint forge-lint golangci gosec


#################
#     forge     #
#################

forge-lint-fix:
	@echo "--> Running forge fmt"
	@cd ./contracts && forge fmt

forge-lint:
	@echo "--> Running forge lint"
	@cd ./contracts && forge fmt --check

#################
# golangci-lint #
#################

golangci_version=v1.54.2

golangci-install:
	@echo "--> Installing golangci-lint $(golangci_version)"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version)

golangci:
	@$(MAKE) golangci-install
	@echo "--> Running linter"
	@go list -f '{{.Dir}}/...' -m | xargs golangci-lint run  --timeout=10m --concurrency 8 -v 

golangci-fix:
	@$(MAKE) golangci-install
	@echo "--> Running linter"
	@go list -f '{{.Dir}}/...' -m | xargs golangci-lint run  --timeout=10m --fix --concurrency 8 -v 


#################
#    golines    #
#################

golines-install:
	@echo "--> Installing golines"
	@go install github.com/segmentio/golines

golines:
	@$(MAKE) golines-install
	@echo "--> Running golines"
	@golines --reformat-tags --shorten-comments --write-output --max-len=99 -l ./.


#################
#    license    #
#################

license-install:
	@echo "--> Installing google/addlicense"
	@go install github.com/google/addlicense

license:
	@$(MAKE) license-install
	@echo "--> Running addlicense with -check"
	@for module in $(MODULES); do \
		(cd $$module && addlicense -check -v -f ./LICENSE.header ./.) || exit 1; \
	done

license-fix:
	@$(MAKE) license-install
	@echo "--> Running addlicense"
	@for module in $(MODULES); do \
		(cd $$module && addlicense -v -f ./LICENSE.header ./.) || exit 1; \
	done


#################
#     gosec     #
#################

gosec-install:
	@echo "--> Installing gosec"
	@go install github.com/securego/gosec/v2/cmd/gosec

gosec:
	@$(MAKE) gosec-install
	@echo "--> Running gosec"
	@gosec -exclude-generated ./...


#################
#     proto     #
#################

protoDir := "cosmos/proto"

buf-install:
	@echo "--> Installing buf"
	@go install github.com/bufbuild/buf/cmd/buf

buf-lint-fix:
	@$(MAKE) buf-install 
	@echo "--> Running buf format"
	@buf format -w --error-format=json $(protoDir)

buf-lint:
	@$(MAKE) buf-install 
	@echo "--> Running buf lint"
	@buf lint --error-format=json $(protoDir)


###############################################################################
###                             Dependencies                                 ###
###############################################################################

sync: |
	@for module in $(MODULES); do \
		echo "Running go mod download in $$module"; \
		(cd $$module && go mod download) || exit 1; \
	done
	go work sync

tidy: |
	@for module in $(MODULES); do \
		echo "Running go mod tidy in $$module"; \
		(cd $$module && go mod tidy) || exit 1; \
	done
.PHONY: lint lint-fix
