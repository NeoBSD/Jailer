.PHONY: all
all: test build

# BINARY
BINARY_DIR=bin
BINARY_NAME=jailer

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
	go build -o ${BINARY_DIR}/${BINARY_NAME} ${LINKER_FLAGS}

.PHONY: cross
cross: linux-amd64 linux-arm linux-arm64 fbsd-amd64 fbsd-arm fbsd-arm64

.PHONY: linux-amd64
linux-amd64: ${BINARY_DIR}
	GOOS=linux GOARCH=amd64 go build -o ${BINARY_DIR}/${BINARY_NAME}-linux_amd64 ${LINKER_FLAGS}

.PHONY: linux-arm
linux-arm: ${BINARY_DIR}
	GOOS=linux GOARCH=arm go build -o ${BINARY_DIR}/${BINARY_NAME}-linux_arm ${LINKER_FLAGS}

.PHONY: linux-arm64
linux-arm64: ${BINARY_DIR}
	GOOS=linux GOARCH=arm64 go build -o ${BINARY_DIR}/${BINARY_NAME}-linux_arm64 ${LINKER_FLAGS}

.PHONY: fbsd-amd64
fbsd-amd64: ${BINARY_DIR}
	GOOS=freebsd GOARCH=amd64 go build -o ${BINARY_DIR}/${BINARY_NAME}-fbsd_amd64 ${LINKER_FLAGS}

.PHONY: fbsd-arm
fbsd-arm: ${BINARY_DIR}
	GOOS=freebsd GOARCH=arm go build -o ${BINARY_DIR}/${BINARY_NAME}-fbsd_arm ${LINKER_FLAGS}

.PHONY: fbsd-arm64
fbsd-arm64: ${BINARY_DIR}
	GOOS=freebsd GOARCH=arm64 go build -o ${BINARY_DIR}/${BINARY_NAME}-fbsd_arm64 ${LINKER_FLAGS}

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