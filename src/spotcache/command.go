/**
 * command - parse and response to commands
 *
 * @author darryl.west <darryl.west@raincitysoftware.com>
 * @created 2017-03-11 13:56:46
 */

package spotcache

import (
    "fmt"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

const (
    ok string = "ok"
)

var (
    RESP_OK []byte = []byte(ok)
)

type Command struct {
	id     []byte
	op     []byte
    key    []byte
	value  []byte
    resp   []byte
}

var db *leveldb.DB

func Opendb(cfg *Config) {
	log.Info("open database in %s", cfg.dbpath)

	var err error

	db, err = leveldb.OpenFile(cfg.dbpath, CreateOptions(cfg))

	if err != nil {
		panic(err)
	}
}

func CreateOptions(dfg *Config) *opt.Options {
	opts := opt.Options{}

	return &opts
}

func Closedb() {
	if db != nil {
		log.Info("closing the database")
		db.Close()
	}
}

func ParseCommand(buf []byte) (*Command, error) {
	return nil, nil
}

func (cmd *Command) Exec() error {
    err := db.Put(cmd.key, cmd.value, nil)
    cmd.resp = RESP_OK
    return err
}

func (cmd *Command) ToString() string {
    return fmt.Sprintf("id:%s,op:%s,key:%s,value:%s,resp:%s", cmd.id, cmd.op, cmd.key, cmd.value, cmd.resp)
}

func (cmd *Command) GetResp() []byte {
    return cmd.resp
}

func CreateCommand(id, op, key, value []byte) *Command {
    cmd := Command{ id:id, op:op, key:key, value:value }

    return &cmd
}
