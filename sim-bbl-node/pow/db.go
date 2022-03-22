package pow

import (
	"encoding/json"
	"fmt"
	"log"
	"path"
	"runtime"
	"time"

	"github.com/boltdb/bolt"
)

var db *bolt.DB
var open bool

func OpenDB() error {
	var err error
	_, filename, _, _ := runtime.Caller(0) // get full path of this file
	dbfile := path.Join(path.Dir(filename), "data.db")
	config := &bolt.Options{Timeout: 1 * time.Second}
	db, err = bolt.Open(dbfile, 0600, config)
	if err != nil {
		log.Fatal(err)
	}
	open = true
	return nil
}

func CloseDB() {
	open = false
	db.Close()
}

func (bh *BlockHash) encode() ([]byte, error) {
	enc, err := json.Marshal(bh)
	if err != nil {
		return nil, err
	}
	return enc, nil
}

func (bh *BlockHash) SaveBlockHash() error {
	if !open {
		return fmt.Errorf("db must be opened before saving!")
	}
	err := db.Update(func(tx *bolt.Tx) error {
		block, err := tx.CreateBucketIfNotExists([]byte("blockhash"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		enc, err := bh.encode()
		if err != nil {
			return fmt.Errorf("could not encode Block %s: %s", bh.Hash, err)
		}
		err = block.Put([]byte(bh.Hash), enc)
		return err
	})
	return err
}

func GetBlockHash(hash string) (map[string]interface{}, error) {
	if !open {
		return nil, fmt.Errorf("db must be opened before saving!")
	}
	var jsonMap map[string]interface{}
	err := db.View(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket([]byte("blockhash"))
		k := []byte(hash)
		jsonMap, err = decode(b.Get(k))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Could not get Block Hash%s ", hash)
		return nil, err
	}
	return jsonMap, nil
}

func (b *Block) SaveBlock() error {
	if !open {
		return fmt.Errorf("db must be opened before saving!")
	}
	err := db.Update(func(tx *bolt.Tx) error {
		block, err := tx.CreateBucketIfNotExists([]byte("block"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		enc, err := b.encode()
		if err != nil {
			return fmt.Errorf("could not encode Block %s: %s", b.Height, err)
		}
		err = block.Put(i32tob(b.Height), enc)
		return err
	})
	return err
}

func (p *Block) encode() ([]byte, error) {
	enc, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return enc, nil
}

func i32tob(val uint32) []byte {
	r := make([]byte, 4)
	for i := uint32(0); i < 4; i++ {
		r[i] = byte((val >> (8 * i)) & 0xff)
	}
	return r
}

func btoi32(val []byte) uint32 {
	r := uint32(0)
	for i := uint32(0); i < 4; i++ {
		r |= uint32(val[i]) << (8 * i)
	}
	return r
}

func decode(data []byte) (map[string]interface{}, error) {

	var jsonMap map[string]interface{}
	err := json.Unmarshal(data, &jsonMap)
	if err != nil {
		return nil, err
	}
	return jsonMap, nil
}

func GetBlock(hight string) (map[string]interface{}, error) {
	if !open {
		return nil, fmt.Errorf("db must be opened before saving!")
	}
	var jsonMap map[string]interface{}
	err := db.View(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket([]byte("block"))
		k := []byte(hight)
		fmt.Printf("---------------test------- ------- %s", b.Get(k))
		jsonMap, err = decode(b.Get(k))
		fmt.Printf("------- =old Hash ------- %s", jsonMap["Hash"])
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return jsonMap, nil
}

func GetNewestBlock(bucket string) {
	db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte(bucket)).Cursor()
		//k, v := c.Next()
		//fmt.Printf("key=%s, value=%s\n", k, v)

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}
		return nil
	})
}

/**
func List(bucket string) {
	db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte(bucket)).Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}
		return nil
	})
}

func ListPrefix(bucket, prefix string) {
	db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte(bucket)).Cursor()
		p := []byte(prefix)
		for k, v := c.Seek(p); bytes.HasPrefix(k, p); k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}
		return nil
	})
}

func ListRange(bucket, start, stop string) {
	db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte(bucket)).Cursor()
		min := []byte(start)
		max := []byte(stop)
		for k, v := c.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = c.Next() {
			fmt.Printf("%s: %s\n", k, v)
		}
		return nil
	})
}
**/
