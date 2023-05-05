package blockCore

import (
	"fmt"
	"github.com/boltdb/bolt"
)

func testBolt() {

	db, err := bolt.Open("test.db", 0600, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("b1"))
		if err != nil {
			panic(err)
		}
		bucket.Put([]byte("username"), []byte("小明"))
		return nil
	})

	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("b1"))

		if bucket != nil {
			v := bucket.Get([]byte("username"))
			fmt.Println(string(v))
		}
		return nil
	})
}
