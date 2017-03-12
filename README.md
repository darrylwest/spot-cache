# Spot Cache

_A fast stand-alone or containerized cache backed by leveldb._

# Overview

A fast cache service written in golang and backed by leveldb.  Server connections are via TCP sockets with cluster/replication to support multiple local machine instances.  It can also be used as a single instance in a small network of machines.

Socket protocol is asynchronous request/replay that uses a thin envelope to match the correct response to it's request.

**Note: this project is in the early stages and doesn't expect to go live until Q2-2017.**

## Server

* level-db backed
* socket API (request/response)
* cluster/replication support
* written in golang
* asynchronous request/response (zeromq inspired)
* Dockerfile to enable containerization


## Client(s)

* golang, nodejs, python implementations
* minimal API

## Examples

(TBD)

## Unit / Integration / Stress Tests

Unit tests are in the test folder.  Run them with 'make test'.  Integration tests exercise the available clients by generating random data to store and retrieve.  Stress tests are written in go and provide data for the benchmarks below (TBD)...

## API

### Command Operations

| func     | params     | response         |
|----------|------------|------------------|
| get      | key        | data, err        |
| put      | key, value | err              |
| del      | key        | err              |
| has      | key        | t/f, err         |
| ttl      | key, value | ttl, err  |
| ping     |            | pong |
| status   |            | data |
| shutdown |            | err  |


### Socket Message Data

The socket channel is binary based using little endian.  For go-lang the encoding/binary package provides an easy way to create a custom client.  

The message format is as follows:

| description | size | examples | comments
|-------------|------|-----|---|
| command id  | 16 bytes | 01BB20AAGCDCW60MZZNP7F7T8H | a [standard ulid](https://github.com/alizain/ulid) works best
| command op  | 2 bytes  | pu, ge, de, ha, pi | first two chars from the full command, put, get, del, etc.
| key size | 2 bytes | 32 | the size in bytes of the data key
| data size | 4 bytes | 256 | the size in bytes of the data value (can be zero)
| data key  | n bytes | mykey:2344 | specified by the key size
| data value | n bytes | my value for this key | 

###### darryl.west | 2017-03-12 | Version 0.90.103-alpha
