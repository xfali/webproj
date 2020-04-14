export GO111MODULE=on

build: {{.Value.App}}
.PHONY: build

clean:
.PHONY: clean

test:
.PHONY: test

{{.Value.App}}: FORCE
    go build -o $@ ./cmd/{{.Value.App}}

vendor: go.mod go.sum
    go mod vendor

install:
    go install -o {{.Value.App}} ./cmd/{{.Value.App}}
.PHONY: install