GOPATH := $(shell go env GOPATH)

.PHONY: setup
setup:
	$(GOPATH)/bin/wire
	go generate ./...
