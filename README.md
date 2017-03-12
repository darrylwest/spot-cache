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
* zero mq (inspired)
* Dockerfile to enable containerization

### Data Types

## Client

* golang, nodejs, python implementations
* minimal API


## API

| func     | params     | response         |
|----------|------------|------------------|
| get      | key        | data, err        |
| put      | key, value | err              |
| del      | key        | err              |
| has      | key        | t/f, err         |
| ping     |            | pong |
| status   |            | data |
| shutdown |            | err  |

###### darryl.west | 2017-03-12 | Version 0.90.102-alpha
