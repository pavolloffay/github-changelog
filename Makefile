IMPORT_LOG=import.log
FMT_LOG=fmt.log
PACKAGES := $(shell go list ./cmd/... ./pkg/...)

BINARY_NAME ?= gch
BIN_DIR ?= "build/"
OUTPUT_BINARY ?= "$(BIN_DIR)/$(BINARY_NAME)"

GOOS ?= $(shell uname -s | awk '{print tolower($$0)}')
GO_FLAGS ?= GOOS=$(GOOS) GOARCH=amd64 CGO_ENABLED=0

.PHONY: install-tools
install-tools:
	go get -u github.com/mjibson/esc

.PHONY: install
install: install-tools
	@which dep > /dev/null || curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	dep ensure -vendor-only

.PHONY: templates
templates:
	esc -pkg templates -o pkg/templates/gen_assets.go -ignore .go -prefix templates templates/*

.PHONY: format
format:
	@echo Formatting code...
	@.travis/import-order-cleanup.sh inplace
	@go fmt $(PACKAGES)

.PHONY: check
check:
	@echo Checking...
	@go fmt $(PACKAGES) > $(FMT_LOG)
	@.travis/import-order-cleanup.sh stdout > $(IMPORT_LOG)
	@[ ! -s "$(FMT_LOG)" -a ! -s "$(IMPORT_LOG)" ] || (echo "Go fmt, license check, or import ordering failures, run 'make format'" | cat - $(FMT_LOG) $(IMPORT_LOG) && false)

.PHONY: build
build: templates format
	@echo Building...
	@${GO_FLAGS} go build -o $(OUTPUT_BINARY) ./cmd/main.go

.PHONY: examples
examples:
	@echo define OAUTH_TOKEN env variable before running this
	go run ./cmd/main.go --template /chrono-list.md > ./examples/chrono-list.md
	go run ./cmd/main.go --template /default-labels.md > ./examples/default-labels.md

