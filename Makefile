SOURCEDIR=cmd/oinc
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

OUTPUTDIR=_output/local/bin/linux/amd64
BINARY=oinc
VERSION=0.0.3-git
BUILD_TIME=$(shell date +%FT%T%z)
GODEPPATH=$(shell pwd)

LDFLAGS=-ldflags "-X github.com/mfojtik/oinc/core.Version=${VERSION} -X github.com/mfojtik/oinc/core.BuildTime=${BUILD_TIME}"

.DEFAULT_GOAL: install

.PHONY: build
	mkdir -p ${OUTPUTDIR} && \
	GOPATH="${GODEPPATH}/Godeps/_workspace:${GOPATH}" go build ${LDFLAGS} -o ${OUTPUTDIR}/${BINARY} ${SOURCEDIR}/main.go

.PHONY: install
install:
	GOPATH="${GODEPPATH}/Godeps/_workspace:${GOPATH}" go install ${LDFLAGS} ./...

.PHONY: clean
clean:
	[ -f ${GOPATH}/bin/${BINARY} ] && rm -f ${GOPATH}/bin/${BINARY} || true
	[ -d _output ] && rm -rf _output || true
