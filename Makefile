.PHONY: clean all sqs-bench

.DEFAULT_GOAL := all

TARGETS=sqs-bench

CUR := $(shell pwd)
OS := $(shell uname)
VERSION := $(shell cat ${CUR}/VERSION)

sqs-bench:
	golint ${CUR}
	GOOS=linux GOARCH=amd64 GO111MODULE=on go build -ldflags "-X main.buildVersion=${VERSION}" -o ${CUR}/dist/sqs-bench_linux ${CUR}/src
	GOOS=darwin GOARCH=amd64 GO111MODULE=on go build -ldflags "-X main.buildVersion=${VERSION}" -o ${CUR}/dist/sqs-bench_darwin ${CUR}/src

all: $(TARGETS)

clean:
	rm -rf dist
