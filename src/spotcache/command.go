//
// command - parse and response to commands.  fetch and update database/cache;
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-03-11 13:56:46

package spotcache

import (
	"errors"
	"fmt"
)

// public constants
const (
	RESP_OK    = "ok"
	RESP_FAIL  = "fail"
	RESP_TRUE  = "true"
	RESP_FALSE = "false"
	RESP_PONG  = "pong"
)

// private responses
var (
	cache *Cache
	ok    = []byte(RESP_OK)
	fail  = []byte(RESP_FAIL)
	yes   = []byte(RESP_TRUE)
	no    = []byte(RESP_FALSE)
	pong  = []byte(RESP_PONG)
)

type Commander struct {
}

type Command struct {
	id    []byte
	op    []byte
	key   []byte
	value []byte
	resp  []byte
}

// create a new command object
func NewCommander(db *Cache) *Commander {
	cache = db

	return &Commander{}
}

//
// parse the buffer and return a command structure, or error if parse is not possible
//
func ParseRequest(buf []byte) (*Command, error) {
    cmd := Command{}
    cmd.id = []byte("flarb")
    cmd.op = []byte("ping")

	return &cmd, nil
}

// execute the command as specified in the command structure
func (cmd *Command) Exec() error {
	// need a hash map of functions to support the API
	var err error
	op := string(cmd.op)

	// TODO: put this into a hash map
	switch op {
	case "put":
		err = cache.Put(cmd.key, cmd.value, 0)
		cmd.resp = ok
	case "get":
		cmd.resp, err = cache.Get(cmd.key)
	case "has":
		r, err := cache.Has(cmd.key)
		if err == nil && r {
			cmd.resp = yes
		} else {
			cmd.resp = no
		}
	case "ping":
		cmd.resp = pong
	case "status":
		cmd.resp = ok
		log.Info("status: %s", cmd.resp)
	case "shutdown":
		log.Info("shutdown command received...")
		cmd.resp = fail
	default:
		msg := fmt.Sprintf("unknown command: %s", op)
		log.Warn(msg)
		err = errors.New(msg)
		cmd.resp = fail
	}

	return err
}

// a string representation of the command buffer
func (cmd *Command) String() string {
	return fmt.Sprintf("id:%s,op:%s,key:%s,value:%s,resp:%s", cmd.id, cmd.op, cmd.key, cmd.value, cmd.resp)
}

// public method that returns the command response
func (cmd *Command) GetResp() []byte {
	return cmd.resp
}

// a public helper method to create a full comman structure, less the response
func CreateCommand(id, op, key, value []byte) *Command {
	cmd := Command{id: id, op: op, key: key, value: value}

	return &cmd
}
