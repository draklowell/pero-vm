CURRENT_DIR := $(abspath .)

ifdef DEBUG
  DEBUG_FLAG := -e DEBUG
endif

ifndef OS
  OS := linux
endif
ifndef ARCH
  ARCH := amd64
endif

PLATFORM = $(OS)-$(ARCH)
CONTAINER_TAG := pero-$(PLATFORM)
VERSION := `cat VERSION`

all: clean build

build-all:
	OS=all ARCH=all make

ifeq ($(OS),all)
build:
	OS=windows make -B build
	OS=linux make -B build
else ifeq ($(ARCH),all)
  ifeq ($(OS),linux)
  build:
	ARCH=amd64 make -B build
	ARCH=i386 make -B build
	ARCH=arm64 make -B build
	ARCH=arm make -B build
  else ifeq ($(OS),windows)
  build:
	ARCH=amd64 make -B build
	ARCH=i386 make -B build
  endif
else
build:
	mkdir -p build

	docker build --tag $(CONTAINER_TAG) --target builder-$(PLATFORM) .

	@echo Using container $(CONTAINER_TAG)

	docker run --rm \
		--mount type=bind,source=$(CURRENT_DIR)/build,target=/build \
		$(DEBUG_FLAG) \
		$(CONTAINER_TAG)
endif

clean:
	rm -rf build

summary:
	cloc common/ internal/ loader/ pero/ wrapper/ tools/

tag:
	@echo "Add tag: '$(VERSION)'"
	git tag $(VERSION)
