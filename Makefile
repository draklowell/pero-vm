SOURCES := lab.draklowell.net/routine-runtime/wrapper/
PYTHON := python3
VERSION := $(shell cat VERSION)

export NAME := routine-runtime-$(VERSION)
export CGO_ENABLED := 1

ifndef OS
  export OS := linux
endif

ifndef ARCH
  export ARCH := amd64
endif

export GOOS := $(OS)
export GOARCH := $(ARCH)

ifeq ($(OS),windows)
  export CC := x86_64-w64-mingw32-gcc
  EXTENSION := dll
else ifeq ($(OS),linux)
  export CC=gcc
  EXTENSION := so
else ifeq ($(OS),darwin)
  EXTENSION := so
endif

ifneq ($(DEBUG),true)
  GOFLAGS := $(GOFLAGS) -ldflags="-s -w"
endif

.PHONY: build

build:
	go build $(GOFLAGS) -buildmode=c-shared -o build/$(NAME)-$(OS)-$(ARCH).$(EXTENSION) $(SOURCES)
	$(PYTHON) tools/make_header.py

build-full: clean
	OS=linux ARCH=amd64 make build

clean:
	go clean
	rm -rf __pycache__
	rm -rf build/

summary:
	cloc common/ internal/ loader/ rrt/ wrapper/ tools/

tag:
	@echo "Add tag: '$(VERSION)'"
	git tag v$(VERSIONVERSION)
