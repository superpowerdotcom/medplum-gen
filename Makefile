SPEC_FILE ?= specs/$(SPEC_VERSION).json
SPEC_VERSION ?= $(shell cat version | sed 's/\./_/g')

# Pattern #1 example: "example : description = Description for example target"
# Pattern #2 example: "### Example separator text
help: HELP_SCRIPT = \
	if (/^([a-zA-Z0-9-\.\/]+).*?: description\s*=\s*(.+)/) { \
		printf "\033[34m%-40s\033[0m %s\n", $$1, $$2 \
	} elsif(/^\#\#\#\s*(.+)/) { \
		printf "\033[33m>> %s\033[0m\n", $$1 \
	}

.PHONY: help
help:
	@perl -ne '$(HELP_SCRIPT)' $(MAKEFILE_LIST)

### Setup

.PHONY: setup
setup: description = Install dependencies necessary for codegen
setup:
	brew install openapi-generator

### Build

.PHONY: generate/go
generate/go: description = Generate Go code from Medplum OpenAPI spec
generate/go: clean/go
	mkdir -p ./build/go
	openapi-generator generate \
	--skip-validate-spec \
	--package-name medplum \
	-i $(SPEC_FILE) \
	-g go -o build/go || @echo "Failed to generate Go code"
	$(RM) -r ./build/go/docs ./build/go/api

.PHONY: clean/go
clean/go: description = Remove Go build artifacts
clean/go:
	$(RM) -r ./build/go
