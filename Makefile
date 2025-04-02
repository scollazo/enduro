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

define NEWLINE


endef

IGNORED_PACKAGES := \
	github.com/artefactual-sdps/enduro/hack/% \
	github.com/artefactual-sdps/enduro/%/fake \
	github.com/artefactual-sdps/enduro/internal/api/design \
	github.com/artefactual-sdps/enduro/internal/api/gen/% \
	github.com/artefactual-sdps/enduro/internal/enums \
	github.com/artefactual-sdps/enduro/internal/persistence/ent/db \
	github.com/artefactual-sdps/enduro/internal/persistence/ent/db/% \
	github.com/artefactual-sdps/enduro/internal/persistence/ent/schema \
	github.com/artefactual-sdps/enduro/internal/storage/enums \
	github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/db \
	github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/db/% \
	github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/schema
PACKAGES = $(shell go list ./...)
TEST_PACKAGES = $(filter-out $(IGNORED_PACKAGES),$(PACKAGES))
TEST_IGNORED_PACKAGES = $(filter $(IGNORED_PACKAGES),$(PACKAGES))

# Configure bine.
include hack/make/tools.mk
export PATH := $(shell go tool bine path):$(PATH)

atlas-hash: tool-atlas # @HELP Recalculate the migration hashes.
	atlas migrate hash \
		--dir="file://internal/db/migrations" \
		--dir-format="atlas"
	atlas migrate hash \
		--dir="file://internal/storage/persistence/migrations" \
		--dir-format="atlas"

db: # @HELP Opens the MySQL shell connected to the enduro development database.
db:
	mysql -h127.0.0.1 -P3306 -uroot -proot123 enduro

deps: # @HELP List available module dependency updates.
deps: tool-gomajor
	gomajor list

env: # @HELP Print Go env variables.
env:
	go env

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

gen-ent: # @HELP Generate Ent assets.
gen-ent: tool-ent
	ent generate ./internal/persistence/ent/schema \
		--feature sql/versioned-migration \
		--target=./internal/persistence/ent/db
	ent generate ./internal/storage/persistence/ent/schema \
		--feature sql/versioned-migration \
		--target=./internal/storage/persistence/ent/db

gen-enums: # @HELP Generate go-enum assets.
gen-enums: ENUM_FLAGS = --names --template=$(CURDIR)/hack/make/enums.tmpl
gen-enums: tool-go-enum
	go-enum $(ENUM_FLAGS) \
		-f internal/enums/sip_type.go \
		-f internal/enums/sip_status.go \
		-f internal/enums/preprocessing_task_outcome.go \
		-f internal/enums/task_status.go \
		-f internal/enums/workflow_status.go \
		-f internal/enums/workflow_type.go \
		-f internal/storage/enums/aip_status.go \
		-f internal/storage/enums/deletion_request_status.go \
		-f internal/storage/enums/location_purpose.go \
		-f internal/storage/enums/location_source.go \
		-f internal/storage/enums/task_status.go \
		-f internal/storage/enums/workflow_status.go \
		-f internal/storage/enums/workflow_type.go

gen-goa: # @HELP Generate Goa assets.
gen-goa: tool-goa
	goa gen github.com/artefactual-sdps/enduro/internal/api/design -o internal/api
	$(MAKE) gen-goa-json-pretty

gen-goa-json-pretty: HTTP_DIR = "internal/api/gen/http"
gen-goa-json-pretty: JSON_FILES = $(shell find $(HTTP_DIR) -type f -name "*.json" | sort -u)
gen-goa-json-pretty: $(JQ)
	for f in $(JSON_FILES); \
		do (cat "$$f" | jq -S '.' >> "$$f".sorted && mv "$$f".sorted "$$f") \
			&& echo "Formatting $$f with jq" || exit 1; \
	done

gen-mock: # @HELP Generate mocks.
gen-mock: tool-mockgen
	mockgen -typed -destination=./internal/a3m/fake/mock_transfer_service_client.go -package=fake buf.build/gen/go/artefactual/a3m/grpc/go/a3m/api/transferservice/v1beta1/transferservicev1beta1grpc TransferServiceClient
	mockgen -typed -destination=./internal/api/auth/fake/mock_ticket_store.go -package=fake github.com/artefactual-sdps/enduro/internal/api/auth TicketStore
	mockgen -typed -destination=./internal/api/auth/fake/mock_token_verifier.go -package=fake github.com/artefactual-sdps/enduro/internal/api/auth TokenVerifier
	mockgen -typed -destination=./internal/ingest/fake/mock_ingest.go -package=fake github.com/artefactual-sdps/enduro/internal/ingest Service
	mockgen -typed -destination=./internal/persistence/fake/mock_persistence.go -package=fake github.com/artefactual-sdps/enduro/internal/persistence Service
	mockgen -typed -destination=./internal/sftp/fake/mock_sftp.go -package=fake github.com/artefactual-sdps/enduro/internal/sftp Client,AsyncUpload
	mockgen -typed -destination=./internal/storage/fake/mock_client.go -package=fake github.com/artefactual-sdps/enduro/internal/storage Client
	mockgen -typed -destination=./internal/storage/fake/mock_storage.go -package=fake github.com/artefactual-sdps/enduro/internal/storage Service
	mockgen -typed -destination=./internal/storage/persistence/fake/mock_persistence.go -package=fake github.com/artefactual-sdps/enduro/internal/storage/persistence Storage
	mockgen -typed -destination=./internal/watcher/fake/mock_service.go -package=fake github.com/artefactual-sdps/enduro/internal/watcher Service
	mockgen -typed -destination=./internal/watcher/fake/mock_watcher.go -package=fake github.com/artefactual-sdps/enduro/internal/watcher Watcher

gosec: # @HELP Run gosec security scanner.
gosec: GOSEC_VERBOSITY ?= "-terse"
gosec: tool-gosec
	gosec \
		$(GOSEC_VERBOSITY) \
		-exclude-dir=dashboard \
		-exclude-dir=hack \
		-exclude-dir=internal/api/gen \
		-exclude-dir=internal/persistence/ent/db \
		-exclude-dir=internal/storage/persistence/ent/db \
		./...

help: # @HELP Print this message.
help:
	echo "TARGETS:"
	grep -E '^.*: *# *@HELP' Makefile             \
	    | awk '                                   \
	        BEGIN {FS = ": *# *@HELP"};           \
	        { printf "  %-30s %s\n", $$1, $$2 };  \
	    '

fmt: # @HELP Format the project Go files with golangci-lint.
fmt: FMT_FLAGS ?=
fmt: tool-golangci-lint
	golangci-lint fmt $(FMT_FLAGS)

lint: # @HELP Lint the project Go files with golangci-lint.
lint: LINT_FLAGS ?= --timeout=5m --fix --output.text.colors
lint: tool-golangci-lint
	golangci-lint run $(LINT_FLAGS)

list-tested-packages: # @HELP Print a list of packages being tested.
list-tested-packages:
	$(foreach PACKAGE,$(TEST_PACKAGES),@echo $(PACKAGE)$(NEWLINE))

list-ignored-packages: # @HELP Print a list of packages ignored in testing.
list-ignored-packages:
	$(foreach PACKAGE,$(TEST_IGNORED_PACKAGES),@echo $(PACKAGE)$(NEWLINE))

mod-tidy: # @HELP Synchronize mod files with current code dependencies.
mod-tidy:
	go mod download
	go mod tidy

mod-tidy-check: # @HELP Check that mod files are tidy.
	go mod tidy -diff
	cd hack/pulumi && go mod tidy -diff

pre-commit: # @HELP Check that code is ready to commit.
pre-commit:
	$(MAKE) -j \
		gosec GOSEC_VERBOSITY="-quiet"\
		lint \
		mod-tidy-check \
		shfmt \
		test-race \
		workflowcheck \

shfmt: SHELL_PROGRAMS := $(shell find $(CURDIR)/hack -name *.sh)
shfmt: tool-shfmt # @HELP Run shfmt to format shell programs in the hack directory.
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

test: # @HELP Run all tests and output a summary using gotestsum.
test: TFORMAT ?= short
test: GOTEST_FLAGS ?=
test: COMBINED_FLAGS ?= $(GOTEST_FLAGS) $(TEST_PACKAGES)
test: tool-gotestsum
	gotestsum --format=$(TFORMAT) -- $(COMBINED_FLAGS)

test-ci: # @HELP Run all tests in CI with coverage and the race detector.
test-ci:
	$(MAKE) test GOTEST_FLAGS="-race -coverprofile=covreport -covermode=atomic"

test-race: # @HELP Run all tests with the race detector.
test-race:
	$(MAKE) test GOTEST_FLAGS="-race"

tilt-am-knownhosts:  # @HELP Update Archivematica Storage Service known_hosts.
tilt-am-knownhosts: HOST ?= host.k3d.internal
tilt-am-knownhosts: PORT ?= 12322
tilt-am-knownhosts:
	ssh-keyscan -H -p $(PORT) $(HOST) > hack/kube/overlays/dev-am/.known_hosts.secret
	tilt trigger "(Tiltfile)"
	tilt wait --for=condition=Ready "uiresource/(Tiltfile)"
	tilt trigger enduro-am
	tilt wait --for=condition=Ready uiresource/enduro-am

tools: # @HELP Install all tools.
tools: $(addprefix tool-, $(TOOLS))

tparse: # @HELP Run all tests and output a coverage report using tparse.
tparse: tool-tparse
	go test -count=1 -json -cover $(TEST_PACKAGES) | tparse -follow -all -notests

upload-sample-transfer: # @HELP Upload sample transfer (small.zip).
upload-sample-transfer: ADDRESS ?= localhost:9002
upload-sample-transfer:
	curl \
		-F "file=@$(CURDIR)/internal/testdata/zipped_transfer/small.zip" \
		http://$(ADDRESS)/ingest/sips/upload

workflowcheck: # @HELP Detect non-determinism in workflow functions.
workflowcheck: tool-workflowcheck
	workflowcheck ./...
