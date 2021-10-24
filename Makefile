__root=$(shell pwd)

.PHONY: build

build:
	mkdir -p build
	go build -o build/simplesrv cmd/main.go

clear:
	rm -rf build

generate:
	rm -rf models restapi
	podman run -u $(id -u):$(id -u) -v ${__root}:${__root}:rw,Z -v /etc/passwd:/etc/passwd -w ${__root} \
        quay.io/goswagger/swagger generate server --template=stratoscale -f ${__root}/swagger.yaml