package main

import (
	"log"

	c "github.com/MRHGST/cloverdb/v2"
	badgerstore "github.com/MRHGST/cloverdb/v2/store/badger"
	"github.com/dgraph-io/badger/v4"
)

func main() {
	store, err := badgerstore.OpenWithOptions(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		log.Fatal(err)
	}

	db, err := c.OpenWithStore(store)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
