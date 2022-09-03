SOURCES=lab.draklowell.net/routine-runtime/wrapper/
PYTHON=python3
VERSION=$(shell cat VERSION)

export NAME=routine-runtime-$(VERSION)
export CGO_ENABLED=1

ifndef OS
export OS=linux
endif
ifndef ARCH
export ARCH=amd64
endif

ifeq ($(OS),windows)
export GOOS=windows
EXTENSION=dll
else ifeq ($(OS),linux)
export GOOS=linux
EXTENSION=so
else ifeq ($(OS),darwin)
export GOOS=darwin
EXTENSION=so
endif
export GOARCH=$(ARCH)

.PHONY: build

build:
	go build $(GOFLAGS) -buildmode=c-shared -o build/$(NAME)-$(OS)-$(ARCH).$(EXTENSION) $(SOURCES)
	$(PYTHON) make_header.py

build-full:
	OS=linux ARCH=amd64 make build

clean:
	go clean
	rm -rf __pycache__
	rm -rf build/

summary:
	cloc common/ internal/ loader/ rrt/ wrapper/

tag:
	@echo "Add tag: '$(VERSION)'"
	git tag v$(VERSIONVERSION)
