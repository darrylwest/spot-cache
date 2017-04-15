SHELL := /bin/bash
export GOPATH:=$(HOME)/.gopath:$(PWD)

build: 
	@[ -d bin ] || mkdir bin
	( go build -o bin/spotcache src/main.go )

install-deps:
	go get -u github.com/golang/lint/golint
	go get github.com/oklog/ulid
	go get -u github.com/darrylwest/cassava-logger/logger
	go get github.com/franela/goblin
	go get github.com/syndtr/goleveldb/leveldb

format:
	( gofmt -s -w src/*.go src/spotcache/*.go test/*/*.go examples/*.go )

lint:
	@( golint src/... && golint test/... && golint examples && golint clients/golang )

qtest:
	@( [ -d $(HOME)/.spotcache ] || mkdir $(HOME)/.spotcache )
	@( cd test/unit ; clear ; go test -cover )

test:
	@( [ -d $(HOME)/.spotcache ] || mkdir $(HOME)/.spotcache )
	@( go vet src/spotcache/*.go ; go vet src/*.go ; cd test/unit ; go test -cover )
	@( make lint )

watch:
	./watcher.js

run:
	( go run src/main.go --env=development )

start:
	( ./bin/spotcache & )

status:
	@( echo "implement a socket client that will request status..." )

ping:
	@( echo "implement a socket client that will request a ping..." )

shutdown:
	@( echo "implement a socket client that will request a shutdown..." )

edit:
	vi -O3 src/*/*.go test/unit/*.go src/*.go

client:
	go run examples/test-client.go

.PHONY: format
.PHONY: test
.PHONY: qtest
.PHONY: watch
.PHONY: run
