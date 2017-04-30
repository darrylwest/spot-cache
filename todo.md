# To Do List

## A list of Stuff that needs to be done...

* dockerize service with scrach container
* finish client tests and implementations
* create random data source to create random put/get/has tests
* swap tcp sockets with unix sockets and benchmark to see if it beats the current benchmarks
* replace fixed buffer with dynamic to read request bytes
* examine and compare socket logic to redis and mongo
* refactor to read a conf file and put in /etc/spotcache/spotcahce.conf
* create a REPL for the cache database 

## Benchmarks

### ping only

87 - 95 ms / 1000 tx for 10 clients
400 - 750 ms / 5000 tx for 10 clients

