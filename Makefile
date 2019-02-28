LINT_FLAGS :=--disable-all --enable=vet --enable=vetshadow --enable=golint --enable=ineffassign --enable=goconst --enable=gofmt --enable=goimports --deadline=120s
LINTER_EXE := gometalinter.v2
LINTER := $(GOPATH)/bin/$(LINTER_EXE)
TESTFLAGS := -v -cover

GO111MODULE := on
all: $(LINTER) deps test lint build

$(LINTER):
	GO111MODULE=off && go get -u gopkg.in/alecthomas/$(LINTER_EXE)
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
