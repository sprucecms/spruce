package sprucelib

import (
	"encoding/json"

	"github.com/boltdb/bolt"
	"golang.org/x/crypto/bcrypt"
)

type BoltDataStore struct {
	DB *bolt.DB
}

func NewBoltDataStore() (db BoltDataStore, err error) {
	// TODO Make the database path configurable.
	bdb, err := bolt.Open("spruce.db", 0600, nil)
	if err != nil {
		return
	}
	ds := BoltDataStore{DB: bdb}
	ds.CreateBuckets()
	return ds, err
}

func (ds BoltDataStore) Close() {
	ds.DB.Close()
}

func (ds BoltDataStore) CreateBuckets() {
	err := ds.DB.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("resources"))
		tx.CreateBucketIfNotExists([]byte("templates"))
		tx.CreateBucketIfNotExists([]byte("user-tokens"))
		ub, _ := tx.CreateBucketIfNotExists([]byte("users"))
		json, _ := json.Marshal(User{
			ID:           1,
			Username:     "admin",
			PasswordHash: "$2a$06$wORnT37BtlThVBTKa/0/rui38N7XA4NkjUGDjCVIIcQMo3FQ95uuS",
		})
		return ub.Put([]byte("admin"), json)
	})
	if err != nil {
		panic(err)
	}
}

func (ds BoltDataStore) GetUserByUsernameAndPassword(username string, password string) (user User, err error) {
	err = ds.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("users"))
		val := bucket.Get([]byte(username))
		if val == nil {
			return ErrUnknownUser
		}
		err = json.Unmarshal(val, &user)
		if err != nil {
			return err
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		user = User{}
	}

	return
}

func (db BoltDataStore) NewTokenForUser(user User) (token string, err error) {
	return "NOTAREALTOKEN", nil
}

func (db BoltDataStore) GetUserByToken(token string) (user User, err error) {
	return
}
