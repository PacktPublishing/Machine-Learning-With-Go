package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func main() {

	// Open an embedded.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("embedded.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a "bucket" in the boltdb file for our data.
	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	// Put the map keys and values into the BoltDB file.
	if err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		err := b.Put([]byte("mykey"), []byte("myvalue"))
		return err
	}); err != nil {
		log.Fatal(err)
	}

	// Output the keys and values in the embedded
	// BoltDB file to standard out.
	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key: %s, value: %s\n", k, v)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}
