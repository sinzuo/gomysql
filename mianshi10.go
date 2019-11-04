package main

import (
	"fmt"

	"github.com/boltdb/bolt"
)

func main() {
	db, _ := bolt.Open("jiang.db", 0600, nil)

	err := db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte("MyFriendsBucket")); err != nil {

			return err
		}

		b := tx.Bucket([]byte("MyFriendsBucket"))
		err := b.Put([]byte("one"), []byte("zhangsan"))
		return err
	})
	if err != nil {
		return
	}
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyFriendsBucket"))
		v := b.Get([]byte("one"))
		fmt.Printf("The answer is: %s\n", v)
		return nil
	})

}
