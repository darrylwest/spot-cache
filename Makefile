SHELL := /bin/bash
export GOPATH:=$(HOME)/.gopath:$(PWD)

build: 
	@[ -d bin ] || mkdir bin
	( go build -o bin/spotcached src/main.go )

build-linux:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o docker/spotcached src/main.go

install-deps:
	go get -u github.com/golang/lint/golint
	go get github.com/oklog/ulid
	go get -u github.com/darrylwest/cassava-logger/logger
	go get -u github.com/darrylwest/go-unique/unique
	go get github.com/franela/goblin
	go get github.com/syndtr/goleveldb/leveldb
	go get github.com/boltdb/bolt/...

format:
	( gofmt -s -w src/*.go src/spotcache/*.go test/*/*.go examples/*.go tools/*.go )

lint:
	@( golint src/... && golint test/... && golint tools/... && golint examples )

qtest:
	@( [ -d $(HOME)/.spotcache ] || mkdir $(HOME)/.spotcache )
	@( cd test/unit && go test -cover )

test:
	@( [ -d $(HOME)/.spotcache ] || mkdir $(HOME)/.spotcache )
	@( go vet src/spotcache/*.go && go vet src/spotclient/*.go && go vet src/*.go && cd test/unit && go test -cover )
	@( make lint )

watch:
	./watcher.js

run:
	( go run src/main.go --env=development )

start:
	( make build )
	./bin/spotcached &

status:
	@( echo "implement a socket client that will request status..." )

ping:
	@( echo "implement a socket client that will request a ping..." )

shutdown:
	@( echo "implement a socket client that will request a shutdown..." )

edit:
	vi -O3 src/*/*.go test/unit/*.go src/*.go

client:
	go run examples/ping-client.go

.PHONY: format
.PHONY: test
.PHONY: qtest
.PHONY: watch
.PHONY: run
