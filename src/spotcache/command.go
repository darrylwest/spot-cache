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

type CommandOp uint8

const (
	NOOP        = CommandOp(0)
	PUT         = CommandOp(1)
	GET         = CommandOp(2)
	HAS         = CommandOp(3)
	DEL         = CommandOp(4)
	EXPIRE      = CommandOp(10)
	TTL         = CommandOp(11)
	SUBSCRIBE   = CommandOp(20)
	UNSUBSCRIBE = CommandOp(21)
	PUBLISH     = CommandOp(22)
	STATUS      = CommandOp(64)
	PING        = CommandOp(128)
	SHUTDOWN    = CommandOp(255)
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
	Id    IdType
	Op    CommandOp
	Key   []byte
	Value []byte
	Resp  []byte
}

// create a new command object
func NewCommander(db *Cache) *Commander {
	cache = db

	return &Commander{}
}

// parse the buffer and return a command structure, or error if parse is not possible
func ParseRequest(buf []byte) (*Command, error) {
	cmd := Command{}
	// cmd.Id =
	cmd.Op = PING

	return &cmd, nil
}

// execute the command as specified in the command structure
func (cmd *Command) Exec() error {
	// need a hash map of functions to support the API
	var err error

	// TODO: put this into a hash map
	switch cmd.Op {
	case PUT:
		err = cache.Put(cmd.Key, cmd.Value, 0)
		cmd.Resp = ok
	case GET:
		cmd.Resp, err = cache.Get(cmd.Key)
	case HAS:
		r, err := cache.Has(cmd.Key)
		if err == nil && r {
			cmd.Resp = yes
		} else {
			cmd.Resp = no
		}
	case PING:
		cmd.Resp = pong
	case STATUS:
		cmd.Resp = ok
		log.Info("status: %s", cmd.Resp)
	case SHUTDOWN:
		log.Info("shutdown command received...")
		cmd.Resp = fail
	default:
		msg := fmt.Sprintf("unknown command id: %d", cmd.Op)
		log.Warn(msg)
		err = errors.New(msg)
		cmd.Resp = fail
	}

	return err
}

// a string representation of the command buffer
func (cmd *Command) String() string {
	return fmt.Sprintf("Id:%s,Op:%d,Key:%s,Value:%s,Resp:%s", cmd.Id, cmd.Op, cmd.Key, cmd.Value, cmd.Resp)
}

// a public helper method to create a full comman structure
func CreateCommand(id IdType, op CommandOp, key, value []byte) *Command {
	cmd := Command{Id: id, Op: op, Key: key, Value: value}

	return &cmd
}
