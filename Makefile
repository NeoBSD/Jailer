.PHONY: all
all: deps test build

# BINARY
BINARY_DIR=bin
BINARY_NAME=$(BINARY_DIR)/jailer

# FLAGS
HEAD = $(shell git rev-parse HEAD)
BUILD_DATE = $(shell date -Iminutes)
BUILD_OS = $(shell uname)
LINKER_FLAGS = -ldflags "-X main.commit=$(HEAD) -X main.date=$(BUILD_DATE) -X main.hostOS=$(BUILD_OS)"
UNAME := $(shell uname)

# COMMANDS
.PHONY: deps
deps:
	go get -u ./...
	go mod tidy

.PHONY: build
build:
	go build -o $(BINARY_NAME) $(LINKER_FLAGS)

.PHONY: test
test:
	go test ./...

.PHONY: coverage
coverage:
	go test -coverprofile coverage.out ./...

.PHONY: clean
clean:
	go clean
	go clean -testcache
	rm -f $(BINARY_NAME)
	rm -f coverage.out

.PHONY: format
format:
	goimports -w .
