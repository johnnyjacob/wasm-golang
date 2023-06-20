GOR00T := $(shell go env GOROOT)
wasmmod:
	GOOS=js GOARCH=wasm go build -o policy.wasm

server:
	GOOS=linux GOARCH=amd64 go build -o server/srv server/server.go

deploy: wasmmod server
	cp index.html deploy
	cp /usr/local/go/misc/wasm/wasm_exec.js ./deploy/
	cp policy.wasm deploy/policy.wasm
	cp server/srv deploy/srv

all: deploy

.PHONY: all server deploy
