//
// SpotCacheDb - utilities to manipulate the cache database; dump, restore, size, etc
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-04-16 07:58:14
//

package main

import (
	"fmt"
	"os"

	"github.com/darrylwest/spot-cache/spotcache"
	"github.com/syndtr/goleveldb/leveldb"
)

func parseArgs() *spotcache.Config {
	cfg := spotcache.NewDefaultConfig()

	// TODO use flags to define backup, restore, copy, size

	fmt.Println(cfg)

	return cfg
}

func main() {
	parseArgs()

	db, err := leveldb.OpenFile(os.Getenv("HOME")+"/.spotcache/cachedb", nil)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()

		fmt.Printf("%s : %s\n", key, value)
	}

	iter.Release()
	err = iter.Error()
	if err != nil {
		panic(err)
	}
}
