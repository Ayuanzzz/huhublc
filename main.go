package main

import (
	"log"
	"fmt"
	"github.com/boltdb/bolt"
)

func main() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//update
	db.Update(func(tx *bolt.Tx) error {
		//创建bucket
		b, err := tx.CreateBucket([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		//写入数据
		if b != nil{
			err := b.Put([]byte("1"),[]byte("11"))
			if err != nil{
				return err
			}
		}
		return nil
	})

	//read
	db.View(func(tx *bolt.Tx) error{
		//获取bucket
		b:=tx.Bucket([]byte("MyBucket"))
		if b != nil{
			value := b.Get([]byte("1"))
			fmt.Printf("value:%s\n", string(value))
		}
		return nil
	})
}