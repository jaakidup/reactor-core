package store

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/jaakidup/project/model"
)

func init() {
	fmt.Println("UserStore init")
}

// MakeUserStore ...
func MakeUserStore() *UserStore {

	db, err := bolt.Open("./db/users.db", 0600, nil)
	if err != nil {
		log.Fatalln("Couldn't open Bolt")
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Users"))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		// _, err = root.CreateBucketIfNotExists([]byte("WEIGHT"))
		// if err != nil {
		// return fmt.Errorf("could not create weight bucket: %v", err)
		// }
		// _, err = root.CreateBucketIfNotExists([]byte("ENTRIES"))
		// if err != nil {
		// return fmt.Errorf("could not create days bucket: %v", err)
		// }
		return nil
	})
	if err != nil {
		fmt.Println("Couldn't set up the Buckets")
	}

	return &UserStore{db: db}
}

// MakeUUID ...
func MakeUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	fmt.Println(uuid)
	return uuid
}

// UserStore ...
type UserStore struct {
	db *bolt.DB
}

// Save ...
func (us UserStore) Save(user model.User) (string, error) {
	log.Println("Storing User", user)

	user.ID = MakeUUID()
	marshalledUser, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	err = us.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Users"))
		err := b.Put([]byte(user.ID), marshalledUser)
		return err
	})
	if err != nil {
		return "nil", err
	}

	return string(user.ID), nil
}

// Get ...
func (us UserStore) Get(id string) (model.User, error) {
	user := model.User{}

	err := us.db.View(func(tx *bolt.Tx) error {
		userdata := tx.Bucket([]byte("Users")).Get([]byte(id))
		json.Unmarshal(userdata, &user)
		return nil
	})
	if err != nil {
		return user, err
	}

	return user, nil
}

// GetAll ...
func (us UserStore) GetAll() ([]model.User, error) {
	users := []model.User{}

	err := us.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Users"))
		b.ForEach(func(k, v []byte) error {
			// fmt.Println(string(k), string(v))
			user := model.User{}
			json.Unmarshal(v, &user)
			users = append(users, user)
			return nil
		})
		return nil
	})
	if err != nil {
		fmt.Println("Couldn't fetch Users")
	}

	return users, nil
}
