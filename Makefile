export GO111MODULE=on

build: starter
.PHONY: build

clean:
.PHONY: clean

test:
.PHONY: test

starter: FORCE
    go build -o $@ ./cmd/starter

vendor: go.mod go.sum
    go mod vendor

install:
    go install -o starter ./cmd/starter
.PHONY: install