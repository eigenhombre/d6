.PHONY: all install doc

PROG=d6

all: ${PROG}

${PROG}: *.go go.mod
	go build .

install:
	go install

doc:
	python updatereadme.py
