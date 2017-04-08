# Spot Cache

```
                       _____                          ______       
 ________________________  /_      ____________ _________  /______ 
 ___  ___/__  __ \  __ \  __/_______  ___/  __ `/  ___/_  __ \  _ \
 __(__  )__  /_/ / /_/ / /_ _/_____/ /__ / /_/ // /__ _  / / /  __/
  /____/ _  .___/\____/\__/        \___/ \__,_/ \___/ /_/ /_/\___/ 
         /_/                                                       
```

_A fast, light-weight cache service written in golang and backed by leveldb._

# Overview

Spot cache is a fast, highly available cache service written in golang and backed by leveldb.  Server connections are via TCP sockets with cluster/replication to support multiple machines.  It can also be used as a single instance in a small network of machines.

Socket protocol is asynchronous request/replay that uses a thin envelope to match the correct response to it's request.


## Installation

**Note: this project is in the early stages and doesn't expect to go live until then end of Q2-2017.**

## Server

### Highlights

* level-db backed
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
| expire   | key, seconds | err  |
| ttl      | key   | seconds, err  |
| subscribe | name | |
| unsubscribe | name | |
| publish  | name, message | |
| ping     |            | pong |
| status   |            | data |
| shutdown |            | err  |


### Socket Message Data

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

## Contributors

_Actively seeking contributors for testing, client implementations, etc._

Here are the rules:

* RainCitySoftware develops all software using TDD.  All submitted code must include a complete set of unit tests, and if appropriate, functional tests.
* We follow golang's idioms and best practices.  All code must be formatted using go-fmt.
* Please submit pull requests. After code reviews and possible modifications your code will be merged.
* Code of conduct: you known the drill. The ACM's [Code of Ethics and Professional Conduct](https://www.acm.org/about-acm/acm-code-of-ethics-and-professional-conduct) says it all--don't be a dick.

## Go Report Card



## License

Apache 2.0

_This project was inspired in part by [Suryandaru Triandana](https://github.com/syndtr/goleveldb)'s excellent port of leveldb to golang._

###### darryl.west | 2017-03-26 | Version 0.90.106-alpha
