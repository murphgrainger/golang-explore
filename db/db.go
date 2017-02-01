package db

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
  ID    string `json:"id" bson:"_id,omitempty"`
  Firstname string `json:"firstname,omitempty"`
  Lastname string `json:"lastname,omitempty"`
  Address string `json:"address,omitempty"`
}

var db *mgo.Database

func init() {
	session, err := mgo.Dial("localhost/api_db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db = session.DB("api_db")
}

func collection() *mgo.Collection {
	return db.C("people")
}

// GetAll returns all items from the database.
func GetAll() ([]Person, error) {
	res := []Person{}

	if err := collection().Find(nil).All(&res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetOne returns a single item from the database.
func GetOne(id string) (*Person, error) {
	res := Person{}

	if err := collection().Find(bson.M{"_id": id}).One(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Save inserts a person to the database.
func Save(person Person) error {
	return collection().Insert(person)
}

// Remove deletes an person from the database
func Remove(id string) error {
	return collection().Remove(bson.M{"_id": id})
}
