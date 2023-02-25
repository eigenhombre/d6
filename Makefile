.PHONY: all install doc docker

PROG=d6

all: ${PROG} doc

${PROG}: *.go go.mod
	go build .

install:
	go install

doc:
	python updatereadme.py

docker:
	docker build -t ${PROG} .
