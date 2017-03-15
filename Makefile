SHELL := /bin/bash
export GOPATH:=$(HOME)/.gopath:$(PWD)

build: 
	@[ -d bin ] || mkdir bin
	( go build -o bin/spotcache src/main.go )

install-deps:
	go get github.com/oklog/ulid
	go get -u github.com/darrylwest/cassava-logger/logger
	go get github.com/franela/goblin
	go get github.com/syndtr/goleveldb/leveldb

format:
	( gofmt -s -w src/*.go src/spotcache/*.go test/*.go )

qtest:
	@( cd test ; go test )

test:
	@( go vet src/spotcache/*.go ; go vet src/*.go ; cd test ; go test )

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

edit:
	vi -O2 src/*/*.go test/*.go src/*.go

.PHONY: format
.PHONY: test
.PHONY: qtest
.PHONY: watch
.PHONY: run
