//
// Cache - an interface definition and thin wrapper around boltdb
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-03-18 13:52:49
//

package spotcache

import (
	"fmt"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

// TTLSeconds compatible with time.Unix() in seconds
type TTLSeconds int64

// Cache the cache object
type Cache struct {
	path   string
	bucket []byte
}

// NewCache create a new cache object
func NewCache(cfg *Config) *Cache {
	cache := Cache{}

	cache.path = cfg.Dbpath
	cache.bucket = []byte("cache")

	return &cache
}

// CreateOptions create the cache options
func (c Cache) CreateOptions() *bolt.Options {
	opts := bolt.Options{}

	return &opts
}

// Open open the cache db
func (c Cache) Open() error {
	log.Info("opening database at path %s", c.path)

	var err error
	db, err = bolt.Open(c.path, 0600, nil)

	if err != nil {
		log.Error(fmt.Sprintf("error opening database at path %s, %v", c.path, err))
		return err
	}

	// open the cache bucket...
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(c.bucket)
		if err != nil {
			log.Error(fmt.Sprintf("error creating bucket: %s %v", c.bucket, err))
			return err
		}

		log.Info("bucket: %s created...", c.bucket)
		return nil
	})

	return err
}

// Close close the cache db
func (c Cache) Close() {
	if db != nil {
		log.Info("closing cache database...")
		db.Close()
		db = nil
	}
}

// Put define the methods get, put, delete, has, ttl, etc...
func (c Cache) Put(key, value []byte, ttl TTLSeconds) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(c.bucket)
		err := b.Put(key, value)
		return err
	})
}

// Get return the data from key; return nil,nil on not found
func (c Cache) Get(key []byte) ([]byte, error) {
	var value []byte

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(c.bucket)
		value = b.Get(key)

		return nil
	})

	log.Debug("Get %s %v %v", key, value, err)

	return value, err
}

// Has return true if cache has the key
func (c Cache) Has(key []byte) (bool, error) {
	var has bool

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(c.bucket)
		has = b.Get(key) != nil
		return nil
	})

	log.Debug("has %s %v", key, has)

	return has, err
}

// Delete delete the data based on key
func (c Cache) Delete(key []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(c.bucket)
		return b.Delete(key)
	})
}

// TTL return the time to live
func (c *Cache) TTL(key []byte) TTLSeconds {
	return 0
}

// Keys return all keys in the database
func (c Cache) Keys() ([]string, error) {
	keys := []string{}

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(c.bucket)
		b.ForEach(func(k, v []byte) error {
			keys = append(keys, string(k))
			return nil
		})

		return nil
	})

	return keys, err
}
