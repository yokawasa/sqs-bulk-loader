.PHONY: clean all sqs-bulk-loader

.DEFAULT_GOAL := all

TARGETS=sqs-bulk-loader

CUR := $(shell pwd)
OS := $(shell uname)
VERSION := $(shell cat ${CUR}/VERSION)

sqs-bulk-loader:
	golint ${CUR}
	GOOS=linux GOARCH=amd64 GO111MODULE=on go build -ldflags "-X main.buildVersion=${VERSION}" -o ${CUR}/dist/sqs-bulk-loader_linux ${CUR}/src
	GOOS=darwin GOARCH=amd64 GO111MODULE=on go build -ldflags "-X main.buildVersion=${VERSION}" -o ${CUR}/dist/sqs-bulk-loader_darwin ${CUR}/src

all: $(TARGETS)

clean:
	rm -rf dist
