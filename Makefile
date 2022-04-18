# Copyright 2016 The Kubernetes Authors.
#
# Modifications Copyright 2020 the Velero contributors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# The binary to build (just the basename).
BIN ?= tinkoff_dollars_bot

# This repo's root import path (under GOPATH).
PKG := github.com/WidowGenerator/tinkoff_dollars_bot

# Where to push the docker image.
REGISTRY ?= tinkoff_dollars_bot

# Image name
IMAGE ?= $(REGISTRY)/$(BIN)

# We allow the Dockerfile to be configurable to enable the use of custom Dockerfiles
# that pull base images from different registries.
VELERO_DOCKERFILE ?= Dockerfile
# BUILDER_IMAGE_DOCKERFILE ?= hack/build-image/Dockerfile

# Calculate the realpath of the build-image Dockerfile as we `cd` into the hack/build
# directory before this Dockerfile is used and any relative path will not be valid.
# BUILDER_IMAGE_DOCKERFILE_REALPATH := $(shell realpath $(BUILDER_IMAGE_DOCKERFILE))

# Build image handling. We push a build image for every changed version of
# /hack/build-image/Dockerfile. We tag the dockerfile with the short commit hash
# of the commit that changed it. When determining if there is a build image in
# the registry to use we look for one that matches the current "commit" for the
# Dockerfile else we make one.
# In the case where the Dockerfile for the build image has been overridden using
# the BUILDER_IMAGE_DOCKERFILE variable, we always force a build.

# ifneq "$(origin BUILDER_IMAGE_DOCKERFILE)" "file"
# 	BUILDER_IMAGE_TAG := "custom"
# else
# 	BUILDER_IMAGE_TAG := $(shell git log -1 --pretty=%h $(BUILDER_IMAGE_DOCKERFILE))
# endif

# BUILDER_IMAGE := $(REGISTRY)/build-image:$(BUILDER_IMAGE_TAG)
# BUILDER_IMAGE_CACHED := $(shell docker images -q ${BUILDER_IMAGE} 2>/dev/null )

# HUGO_IMAGE := hugo-builder

# Which architecture to build - see $(ALL_ARCH) for options.
# if the 'local' rule is being run, detect the ARCH from 'go env'
# if it wasn't specified by the caller.
local : ARCH ?= $(shell go env GOOS)-$(shell go env GOARCH)
ARCH ?= linux-amd64

VERSION ?= main

TAG_LATEST ?= false

ifeq ($(TAG_LATEST), true)
	IMAGE_TAGS ?= $(IMAGE):$(VERSION) $(IMAGE):latest
else
	IMAGE_TAGS ?= $(IMAGE):$(VERSION)
endif

ifeq ($(shell docker buildx inspect 2>/dev/null | awk '/Status/ { print $$2 }'), running)
	BUILDX_ENABLED ?= true
else
	BUILDX_ENABLED ?= false
endif

define BUILDX_ERROR
buildx not enabled, refusing to run this recipe
see: https://velero.io/docs/main/build-from-source/#making-images-and-updating-velero for more info
endef

# The version of restic binary to be downloaded
RESTIC_VERSION ?= 0.12.1

CLI_PLATFORMS ?= linux-amd64 linux-arm linux-arm64 darwin-amd64 darwin-arm64 windows-amd64 linux-ppc64le
BUILDX_PLATFORMS ?= $(subst -,/,$(ARCH))
BUILDX_OUTPUT_TYPE ?= docker

# set git sha and tree state
GIT_SHA = $(shell git rev-parse HEAD)
ifneq ($(shell git status --porcelain 2> /dev/null),)
	GIT_TREE_STATE ?= dirty
else
	GIT_TREE_STATE ?= clean
endif

# The default linters used by lint and local-lint
LINTERS ?= "gosec,goconst,gofmt,goimports,unparam"

###
### These variables should not need tweaking.
###

platform_temp = $(subst -, ,$(ARCH))
GOOS = $(word 1, $(platform_temp))
GOARCH = $(word 2, $(platform_temp))
GOPROXY ?= https://proxy.golang.org

local-build:
	GOOS=$(GOOS) \
	GOARCH=$(GOARCH) \
	VERSION=$(VERSION) \
	REGISTRY=$(REGISTRY) \
	PKG=$(PKG) \
	BIN=$(BIN) \
	GIT_SHA=$(GIT_SHA) \
	GIT_TREE_STATE=$(GIT_TREE_STATE) \
	OUTPUT_DIR=$$(pwd)/_output/bin/$(GOOS)/$(GOARCH) \
	./scripts/build.sh

local-test:
	./scripts/test.sh

local-lint:
	./scripts/lint.sh

local-ci: local-lint local-test

format:
	./scripts/format.sh

install-instruments:
	./scripts/install-instruments.sh

install-deps:
	go install ./cmd/main.go

##################################### ENUMS #####################################
ENUM_DIR_PREFIX = internal/db/enums
STANDARD_ENUMS = $(ENUM_DIR_PREFIX)/city_enum.go $(ENUM_DIR_PREFIX)/currency_enum.go $(ENUM_DIR_PREFIX)/bank_enum.go

GOENUM = ./bin/go-enum

$(STANDARD_ENUMS): GO_ENUM_FLAGS=--nocase --marshal --names --ptr

enums: $(STANDARD_ENUMS)

# The generator statement for go enum files.  Files that invalidate the
# enum file: source file, the binary itself, and this file (in case you want to generate with different flags)
%_enum.go: %.go $(GOENUM) Makefile
	$(GOENUM) -f $*.go $(GO_ENUM_FLAGS)

#################################################################################