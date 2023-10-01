#!/usr/bin/make -f
include scripts/cosmos.mk scripts/constants.mk


# Specify the default target if none is provided
.DEFAULT_GOAL := build

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


###############################################################################
###                                  Build                                  ###
###############################################################################

BUILD_TARGETS := build install

build: BUILD_ARGS=-o $(OUT_DIR)/

build-linux-amd64:
	GOOS=linux GOARCH=amd64 LEDGER_ENABLED=false $(MAKE) build

build-linux-arm64:
	GOOS=linux GOARCH=arm64 LEDGER_ENABLED=false $(MAKE) build

$(BUILD_TARGETS): tidy $(OUT_DIR)/
	@echo "Building ${TESTAPP_DIR}"
	@cd ${CURRENT_DIR}/$(TESTAPP_DIR) && go $@ -mod=readonly $(BUILD_FLAGS) $(BUILD_ARGS) ./...

$(OUT_DIR)/:
	mkdir -p $(OUT_DIR)/

# build:
# 	@$(MAKE) forge-build

build-clean: 
	@$(MAKE) clean build

clean:
	@$(MAKE) forge-clean

#################
#     forge     #
#################

forge-build: |
	@forge build --extra-output-files bin --extra-output-files abi  --root $(CONTRACTS_DIR)

forge-clean: |
	@forge clean --root $(CONTRACTS_DIR)


#################
#     proto     #
#################

protoImageName    := "ghcr.io/cosmos/proto-builder"
protoImageVersion := "0.14.0"

proto:
	@$(MAKE) buf-lint-fix buf-lint proto-build

proto-build:
	@docker run --rm -v ${CURRENT_DIR}:/workspace --workdir /workspace $(protoImageName):$(protoImageVersion) sh ./cosmos/proto/scripts/proto_generate.sh


###############################################################################
###                           Tests & Simulation                            ###
###############################################################################

#################
#    polard     #
#################

start:
	@./e2e/testapp/entrypoint.sh

#################
#     unit      #
#################

install-ginkgo:
	@echo "Installing ginkgo..."
	@go install github.com/onsi/ginkgo/v2/ginkgo@latest

test-unit:
	@$(MAKE) install-ginkgo forge-test
	@echo "Running unit tests..."
	@ginkgo -r --randomize-all --fail-on-pending -trace --skip .*e2e* ./...

test-unit-race:
	@$(MAKE) install-ginkgo forge-test
	@echo "Running unit tests with race detection..."
	@ginkgo --race -r --randomize-all --fail-on-pending -trace --skip .*e2e* ./...

test-unit-cover:
	@$(MAKE) install-ginkgo forge-test
	@echo "Running unit tests with coverage..."
	@ginkgo -r --randomize-all --fail-on-pending -trace --skip .*e2e* \
	--junit-report out.xml --cover --coverprofile "coverage-testunitcover.txt" --covermode atomic \
		./...

#################
#     forge     #
#################

forge-test:
	@echo "Running forge test..."
	@forge test --root $(CONTRACTS_DIR)

#################
#      e2e      #
#################

test-e2e:
	# TODO: docker build before running
	@$(MAKE) test-e2e-no-build

test-e2e-no-build:
	@$(MAKE) install-ginkgo
	@echo "Running localnet tests..."
	@ginkgo -r --randomize-all --fail-on-pending -trace -timeout 30m ./e2e/precompiles/...


#################
#     hive      #
#################

#################
#   localnet    #
#################

test-localnet:
	# TODO: docker build before running
	@$(MAKE) test-localnet-no-build

test-localnet-no-build:
	@$(MAKE) install-ginkgo
	@echo "Running localnet tests..."
	@ginkgo -r --randomize-all --fail-on-pending -trace -timeout 30m ./e2e/localnet/...



test-sim-after-import:
	@echo "Running application simulation-after-import. This may take several minutes..."
	@cd ${CURRENT_DIR}/e2e/testapp && $(BINDIR)/runsim -Jobs=4 -SimAppPkg=. -ExitOnFail 50 5 TestAppSimulationAfterImport


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
	@cd $(CONTRACTS_DIR) && forge fmt

forge-lint:
	@echo "--> Running forge lint"
	@cd $(CONTRACTS_DIR) && forge fmt --check

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
###                             Dependencies                                ###
###############################################################################

sync: |
	@for module in $(MODULES); do \
		echo "Running go mod download in $$module"; \
		(cd $$module && go mod download) || exit 1; \
	done
	@echo "Running go mod sync"
	@go work sync

tidy: |
	@for module in $(MODULES); do \
		echo "Running go mod tidy in $$module"; \
		(cd $$module && go mod tidy) || exit 1; \
	done

