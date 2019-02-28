LINT_FLAGS := run --deadline=120s
LINTER_EXE := golangci-lint
LINTER := $(GOPATH)/bin/$(LINTER_EXE)
TESTFLAGS := -v -cover

GO111MODULE := on
all: $(LINTER) deps test lint build

$(LINTER):
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.15.0
	$(LINTER) --install

.PHONY: lint
lint: $(LINTER)
	$(LINTER) $(LINT_FLAGS) ./...

.PHONY: deps
deps:
	go mod download

.PHONY: build
build:
	go build .

.PHONY: test
test:
	go test $(TESTFLAGS) ./...
