package main

import (
	"fmt"

	"bytes"

	"encoding/gob"

	"github.com/boltdb/bolt"
)

type A struct {
	Name string
	Age  int
}

func structToBuf(data A) []byte {
	var buf = &bytes.Buffer{}
	fmt.Println(data.Name)
	c1 := gob.NewEncoder(buf)
	c1.Encode(data)

	fmt.Println(buf.Bytes())
	return buf.Bytes()
}

func bufToStruct(data []byte) interface{} {
	fmt.Println(data)
	var buf = bytes.NewBuffer(data)
	var std A
	c1 := gob.NewDecoder(buf)

	c1.Decode(&std)
	fmt.Println("name is: " + std.Name)
	return std
}

func main() {
	var s1 = A{Name: "jiang", Age: 10}

	b1 := structToBuf(s1)

	db, _ := bolt.Open("jiang.db", 0600, nil)

	err := db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte("MyFriendsBucket")); err != nil {

			return err
		}

		b := tx.Bucket([]byte("MyFriendsBucket"))
		err := b.Put([]byte("one"), b1)
		return err
	})
	if err != nil {
		return
	}
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyFriendsBucket"))
		v := b.Get([]byte("one"))
		d1 := bufToStruct(v)
		switch d2 := d1.(type) {
		case A:
			fmt.Println("this name " + d2.Name)
		}
		fmt.Printf("The answer is: %s\n", v)
		return nil
	})

}
