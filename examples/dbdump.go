package main

import (
	"fmt"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
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
