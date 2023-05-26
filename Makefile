# Makefile

# Specify the default target if none is provided
.DEFAULT_GOAL := codeqlbuild

# Helper rule to display available targets
help:
	@echo "This Makefile is an alias for Mage tasks. Run 'mage' to see available Mage targets."
	@echo "You can use 'make <target>' to call the corresponding 'mage <target>' command."

# Required rule for gh action codeql to run.
codeqlbuild:
	@mage contracts:build generate build

# Rule to forward any target to Mage
%:
	@mage $@

# Rule to setup the project. This is a special case because it's not a Mage target.
setup:
	@go run magefiles/setup/setup.go