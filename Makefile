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

all: clean build

build:
	mkdir build

	docker build . 2>&1 | tee /tmp/rrt.build.dockeroutput
	$(eval CONTAINER := $(shell awk 'END {print $$NF}' /tmp/rrt.build.dockeroutput))

	docker run \
		--mount type=bind,source=$(CURRENT_DIR)/build,target=/build \
		-e OS=$(OS) -e ARCH=$(ARCH) $(DEBUG_FLAG) \
		$(CONTAINER)
	docker image rm -f $(CONTAINER)

clean:
	rm -rf build

summary:
	cloc common/ internal/ loader/ rrt/ wrapper/ tools/

tag:
	@echo "Add tag: '$(VERSION)'"
	git tag v$(VERSIONVERSION)
