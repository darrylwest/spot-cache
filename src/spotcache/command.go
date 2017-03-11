/**
 * command - parse and response to commands
 *
 * @author darryl.west <darryl.west@raincitysoftware.com>
 * @created 2017-03-11 13:56:46
 */

package spotcache

import (
    "github.com/syndtr/goleveldb/leveldb"
    "github.com/syndtr/goleveldb/leveldb/opt"
)

type Command struct {
    reqid string
    cmd string
    params interface{}
}

var db *leveldb.DB

func opendb(cfg *Config) {
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

func closedb() {
    if db != nil {
        log.Info("closing the database")
        db.Close()
    }
}

func ParseCommand(buf []byte) (*Command, error) {
    return nil, nil
}

func (command *Command) exec() {
}

