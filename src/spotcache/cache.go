//
// Cache - an interface definition and thin wrapper around leveldb
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-03-18 13:52:49
//

package spotcache

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

var db *leveldb.DB

type TTL uint64

type Cache struct {
	path string
}

func NewCache(cfg *Config) *Cache {
	cache := Cache{}

	cache.path = cfg.Dbpath

	return &cache
}

func (c Cache) CreateOptions() *opt.Options {
	opts := opt.Options{}

	return &opts
}

func (c Cache) Open() error {
	var err error
	db, err = leveldb.OpenFile(c.path, c.CreateOptions())

	if err != nil {
		log.Error(fmt.Sprintf("error opening database at path %s, %v", c.path, err))
	}

	return err
}

func (c Cache) Close() {
	if db != nil {
		log.Info("closing cache database...")
		db.Close()
	}
}

// define the methods get, put, delete, has, ttl, etc...
func (c *Cache) Put(key, value []byte, ttl TTL) error {
	return db.Put(key, value, nil)
}

func (c *Cache) Get(key []byte) ([]byte, error) {
	return db.Get(key, nil)
}

func (c *Cache) Has(key []byte) (bool, error) {
	return db.Has(key, nil)
}

func (c *Cache) Ttl(key []byte) TTL {
	return 0
}

// should have a way of dumping the entire database or at least generating stats
