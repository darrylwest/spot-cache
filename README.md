# Spot Cache

_A small targeted application cache._

# Overview

## Server

* levelup backed
* socket API (req/resp)
* clusterable
* written in golang
* zero mq (inspired)
* dockerfile included

### Data Types

## Client

* golang, nodejs, python implementations
* minimal API


## API

| func | params     | response         |
|------|------------|------------------|
| get  | key        | data,err         |
| put  | key, value | err              |
| del  | key        | err              |
| has  | key        | t/f, err         |
| ping |            | pong |
| stat |            | data |
| halt |    			| err  |

###### darryl.west | 2017-03-10
