package store

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/jaakidup/reactor-core/model"
)

var updateUserChan chan model.User
var db = openDB()
var dbBucket = dbBucketName()

func dbBucketName() []byte {
	return []byte("Users")
}

func openDB() *bolt.DB {
	db, err := bolt.Open("./db/users.db", 0600, nil)
	if err != nil {
		log.Fatalln("Couldn't open Bolt")
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(dbBucket)
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Couldn't set up the Buckets")
	}

	return db
}

func init() {
	fmt.Println("UserStore init")

	updateUserChan = make(chan model.User)

	go func() {
		for {
			user := <-updateUserChan
			marshalledUser, err := json.Marshal(user)
			if err != nil {
				log.Println("Failed Marshalling user for saving in DB")
			}

			err = db.Update(func(tx *bolt.Tx) error {
				b := tx.Bucket(dbBucket)
				err := b.Put([]byte(user.ID), marshalledUser)
				return err
			})
			if err != nil {
				log.Println("Failed saving user with ID: ", user.ID)
				log.Println("Handle this with notification to admin")
			}

		}
	}()

	fmt.Println("User Write Instance started")
}

// MakeUserStore ...
func MakeUserStore() *UserStore {
	return &UserStore{}
}

// UserStore ...
type UserStore struct {
}

// Update ...
func (UserStore) Update(user model.User) (string, error) {
	log.Println("Storing User", user)

	if user.ID == "" {
		log.Println("User doesn't have an ID, so let's create one.")
		user.ID = GenerateUUID()
		log.Println("Generating new UUID")
	}

	// updateUserChan takes a user in a separate goroutine
	// as it only allows a single write instance
	updateUserChan <- user

	return string(user.ID), nil
}

// Get ...
func (UserStore) Get(id string) (model.User, error) {
	user := model.User{}

	err := db.View(func(tx *bolt.Tx) error {
		userdata := tx.Bucket(dbBucket).Get([]byte(id))
		json.Unmarshal(userdata, &user)
		return nil
	})
	if err != nil {
		return user, err
	}

	return user, nil
}

// GetAll ...
func (UserStore) GetAll() ([]model.User, error) {
	users := []model.User{}

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(dbBucket)
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
