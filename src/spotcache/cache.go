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

// TTLSeconds compatible with time.Unix() in seconds
type TTLSeconds int64

// Cache the cache object
type Cache struct {
	path string
}

// NewCache create a new cache object
func NewCache(cfg *Config) *Cache {
	cache := Cache{}

	cache.path = cfg.Dbpath

	return &cache
}

// CreateOptions create the cache options
func (c Cache) CreateOptions() *opt.Options {
	opts := opt.Options{}

	return &opts
}

// Open open the cache db
func (c Cache) Open() error {
	var err error
	db, err = leveldb.OpenFile(c.path, c.CreateOptions())

	if err != nil {
		log.Error(fmt.Sprintf("error opening database at path %s, %v", c.path, err))
	}

	return err
}

// Close close the cache db
func (c Cache) Close() {
	if db != nil {
		log.Info("closing cache database...")
		db.Close()
	}
}

// Put define the methods get, put, delete, has, ttl, etc...
func (c *Cache) Put(key, value []byte, ttl TTLSeconds) error {
	return db.Put(key, value, nil)
}

// Get return the data from key
func (c *Cache) Get(key []byte) ([]byte, error) {
	return db.Get(key, nil)
}

// Has return true if cache has the key
func (c *Cache) Has(key []byte) (bool, error) {
	return db.Has(key, nil)
}

// Delete delete the data based on key
func (c *Cache) Delete(key []byte) error {
	return db.Delete(key, nil)
}

// TTL return the time to live
func (c *Cache) TTL(key []byte) TTLSeconds {
	return 0
}

// Keys return all keys in the database
func (c *Cache) Keys() ([]string, error) {
	keys := []string{}

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		keys = append(keys, string(iter.Key()))
	}
	iter.Release()

	return keys, iter.Error()
}
