# Spot Cache

_A fast, light-weight cache service written in golang and backed by leveldb._

# Overview

A fast cache service written in golang and backed by leveldb.  Server connections are via TCP sockets with cluster/replication to support multiple machines.  It can also be used as a single instance in a small network of machines.

Socket protocol is asynchronous request/replay that uses a thin envelope to match the correct response to it's request.

_This project was inspired in part by [Suryandaru Triandana](https://github.com/syndtr/goleveldb)'s excellent port of leveldb to golang._

## Installation

**Note: this project is in the early stages and doesn't expect to go live until Q2-2017.**

## Server

* level-db backed
* socket API (request/response)
* cluster/replication support
* written in golang
* asynchronous request/response (zeromq inspired)
* Dockerfile to enable containerization
* unix socket for status and shutdown

## Client(s)

* golang, nodejs, python implementations
* minimal API

## Examples

_coming soon..._

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

#### Request Messages

Requests are usually sent by the client and include operations to manipulate data to be cached, accessed or expired.  Other ops include the ability to list and subscribe to available events.  There are also requests that may come from the server to one or more clients.

The message request format is as follows:

| description | bytes | examples | comments
|-------------|------|-----|---|
| id   | 26 | 01BB20AAGCDCW60MZZNP7F7T8H | a [standard ulid](https://github.com/alizain/ulid) works best
| session | 12 | bbeof787vpkw | session id granted by server (nano seconds at base 36)
| op   | 2  | pu, ge, de, ha, pi | first two chars from the full command, put, get, del, etc.
| meta size | 2 | 0, 128 | the size in bytes of any associated meta data, and/or hmac or other signature
| key size | 2 | 32, 128 | the size in bytes of the data key
| data size | 4 | 256, 64,000 | the size in bytes of the data value (can be zero)
| metadata | n | anything | this can be zero to any length
| key  | n | mykey:2344 | specified by the key size
| value | n | my value for this key | this could be encrypted, compressed, plain, JSON encoded or whatever

#### Response Messages

After a request has been processed a response is returned that includes the following:

| description | bytes | examples | comments
|-------------|------|-----|---|
| id   | 26 | 01BB20AAGCDCW60MZZNP7F7T8H | a [standard ulid](https://github.com/alizain/ulid) works best
| status | 12 | 0 (ok), > 0 error code | the status of ok/fail is returned here where 0 = success
| meta size | 2 | 0, 128 | the size in bytes of any associated meta data, and/or hmac or other signature
| key size | 2 | 32, 128 | the size in bytes of the data key (may be zero)
| data size | 4 | 256, 64,000 | the size in bytes of the data value (can be zero)
| metadata | n | anything | this can be zero to any length
| key  | n | mykey:2344 | specified by the key size
| value | n | my value, if any | data or status message, etc. | 

#### Event Messages

(TBD) 

Events could include signals to switch to an alternate route, an alert that a set of cached items is about to expire, a request to suspend ops and reconnect in n-seconds (for updates), etc.

## Contributors

_Actively seeking contributors for testing, client implementations, etc._

Here are the rules:

* RainCitySoftware develops all software using TDD.  All submitted code must include a complete set of unit tests, and if appropriate, functional tests.
* We follow golang's idioms and best practices.  All code must be formatted using go-fmt.
* Please submit pull requests. After code reviews and possible modifications your code will be merged.
* Code of conduct: you known the drill. The ACM's [Code of Ethics and Professional Conduct](https://www.acm.org/about-acm/acm-code-of-ethics-and-professional-conduct) says it all--don't be a dick.

## License

Apache 2.0

###### darryl.west | 2017-03-19 | Version 0.90.105-alpha