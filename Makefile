GO:= $(shell which go)
GLIDE	:= $(shell which glide)
INSTALL_DIR ?= /usr/local/bin
RELEASE_DIR ?= release
APP ?= $(shell basename `pwd`)

.PHONY: go\:build
## Build binary
go\:build: $(GO)
	$(call assert-set,GO)
	$(GO) build -o $(RELEASE_DIR)/$(APP)

.PHONY: go\:build-all
## Build binary for all platforms
go\:build-all: $(GO)
	$(call assert-set,GO)
ifeq ($(RELEASE_ARCH),)
	gox -output "${RELEASE_DIR}/${APP}_{{.OS}}_{{.Arch}}"
else
	gox -osarch="$(RELEASE_ARCH)" -output "${RELEASE_DIR}/${APP}_{{.OS}}_{{.Arch}}"
endif

.PHONY: go\:deps
## Install dependencies
go\:deps: $(GLIDE)
	$(call assert-set,GLIDE)
	$(GLIDE) install --strip-vendor
	$(GLIDE) update  --strip-vendor

.PHONY: go\:deps-build
## Install dependencies for build
go\:deps-build:
	$(call assert-set,GOPATH)
	mkdir -p  $(GOPATH)/bin
	which $(GLIDE) || (curl https://glide.sh/get | sh)

.PHONY: deps-dev
## Install development dependencies
go\:deps-dev: $(GO)
	$(call assert-set,GO)
	$(GO) get -d -v "github.com/golang/lint"
	$(GO) install -v "github.com/golang/lint/golint"
	$(GO) get -d -v github.com/mitchellh/gox
	$(GO) install -v github.com/mitchellh/gox

## Clean compiled binary
go\:clean:
	rm -rf $(RELEASE_DIR)

## Clean compiled binary and dependency
go\:clean-all: go\:clean
	rm -rf vendor
	rm -rf glide.lock

## Install cli
go\:install: ${RELEASE_DIR}/$(APP) go\:build
	cp $(RELEASE_DIR)/$(APP) $(INSTALL_DIR)

