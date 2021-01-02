.PHONY: all
all: test build

# BINARY
BINARY_DIR=bin
JAILER=jailer
JAILER_COMPOSE=jailer-compose

# FLAGS
HEAD = ${shell git rev-parse HEAD}
BUILD_DATE = ${shell date -Iminutes}
BUILD_OS = ${shell uname}
LINKER_FLAGS = -ldflags "-X main.commit=${HEAD} -X main.date=${BUILD_DATE} -X main.hostOS=${BUILD_OS}"
UNAME := ${shell uname}

# COMMANDS
.PHONY: deps
deps:
	go get -u ./...

${BINARY_DIR}:
	mkdir -p ${BINARY_DIR}

.PHONY: build
build: ${BINARY_DIR}
	go build -o ${BINARY_DIR}/${JAILER} ${LINKER_FLAGS} ./cmd/${JAILER}
	go build -o ${BINARY_DIR}/${JAILER_COMPOSE} ${LINKER_FLAGS} ./cmd/${JAILER_COMPOSE}

.PHONY: fbsd-arm64
fbsd-arm64: ${BINARY_DIR}
	GOOS=freebsd GOARCH=arm64 go build -o ${BINARY_DIR}/${JAILER}-fbsd_arm64 ${LINKER_FLAGS} ./cmd/${JAILER}
	GOOS=freebsd GOARCH=arm64 go build -o ${BINARY_DIR}/${JAILER_COMPOSE}-fbsd_arm64 ${LINKER_FLAGS} ./cmd/${JAILER_COMPOSE}

.PHONY: fbsd-amd64
fbsd-amd64: ${BINARY_DIR}
	GOOS=freebsd GOARCH=amd64 go build -o ${BINARY_DIR}/${JAILER}-fbsd_amd64 ${LINKER_FLAGS} ./cmd/${JAILER}
	GOOS=freebsd GOARCH=amd64 go build -o ${BINARY_DIR}/${JAILER_COMPOSE}-fbsd_amd64 ${LINKER_FLAGS} ./cmd/${JAILER_COMPOSE}

.PHONY: test
test:
	go test -race ./...

.PHONY: coverage
coverage:
	go test -race -coverprofile coverage.out ./...

.PHONY: report
report: coverage
	go tool cover -html=coverage.out

.PHONY: clean
clean:
	go clean
	go clean -testcache
	rm -rf ${BINARY_DIR}
	rm -f coverage.out

.PHONY: format
format:
	goimports -w .

.PHONY: stats
stats:
	cloc --by-file .