GOARCH := amd64

GOLANG_DOCKER_IMAGE := golang:1.11

PACKAGE_PATH := github.com/fucangyu/mblog
WORKDIR := /go/src/$(PACKAGE_PATH)
OUTPUT_PATH := $(WORKDIR)/mblog
IMAGE_VERSION := fucangyu:mblog

build-linux:
	    docker run --rm -v $(PWD):$(WORKDIR) -e GOOS=linux -e GOARCH=$(GOARCH) -w $(WORKDIR) $(GOLANG_DOCKER_IMAGE) ./build.sh

build-darwin:
	    docker run --rm -v $(PWD):$(WORKDIR) -e GOOS=darwin -e GOARCH=$(GOARCH) -w $(WORKDIR) $(GOLANG_DOCKER_IMAGE) $(WORKDIR)/build.sh

image-build:
	    docker build -t $(IMAGE_VERSION) .

test:
	    docker run --rm -v $(PWD):$(WORKDIR) $(GOLANG_DOCKER_IMAGE) go test -v $(PACKAGE_PATH)/...

