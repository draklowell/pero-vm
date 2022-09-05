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

CONTAINER := $(shell awk 'END {print $$NF}' build.docker.tmp)

all: clean build

build:
	mkdir build

	docker build . 2>&1 | tee build.docker.tmp

	@echo Using container $(CONTAINER)

	docker run \
		--mount type=bind,source=$(CURRENT_DIR)/build,target=/build \
		-e OS=$(OS) -e ARCH=$(ARCH) $(DEBUG_FLAG) \
		$(CONTAINER)

	docker image rm -f $(CONTAINER)
	rm -f build.docker.tmp

clean:
	rm -rf build
	rm -f build.docker.tmp

summary:
	cloc common/ internal/ loader/ rrt/ wrapper/ tools/

tag:
	@echo "Add tag: '$(VERSION)'"
	git tag v$(VERSIONVERSION)
