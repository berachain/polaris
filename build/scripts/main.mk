#!/usr/bin/make -f
include build/scripts/cosmos.mk build/scripts/constants.mk


# Specify the default target if none is provided
.DEFAULT_GOAL := build

###############################################################################
###                                  Build                                  ###
###############################################################################

BUILD_TARGETS := build install

build: BUILD_ARGS=-o $(OUT_DIR)/

build-linux-amd64:
	GOOS=linux GOARCH=amd64 LEDGER_ENABLED=false $(MAKE) build

build-linux-arm64:
	GOOS=linux GOARCH=arm64 LEDGER_ENABLED=false $(MAKE) build

$(BUILD_TARGETS): forge-build sync $(OUT_DIR)/
	@echo "Building ${TESTAPP_DIR}"
	@cd ${CURRENT_DIR}/$(TESTAPP_DIR) && go $@ -mod=readonly $(BUILD_FLAGS) $(BUILD_ARGS) ./...

$(OUT_DIR)/:
	mkdir -p $(OUT_DIR)/

build-clean: 
	@$(MAKE) clean build

clean:
	@rm -rf .tmp/ 
	@rm -rf $(OUT_DIR)
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
	@docker run --rm -v ${CURRENT_DIR}:/workspace --workdir /workspace $(protoImageName):$(protoImageVersion) sh ./build/scripts/proto_generate.sh

###############################################################################
###                                 Docker                                  ###
###############################################################################

# Variables
DOCKER_TYPE ?= base
ARCH ?= arm64
GO_VERSION ?= 1.21.3
IMAGE_NAME ?= polard
IMAGE_VERSION ?= v0.0.0
BASE_IMAGE ?= polard/base:$(IMAGE_VERSION)

# Docker Paths
BASE_DOCKER_PATH = ./e2e/testapp/docker
EXEC_DOCKER_PATH = $(BASE_DOCKER_PATH)/base.Dockerfile
LOCAL_DOCKER_PATH = $(BASE_DOCKER_PATH)/local/Dockerfile
SEED_DOCKER_PATH =  $(BASE_DOCKER_PATH)/seed/Dockerfile
VAL_DOCKER_PATH =  $(BASE_DOCKER_PATH)/validator/Dockerfile
LOCALNET_CLIENT_PATH = ./e2e/precompile/polard
LOCALNET_DOCKER_PATH = $(LOCALNET_CLIENT_PATH)/Dockerfile

# Image Build
docker-build:
	@echo "Build a release docker image for the Cosmos SDK chain..."
	@$(MAKE) docker-build-$(DOCKER_TYPE)

# Docker Build Types
docker-build-base:
	$(call docker-build-helper,$(EXEC_DOCKER_PATH),base)

docker-build-local:
	$(call docker-build-helper,$(LOCAL_DOCKER_PATH),local,--build-arg BASE_IMAGE=$(BASE_IMAGE))

docker-build-seed:
	$(call docker-build-helper,$(SEED_DOCKER_PATH),seed,--build-arg BASE_IMAGE=$(BASE_IMAGE))

docker-build-validator:
	$(call docker-build-helper,$(VAL_DOCKER_PATH),validator,--build-arg BASE_IMAGE=$(BASE_IMAGE))

docker-build-localnet:
	$(call docker-build-helper,$(LOCALNET_DOCKER_PATH),localnet,--build-arg BASE_IMAGE=$(BASE_IMAGE))

# Docker Build Function
define docker-build-helper
	docker build \
	--build-arg GO_VERSION=$(GO_VERSION) \
	--platform linux/$(ARCH) \
	--build-arg PRECOMPILE_CONTRACTS_DIR=$(CONTRACTS_DIR) \
	--build-arg GOOS=linux \
	--build-arg GOARCH=$(ARCH) \
	-f $(1) \
	-t $(IMAGE_NAME)/$(2):$(IMAGE_VERSION) \
	$(if $(3),$(3)) \
	.

endef

.PHONY: docker-build-localnet

###############################################################################
###                                 CodeGen                                 ###
###############################################################################

generate:
	@$(MAKE) abigen-install moq-install mockery
	@for module in $(MODULES); do \
		echo "Running go generate in $$module"; \
		(cd $$module && go generate ./...) || exit 1; \
	done

abigen-install:
	@echo "--> Installing abigen"
	@go install github.com/ethereum/go-ethereum/cmd/abigen@latest

moq-install:
	@echo "--> Installing moq"
	@go install github.com/matryer/moq@latest

mockery-install:
	@echo "--> Installing mockery"
	@go install github.com/vektra/mockery/v2@latest

mockery:
	@$(MAKE) mockery-install
	@echo "Running mockery..."
	@mockery


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
	--junit-report out.xml --cover --coverprofile "coverage-test-unit-cover.txt" --covermode atomic \
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
	@$(MAKE) test-e2e-no-build

test-e2e-no-build:
	@$(MAKE) install-ginkgo
	@echo "Running e2e tests..."
	@ginkgo -r --randomize-all --fail-on-pending -trace -timeout 30m ./e2e/precompile/...


#################
#     hive      #
#################

ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

HIVE_CLONE := $(GOPATH)/src
CLONE_PATH := $(HIVE_CLONE)/.hive-e2e
SIMULATORS_ROOT := $(CLONE_PATH)/simulators
SIMULATORS_PATH := $(SIMULATORS_ROOT)/polaris/
BASE_HIVE_DOCKER_PATH := ./e2e/hive
CLIENTS_PATH := $(CLONE_PATH)/clients/polard/
SIMULATIONS := \
	rpc:init/genesis.json:ethclient.hive \
	rpc-compat:Dockerfile:tests \

# .PHONY: setup test testv view

hive-setup:
	@echo $(HIVE_CLONE)
	@echo "--> Setting up Hive testing environment..."
	@test ! -d $(HIVE_CLONE) && mkdir $(HIVE_CLONE) || true
	@rm -rf $(CLONE_PATH)
	@git clone https://github.com/ethereum/hive $(CLONE_PATH) --depth=1
	@mkdir $(SIMULATORS_PATH)
	@cp -rf $(BASE_HIVE_DOCKER_PATH)/clients/polard $(CLIENTS_PATH)
	@echo "Copying files...";
	@$(foreach sim,$(SIMULATIONS), \
		$(eval SIM_NAME = $(word 1, $(subst :, ,$(sim)))) \
		$(eval FILES = $(wordlist 2, $(words $(subst :, ,$(sim))), $(subst :, ,$(sim)))) \
		cp -rf $(SIMULATORS_ROOT)/ethereum/$(SIM_NAME) $(SIMULATORS_PATH); \
		$(foreach file,$(FILES), \
			cp -rf $(BASE_HIVE_DOCKER_PATH)/simulators/$(SIM_NAME)/$(file) \
			$(SIMULATORS_PATH)/$(SIM_NAME)/$(file); \
			if [ "$(file)" = "ethclient.hive" ]; then \
				cp -rf $(SIMULATORS_PATH)/$(SIM_NAME)/$(file) $(SIMULATORS_PATH)/$(SIM_NAME)/ethclient.go; \
			fi; \
		) \
	)
	@cd $(CLONE_PATH) && go install ./...

hive-view:
	@cd $(CLONE_PATH) && \
		go build ./cmd/hiveview && \
		hiveview --serve

# SHELL := /bin/zsh  # Explicitly set to zsh as that is what you are using

test-hive:
	@cd $(CLONE_PATH) && \
		hive --sim polaris/rpc -client polard

test-hive-v:
	@cd $(CLONE_PATH) && \
		hive --sim polaris/rpc -client polard --docker.output



#################
#   localnet    #
#################

test-localnet:
	@$(MAKE) test-localnet-no-build

test-localnet-no-build:
	@$(MAKE) install-ginkgo
	@echo "Running localnet tests..."
	@ginkgo -r --randomize-all --fail-on-pending -trace -timeout 30m ./e2e/localnet/...


###############################################################################
###                              Formatting                                 ###
###############################################################################

###############################################################################
###                                Linting                                  ###
###############################################################################

format:
	@$(MAKE) license-fix buf-lint-fix forge-lint-fix golangci-fix

lint:
	@$(MAKE) license buf-lint forge-lint golangci gosec nilaway-install nilaway


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
#    license    #
#################

license-install:
	@echo "--> Installing google/addlicense"
	@go install github.com/google/addlicense

license:
	@$(MAKE) license-install
	@echo "--> Running addlicense with -check"
	@for module in $(MODULES); do \
		(cd $$module && addlicense -check -v -f ./LICENSE.header -ignore "lib/forge-std/**/*" -ignore "lib/solmate/**/*" ./.) || exit 1; \
	done

license-fix:
	@$(MAKE) license-install
	@echo "--> Running addlicense"
	@for module in $(MODULES); do \
		(cd $$module && addlicense -v -f ./LICENSE.header ./.) || exit 1; \
	done

#################
#    nilaway    #
#################

nilaway-install:
	@echo "--> Installing nilaway"
	@go install go.uber.org/nilaway/cmd/nilaway@latest

nilaway:
	@for module in $(MODULES); do \
		(cd $$module && find . -type f -name '*.go' ! -name '*.abigen.go' -exec nilaway {} \;) || exit 1; \
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

protoDir := "proto"

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

repo-rinse: |
	git clean -xfd
	git submodule foreach --recursive git clean -xfd
	git submodule foreach --recursive git reset --hard
	git submodule update --init --recursive


.PHONY: build build-linux-amd64 build-linux-arm64 \
	$(BUILD_TARGETS) $(OUT_DIR)/ build-clean clean \
	forge-build forge-clean proto proto-build docker-build \
	docker-build-base docker-build-local docker-build-seed \
	docker-build-validator docker-build-localnet generate \
	abigen-install moq-install mockery-install mockery \
	start test-unit test-unit-race test-unit-cover forge-test \
	test-e2e test-e2e-no-build hive-setup hive-view test-hive \
	test-hive-v test-localnet test-localnet-no-build format lint \
	forge-lint-fix forge-lint golangci-install golangci golangci-fix \
	license-install license license-fix \
	gosec-install gosec buf-install buf-lint-fix buf-lint sync tidy repo-rinse
