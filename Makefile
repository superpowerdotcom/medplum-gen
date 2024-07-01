OPENAPI_GENERATOR_BIN ?= openapi-generator-cli
SPEC_FILE ?= specs/$(SPEC_VERSION).json
SPEC_VERSION ?= $(shell cat version | sed 's/\./_/g')
GO_BUILD_DIR ?= ./build/go
DATE ?= $(shell date)

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
	mkdir -p /usr/local/bin && \
    curl https://raw.githubusercontent.com/OpenAPITools/openapi-generator/master/bin/utils/openapi-generator-cli.sh > /usr/local/bin/openapi-generator-cli && \
    chmod u+x /usr/local/bin/openapi-generator-cli

### Test

.PHONY: validate
validate: description = Validate spec file
validate:
	$(OPENAPI_GENERATOR_BIN) validate -i $(SPEC_FILE)

### Build

.PHONY: generate
generate: description = Generate ALL code
generate: generate/go

.PHONY: generate/go
generate/go: description = Generate Go code from Medplum OpenAPI spec
generate/go: clean/go
	mkdir -p $(GO_BUILD_DIR)
	$(OPENAPI_GENERATOR_BIN) generate \
	--skip-validate-spec \
	--package-name medplum \
	--git-user-id superpowerdotcom \
	--git-repo-id medplum-gen/build/go \
	--global-property models,supportingFiles,apis=false,modelDocs=false \
	-i $(SPEC_FILE) \
	-g go -o $(GO_BUILD_DIR) || (echo ">> ERROR: Failed to generate Go code" && exit 1)

	# Remove unnecessary files/docs/etc.
	$(RM) -r $(GO_BUILD_DIR)/docs \
		$(GO_BUILD_DIR)/api \
		$(GO_BUILD_DIR)/.gitignore \
		$(GO_BUILD_DIR)/.openapi-generator \
		$(GO_BUILD_DIR)/.openapi-generator-ignore \
		$(GO_BUILD_DIR)/.gitignore \
		$(GO_BUILD_DIR)/git_push.sh \
		$(GO_BUILD_DIR)/.travis.yml \
		$(GO_BUILD_DIR)/client.go \
		$(GO_BUILD_DIR)/configuration.go

	# Cleanup README
	sed -n '/## Installation/q;p' build/go/README.md > $(GO_BUILD_DIR)/README.md.tmp
	echo "<sub>\`make generate/go\` ran at \`$(DATE)\`.</sub>" >> $(GO_BUILD_DIR)/README.md.tmp
	mv -f $(GO_BUILD_DIR)/README.md.tmp $(GO_BUILD_DIR)/README.md

.PHONY: clean
clean: description = Remove ALL build artifacts
clean: clean/go

.PHONY: clean/go
clean/go: description = Remove Go build artifacts
clean/go:
	$(RM) -r $(GO_BUILD_DIR)


