BUILD_DIR ?= .build
PROJECT_NAME ?= "app"

include scripts/makefiles/third_party/pasdam/makefiles/docker.mk
include scripts/makefiles/third_party/pasdam/makefiles/go.mk
include scripts/makefiles/third_party/pasdam/makefiles/go.mod.mk
include scripts/makefiles/third_party/pasdam/makefiles/help.mk

.DEFAULT_GOAL := help

GO_MAIN_DIR := ./cmd/cli

## clean: Remove all artifacts
.PHONY: clean
clean: go-clean docker-clean

## gitlab-ci-test: Run the stages locally to verify that they execute correctly
.PHONY: gitlab-ci-test
gitlab-ci-test:
	@gitlab-runner exec docker inspect
	@gitlab-runner exec docker build
