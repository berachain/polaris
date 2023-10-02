#!/usr/bin/make -f
MODULES := $(shell find . -type f -name 'go.mod' -exec dirname {} \;)
# Exclude root module
MODULES := $(filter-out ./,$(MODULES))
CONTRACTS_DIR := ./contracts

