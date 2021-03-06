SHELL := /bin/bash
export GOPATH:=$(HOME)/.gopath:$(PWD)

build: 
	@[ -d bin ] || mkdir bin
	( go build -o bin/spotcache-cli src/spotcache-cli.go )
	( go build -o bin/spotcached src/spotcached.go )

build-linux:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o docker/spotcached src/spotcached.go

install-deps:
	go get -u github.com/golang/lint/golint
	go get -u github.com/darrylwest/cassava-logger/logger
	go get -u github.com/darrylwest/go-unique/unique
	go get -u github.com/darrylwest/spot-cache/spotcache
	go get github.com/franela/goblin
	go get github.com/syndtr/goleveldb/leveldb
	go get github.com/boltdb/bolt/...

format:
	( gofmt -s -w src/*.go src/spotcache/*.go src/spotclient/*.go test/*/*.go examples/*.go tools/*.go )

lint:
	@( golint src/... && golint test/... && golint tools/... && golint examples )

qtest:
	@( [ -d $(HOME)/.spotcache ] || mkdir $(HOME)/.spotcache )
	@( cd test/unit && go test -cover )
	@( cd test/client && go test -cover )

test:
	@( [ -d $(HOME)/.spotcache ] || mkdir $(HOME)/.spotcache )
	@( go vet src/spotclient/*.go && cd test/client && go test -cover )
	@( go vet src/spotcache/*.go && go vet src/*.go && cd test/unit && go test -cover )
	@( make lint )

test-client:
	@( go vet src/spotclient/*.go && cd test/client && go test -cover )
	@( make lint )

watch:
	./watcher.js

run:
	( go run src/spotcached.go --env=development )

start:
	( make build )
	./bin/spotcached &

# requires cli
status:
	@( echo "implement a socket client that will request status..." )

ping:
	go run src/spotcache-cli.go

shutdown:
	@( echo "implement a socket client that will request a shutdown..." )

edit:
	vi -O3 src/*/*.go test/unit/*.go src/*.go

# examples...
ping-client:
	go run examples/ping-client.go

writer-client:
	go run examples/writer-client.go

.PHONY: format test qtest watch run
