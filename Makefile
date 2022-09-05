CURRENT_DIR := $(abspath .)

ifdef DEBUG
  DEBUG_FLAG := -e DEBUG
endif

ifndef OS
  OS := all
endif
ifndef ARCH
  ARCH := all
endif

CONTAINER_TAG := rrt-builder

all: clean build

build:
	mkdir build

	docker build --tag $(CONTAINER_TAG) .

	@echo Using container $(CONTAINER_TAG)

	docker run --rm \
		--mount type=bind,source=$(CURRENT_DIR)/build,target=/build \
		-e OS=$(OS) -e ARCH=$(ARCH) $(DEBUG_FLAG) \
		$(CONTAINER_TAG)

clean:
	rm -rf build

summary:
	cloc common/ internal/ loader/ rrt/ wrapper/ tools/

tag:
	@echo "Add tag: '$(VERSION)'"
	git tag v$(VERSIONVERSION)
