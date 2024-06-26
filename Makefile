IMG ?= arbctl:latest

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

ARCH=$(shell uname -m)
ifeq (x86_64, $(ARCH))
	ARCH=amd64
else ifeq (aarch64, $(ARCH))
	ARCH=arm64
endif
OS=$(shell uname -s)
os=$(shell uname -s | tr '[:upper:]' '[:lower:]')
GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

# Setting SHELL to bash allows bash commands to be executed by recipes.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

LOCALBIN ?= bin
ABS_LOCALBIN=$(shell pwd)/$(LOCALBIN)
export PATH := $(ABS_LOCALBIN):$(PATH)

ENSURE_DEPS=$(shell pwd)/scripts/ensure-deps.sh
## Tool Binaries
GOLANGCI_LINT ?= $(LOCALBIN)/golangci-lint
LINTER ?= $(GOLANGCI_LINT)
ABIGEN ?= $(LOCALBIN)/abigen

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

.PHONY: all
all: build

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: lint
lint: $(LINTER) ## Lint code
	$(LINTER) run -v

##@ Build

.PHONY: build
build: ## Build binary.
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

# If you wish built the manager image targeting other platforms you can use the --platform flag.
# (i.e. docker build --platform linux/arm64 ). However, you must enable docker buildKit for it.
# More info: https://docs.docker.com/develop/develop-images/build_enhancements/
.PHONY: image
image: ## Build docker image with local proxy setting
	docker build . --progress=plain -t ${IMG} --build-arg GOPROXY=$(shell go env GOPROXY)

##@ Build Dependencies
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

ENSURE_DEPS: $(LOCALBIN) $(ENSURE_DEPS)
	scripts/ensure-deps.sh

$(LINTER): ENSURE_DEPS
$(ABIGEN): ENSURE_DEPS

.PHONY: deps
deps: $(LOCALBIN) $(LINTER) $(ABIGEN)  ## Install binary dependencies

##@ Code Generation
.PHONY: generate
generate: deps ## Generate go files by go generate
	go mod tidy
	go generate ./...

.PHONY: generate-all
generate-all: ## Generate all files
	make config
