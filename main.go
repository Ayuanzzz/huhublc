package main

import (
	"fmt"
	"huhuBLC/BLC"

	"github.com/boltdb/bolt"
)

func main() {
	bc := BLC.CreateBlockChainWithGenesisBlock()

	// bc.AddBlock([]byte("alice send 100 eth to bob"))
	// bc.AddBlock([]byte("bob send 100 eth to alice"))

	bc.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocks"))
		if b != nil {
			hash := b.Get([]byte("1"))
			fmt.Printf("hash: %v\n", hash)
		}
		return nil
	})
}
