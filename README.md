# Spot Cache

```
                       _____                          ______       
 ________________________  /_      ____________ _________  /______ 
 ___  ___/__  __ \  __ \  __/_______  ___/  __ `/  ___/_  __ \  _ \
 __(__  )__  /_/ / /_/ / /_ _/_____/ /__ / /_/ // /__ _  / / /  __/
  /____/ _  .___/\____/\__/        \___/ \__,_/ \___/ /_/ /_/\___/ 
         /_/                                                       
```

_A lightning fast, light-weight cache service written in golang and backed by leveldb, Boltdb, or RethinkDb._

[![build](https://travis-ci.org/darrylwest/spot-cache.svg?branch=master)](https://travis-ci.org/darrylwest/spot-cache/)
[![reportcard](https://goreportcard.com/badge/github.com/darrylwest/spot-cache)](https://goreportcard.com/report/github.com/darrylwest/spot-cache)

# Overview

Spot cache is a fast, highly available cache service written in golang with backing from leveldb or bolt.  Server connections are via TCP sockets with cluster/replication to support multiple machines.  It can also be used as a single instance in a small network of machines.

Socket protocol is asynchronous request/replay that uses a thin envelope to match the correct response to it's request.


## Installation

**Note: this project is in the early stages and doesn't expect to go live until the middle of Q3-2017.**

## Server

### Highlights

* [level-db](https://github.com/syndtr/goleveldb) or [boltdb](https://github.com/boltdb/bolt) backed
* socket API (request/response)
* cluster/replication support
* written in golang
* asynchronous request/response (zeromq inspired)
* Dockerfile to enable containerization
* lightweight enough to support a small cluster inside a single container

### Cluster Configuration



## Client Implementations

* golang, nodejs, python implementations
* minimal API

## Examples

_coming soon..._

## Unit / Integration / Benchmark Tests

Unit tests are in the test folder.  Run them with 'make test'.  Integration tests exercise the available clients by generating random data to store and retrieve.  Stress tests are written in go and provide data for the benchmarks below (TBD)...

## API

### Command Operations

| func     | params     | response         | description |
|----------|------------|------------------|-------------|
| get      | key        | data, err        | get the data for a given key |
| put      | key, value | err              | put data for a key |
| del      | key        | err              | delete data for a key |
| has      | key        | t/f, err         | return true if key exists
| batch    | key, value | data, err        | create/execute a batch of ops |
| expire   | key, seconds | err  | set the expiration in seconds |
| ttl      | key   | seconds, err  | return the expiration in seconds |
| subscribe | name | | subscribe to a channel |
| unsubscribe | name | | unsubscribe from a channel |
| publish  | name, message | | publish to a channel |
| keys     | query      | data, err | return a set of keys |
| backup   |            | err  | do a snap backup of the current database |
| clear    |            | err  | clear all data from the database |
| ping     |            | pong | send a pong response |
| status   |            | data | return the status of the cache |
| shutdown |            | err  | shutdown the cache service |


### Socket Message Data

#### Request Message Format

The socket channel is binary based using little endian.  For go-lang the encoding/binary package provides an easy way to create a custom client.  

The message format is as follows:

| description | bytes | examples | comments
|-------------|------|-----|---|
| id   | 26 | 01BB20AAGCDCW60MZZNP7F7T8H | a [standard ulid](https://github.com/alizain/ulid) works best
| session | 12 | bbeof787vpkw | session id granted by server (nano seconds at base 36)
| op   | 1  | see op codes | op codes for put, get, del, etc.
| meta size | 2 | 0, 32, 128 | the size of any meta data (can be zero), expire, hmac, etc.
| key size | 2 | 32, 128 | the size in bytes of the data key
| data size | 4 | 256, 64,000 | the size in bytes of the data value (can be zero)
| meta data | n | expire:60 | any meta data, e.g., expire, hmac, app-key, etc
| data key  | n | mykey:2344 | specified by the key size
| data value | n | my value for this key | 

#### Response Message Format

| description | bytes | examples | comments
|-------------|------|-----|---|
| id   | 26 | 01BB20AAGCDCW60MZZNP7F7T8H | the request's id
| session | 12 | bbeof787vpkw | the request's session id
| op   | 1  | see op codes | the request's op code |
| meta size | 2 | any returned meta data |
| data size | 4 | response data size |
| meta data | n | |
| data value | n | the response data |

## Contributors

_This project is actively seeking contributors for testing, client implementations, etc._

Here are the rules:

* RainCitySoftware develops all software using TDD.  All submitted code must include a complete set of unit tests, and if appropriate, functional tests.
* We follow golang's idioms and best practices.  All code must be formatted using go-fmt.
* Please submit pull requests. After code reviews and possible modifications your code will be merged.
* Code of conduct: you known the drill. The ACM's [Code of Ethics and Professional Conduct](https://www.acm.org/about-acm/acm-code-of-ethics-and-professional-conduct) says it all--don't be a dick.


## License

Apache 2.0

_This project was inspired in part by [Suryandaru Triandana](https://github.com/syndtr/goleveldb)'s excellent port of leveldb to golang._

###### Copyright © 2014-2017, Rain City Software | darryl.west | Version 0.90.116
