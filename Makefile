# GO ENV
GOCMD=go
GOFMT=goimports
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

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
all: deps test build
deps:
	$(GOGET) -u ./...
	$(GOCMD) mod tidy
build:
	$(GOBUILD) -o $(BINARY_NAME) $(LINKER_FLAGS)
test:
	$(GOTEST) ./...
test-coverage:
	$(GOTEST) -coverprofile cover.out ./... && go tool cover -html=cover.out -o cover.html
clean:
	$(GOCLEAN)
	$(GOCLEAN) -testcache
	rm -f $(BINARY_NAME)
	rm -f cover.out cover.html
format:
	$(GOFMT) -w .
