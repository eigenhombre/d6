.PHONY: all install

PROG=d6

all: ${PROG}

${PROG}: *.go go.mod
	go build .

install:
	go install
