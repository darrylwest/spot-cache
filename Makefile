SHELL := /bin/bash

build: 
	@[ -d bin ] || mkdir bin
	( . ./.setpath ; go build -o bin/webserver src/main.go )

install-deps:
	go get github.com/pborman/uuid
	go get -u github.com/darrylwest/cassava-logger/logger
	go get github.com/franela/goblin

format:
	( gofmt -s -w src/*.go src/webserver/*.go test/*.go )

qtest:
	@( . ./.setpath ; cd test ; go test )

test:
	@( . ./.setpath ; go vet src/webserver/*.go ; go vet src/*.go ; cd test ; go test )

watch:
	./watcher.js

run:
	( go run src/main.go --env=development )

status:
	@( echo "implement a socket client that will request status..." )

ping:
	@( echo "implement a socket client that will request a ping..." )

shutdown:
	@( echo "implement a socket client that will request a shutdown..." )

.PHONY: format
.PHONY: test
.PHONY: qtest
.PHONY: watch
.PHONY: run
