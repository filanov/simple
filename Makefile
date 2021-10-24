__root=$(shell pwd)

SERVICE=simplesrv
REGISTRY=registry.local:5000

.PHONY: build

build:
	mkdir -p build
	go build -o build/simplesrv cmd/main.go

clear:
	rm -rf build

container-build: build
	docker build -f Dockerfile.simplesrv . -t $(SERVICE)

generate:
	rm -rf models restapi
	docker run --user $(shell id -u):$(shell id -g) -v ${__root}:${__root}:rw,Z -v /etc/passwd:/etc/passwd -w ${__root} \
        quay.io/goswagger/swagger generate server --template=stratoscale -f ${__root}/swagger.yaml

container-push: container-build
	docker tag $(SERVICE) $(REGISTRY)/$(SERVICE)
	docker push $(REGISTRY)/$(SERVICE)