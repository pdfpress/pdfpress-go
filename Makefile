COMMIT_ID := $(shell git -C "$(MD)" rev-parse HEAD | head -c8)
GO           ?= go
GOFMT        ?= $(GO)fmt
GOOPTS       ?=

all: test lint vet

build:
	$(GO) build $(GOOPTS) ./...

test:
	$(GO) test -race ./...

vet:
	$(GO) vet ./...

lint:
	gometalinter --disable-all --enable=errcheck --enable=golint --enable=ineffassign --enable=goimports --enable=lll --line-length=120 --enable=varcheck --enable=interfacer --enable=unconvert --enable=structcheck --enable=megacheck ./..
