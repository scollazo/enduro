PROJECT := enduro-sdps
MAKEDIR := hack/make
SHELL   := /bin/bash

.DEFAULT_GOAL := help
.PHONY: *

DBG_MAKEFILE ?=
ifeq ($(DBG_MAKEFILE),1)
    $(warning ***** starting Makefile for goal(s) "$(MAKECMDGOALS)")
    $(warning ***** $(shell date))
else
    # If we're not debugging the Makefile, don't echo recipes.
    MAKEFLAGS += -s
endif

include hack/make/bootstrap.mk
include hack/make/dep_ent.mk
include hack/make/dep_goa.mk
include hack/make/dep_golangci_lint.mk
include hack/make/dep_gomajor.mk
include hack/make/dep_gosec.mk
include hack/make/dep_gotestsum.mk
include hack/make/dep_migrate.mk
include hack/make/dep_mockgen.mk
include hack/make/dep_shfmt.mk
include hack/make/dep_tparse.mk

define NEWLINE


endef

IGNORED_PACKAGES := \
	github.com/artefactual-sdps/enduro/hack/% \
	github.com/artefactual-sdps/enduro/%/fake \
	github.com/artefactual-sdps/enduro/internal/api/design \
	github.com/artefactual-sdps/enduro/internal/api/gen/% \
	github.com/artefactual-sdps/enduro/internal/persistence/ent/db \
	github.com/artefactual-sdps/enduro/internal/persistence/ent/db/% \
	github.com/artefactual-sdps/enduro/internal/persistence/ent/schema \
	github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/db \
	github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/db/% \
	github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/schema
PACKAGES := $(shell go list ./...)
TEST_PACKAGES := $(filter-out $(IGNORED_PACKAGES),$(PACKAGES))
TEST_IGNORED_PACKAGES := $(filter $(IGNORED_PACKAGES),$(PACKAGES))

export PATH:=$(GOBIN):$(PATH)

env: # @HELP Print Go env variables.
env:
	go env

deps: # @HELP List available module dependency updates.
deps: $(GOMAJOR)
	gomajor list

tparse: # @HELP Run all tests and output a coverage report using tparse.
tparse: $(TPARSE)
	go test -count=1 -json -cover $(TEST_PACKAGES) | tparse -follow -all -notests

test: # @HELP Run all tests and output a summary using gotestsum.
test: TFORMAT ?= short
test: GOTEST_FLAGS ?=
test: COMBINED_FLAGS ?= $(GOTEST_FLAGS) $(TEST_PACKAGES)
test: $(GOTESTSUM)
	gotestsum --format=$(TFORMAT) -- $(COMBINED_FLAGS)

test-race: # @HELP Run all tests with the race detector.
test-race:
	$(MAKE) test GOTEST_FLAGS="-race"

test-ci: # @HELP Run all tests in CI with coverage and the race detector.
test-ci:
	$(MAKE) test GOTEST_FLAGS="-race -coverprofile=covreport -covermode=atomic"

list-tested-packages: # @HELP Print a list of packages being tested.
list-tested-packages:
	$(foreach PACKAGE,$(TEST_PACKAGES),@echo $(PACKAGE)$(NEWLINE))

list-ignored-packages: # @HELP Print a list of packages ignored in testing.
list-ignored-packages:
	$(foreach PACKAGE,$(TEST_IGNORED_PACKAGES),@echo $(PACKAGE)$(NEWLINE))

lint: # @HELP Lint the project Go files with golangci-lint.
lint: OUT_FORMAT ?= colored-line-number
lint: LINT_FLAGS ?= --timeout=5m --fix
lint: $(GOLANGCI_LINT)
	golangci-lint run --out-format $(OUT_FORMAT) $(LINT_FLAGS)

gen-goa: # @HELP Generate Goa assets.
gen-goa: $(GOA)
	goa gen github.com/artefactual-sdps/enduro/internal/api/design -o internal/api
	$(MAKE) gen-goa-json-pretty

gen-goa-json-pretty: HTTP_DIR = "internal/api/gen/http"
gen-goa-json-pretty: JSON_FILES = $(shell find $(HTTP_DIR) -type f -name "*.json" | sort -u)
gen-goa-json-pretty: $(JQ)
	for f in $(JSON_FILES); \
		do (cat "$$f" | jq -S '.' >> "$$f".sorted && mv "$$f".sorted "$$f") \
			&& echo "Formatting $$f with jq" || exit 1; \
	done

gen-dashboard-client: # @HELP Generate the Dashboard web client from the OpenAPI spec.
gen-dashboard-client:
	rm -rf $(CURDIR)/dashboard/src/openapi-generator
	docker container run --rm --user $(shell id -u):$(shell id -g) --volume $(CURDIR):/local openapitools/openapi-generator-cli:v7.1.0 \
		generate \
			--input-spec /local/internal/api/gen/http/openapi3.yaml \
			--generator-name typescript-fetch \
			--output /local/dashboard/src/openapi-generator/ \
			-p "generateAliasAsModel=false" \
			-p "withInterfaces=true" \
			-p "supportsES6=true"
	echo "@@@@ Please, review all warnings generated by openapi-generator-cli above!"

gen-mock: # @HELP Generate mocks.
gen-mock: $(MOCKGEN)
	mockgen -typed -destination=./internal/api/auth/fake/mock_ticket_store.go -package=fake github.com/artefactual-sdps/enduro/internal/api/auth TicketStore
	mockgen -typed -destination=./internal/package_/fake/mock_package_.go -package=fake github.com/artefactual-sdps/enduro/internal/package_ Service
	mockgen -typed -destination=./internal/persistence/fake/mock_persistence.go -package=fake github.com/artefactual-sdps/enduro/internal/persistence Service
	mockgen -typed -destination=./internal/sftp/fake/mock_sftp.go -package=fake github.com/artefactual-sdps/enduro/internal/sftp Client
	mockgen -typed -destination=./internal/storage/fake/mock_storage.go -package=fake github.com/artefactual-sdps/enduro/internal/storage Service
	mockgen -typed -destination=./internal/storage/persistence/fake/mock_persistence.go -package=fake github.com/artefactual-sdps/enduro/internal/storage/persistence Storage
	mockgen -typed -destination=./internal/upload/fake/mock_upload.go -package=fake github.com/artefactual-sdps/enduro/internal/upload Service
	mockgen -typed -destination=./internal/watcher/fake/mock_service.go -package=fake github.com/artefactual-sdps/enduro/internal/watcher Service
	mockgen -typed -destination=./internal/watcher/fake/mock_watcher.go -package=fake github.com/artefactual-sdps/enduro/internal/watcher Watcher

gen-ent: # @HELP Generate Ent assets.
gen-ent: $(ENT)
	ent generate ./internal/persistence/ent/schema \
		--feature sql/versioned-migration \
		--target=./internal/persistence/ent/db
	ent generate ./internal/storage/persistence/ent/schema \
		--feature sql/versioned-migration \
		--target=./internal/storage/persistence/ent/db

shfmt: SHELL_PROGRAMS := $(shell find $(CURDIR)/hack -name *.sh)
shfmt: $(SHFMT) # @HELP Run shfmt to format shell programs in the hack directory.
	shfmt \
		--list \
		--write \
		--diff \
		--simplify \
		--language-dialect=posix \
		--indent=0 \
		--case-indent \
		--space-redirects \
		--func-next-line \
			$(SHELL_PROGRAMS)

gosec: # @HELP Run gosec security scanner.
gosec: $(GOSEC)
	gosec \
		-exclude-dir=dashboard \
		-exclude-dir=hack \
		-exclude-dir=internal/api/gen \
		-exclude-dir=internal/persistence/ent/db \
		-exclude-dir=internal/storage/persistence/ent/db \
		./...

tilt-trigger-internal: # @HELP Restart enduro-internal and wait until ready.
tilt-trigger-internal:
	tilt trigger enduro-internal
	tilt wait --for=condition=Ready uiresource/enduro-internal

help: # @HELP Print this message.
help:
	echo "TARGETS:"
	grep -E '^.*: *# *@HELP' Makefile             \
	    | awk '                                   \
	        BEGIN {FS = ": *# *@HELP"};           \
	        { printf "  %-30s %s\n", $$1, $$2 };  \
	    '
