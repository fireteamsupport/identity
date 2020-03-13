.PHONY: all build clean docker docker-build docker-push docker-test

GOBUILD = go build
GORUN = go run
DOCKER = docker
APPNAME = identity

all: clean build


clean:
	rm -rf ./bin

build: clean
	$(GOBUILD) -o bin/$(APPNAME) cmd/$(APPNAME)/*.go

docker-test:
	test $(DOCKERREPO)

docker-build: docker-test
	$(DOCKER) build -t $(DOCKERREPO)

docker-push:
	$(DOCKER) push $(DOCKERREPO)

docker: docker-build docker-push
