.PHONY: all build clean docker docker-build docker-push docker-test

GOBUILD = go build
GOGENERATE = go generate
GORUN = go run
DOCKER = docker
APPNAME = identity

all: clean build


clean:
	rm -rf ./bin

generate:
	if [ -f cmd/$(APPNAME)/wire_gen.go ]; then $(GOGENERATE) cmd/$(APPNAME)/wire.go; else wire cmd/$(APPNAME)/wire.go; fi;

build: clean generate
	$(GOBUILD) -o bin/$(APPNAME) cmd/$(APPNAME)/main.go cmd/$(APPNAME)/wire_gen.go

docker-test:
	test $(DOCKERREPO)

docker-build: docker-test
	$(DOCKER) build -t $(DOCKERREPO)

docker-push:
	$(DOCKER) push $(DOCKERREPO)

docker: docker-build docker-push
