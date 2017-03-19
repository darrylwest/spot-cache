//
// command - parse and response to commands.  fetch and update database/cache;
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-03-11 13:56:46

package spotcache

import (
	"bytes"
	"encoding/binary"
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

// command object to support executions
type Command struct {
	Id    []byte
	Op    []byte
	Key   []byte
	Value []byte
	Resp  []byte
}

// request object as created by the client
type Request struct {
	Id       []byte
	Session  []byte
	Op       []byte
	MetaSize uint16
	KeySize  uint16
	DataSize uint32
	Metadata []byte
	Key      []byte
	Value    []byte
}

// create a new command object
func NewCommander(db *Cache) *Commander {
	cache = db

	return &Commander{}
}

// parse the buffer and return a command structure, or error if parse is not possible
func ParseRequest(buf []byte) (*Command, error) {
	cmd := Command{}
	cmd.Id = []byte("flarb")
	cmd.Op = []byte("ping")

	return &cmd, nil
}

// execute the command as specified in the command structure
func (cmd *Command) Exec() error {
	// need a hash map of functions to support the API
	var err error
	op := string(cmd.Op)

	// TODO: put this into a hash map
	switch op {
	case "put":
		err = cache.Put(cmd.Key, cmd.Value, 0)
		cmd.Resp = ok
	case "get":
		cmd.Resp, err = cache.Get(cmd.Key)
	case "has":
		r, err := cache.Has(cmd.Key)
		if err == nil && r {
			cmd.Resp = yes
		} else {
			cmd.Resp = no
		}
	case "ping":
		cmd.Resp = pong
	case "status":
		cmd.Resp = ok
		log.Info("status: %s", cmd.Resp)
	case "shutdown":
		log.Info("shutdown command received...")
		cmd.Resp = fail
	default:
		msg := fmt.Sprintf("unknown command: %s", op)
		log.Warn(msg)
		err = errors.New(msg)
		cmd.Resp = fail
	}

	return err
}

// a string representation of the command buffer
func (cmd *Command) String() string {
	return fmt.Sprintf("Id:%s,Op:%s,Key:%s,Value:%s,Resp:%s", cmd.Id, cmd.Op, cmd.Key, cmd.Value, cmd.Resp)
}

// TODO : create a command helper object to enable createing put, get, has, etc to share with client applications

// a public helper method to create a full comman structure
func CreateCommand(id, op, key, value []byte) *Command {
	cmd := Command{Id: id, Op: op, Key: key, Value: value}

	return &cmd
}
