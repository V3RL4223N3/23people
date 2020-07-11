package models

import (
	"fmt"
	"github.com/V3RL4223N3/23people/app/models/mongodb"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Person struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	NationalId string        `json:"nationalId" bson:"national_id"`
	Name       string        `json:"name" bson:"name"`
	LastName   string        `json:"lastName" bson:"last_name"`
	Age        int           `json:"age" bson:"age"`
	PictureUrl string        `json:"pictureUrl" bson:"picture_url"`
	CreatedAt  time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at" bson:"updated_at"`
}

func newPersonCollection() *mongodb.Collection {
	return mongodb.NewCollectionSession("persons")
}

// AddPerson insert a new Person into database and returns
// last inserted person on success.
func AddPerson(m Person) (person Person, err error) {
	c := newPersonCollection()
	defer c.Close()
	m.ID = bson.NewObjectId()
	m.CreatedAt = time.Now()
	checkPerson, err := GetPerson(m.NationalId)

	if err != nil {
		if checkPerson.NationalId == "" {
			return m, c.Session.Insert(m)
		}

	}

	return m, err
}

// UpdatePerson update a Person into database and returns
// last nil on success.
func (m Person) UpdatePerson(id string) error {
	c := newPersonCollection()
	defer c.Close()

	var person Person
	fmt.Println(m.NationalId)

	err := c.Session.Find(bson.M{"national_id": id}).One(&person)

	err = c.Session.Update(bson.M{
		"national_id": m.NationalId,
	}, bson.M{
		"$set": bson.M{
			"national_id": m.NationalId, "name": m.Name, "last_name": m.LastName, "age": m.Age, "picture_url": m.PictureUrl, "updatedAt": time.Now()},
	})
	return err
}

// DeletePerson Delete Person from database and returns
// last nil on success.
func (m Person) DeletePerson() error {
	c := newPersonCollection()
	defer c.Close()

	err := c.Session.Remove(bson.M{"_id": m.ID})
	return err
}

// GetPersons Get all Person from database and returns
// list of Person on success
func GetPersons() ([]Person, error) {
	var (
		persons []Person
		err     error
	)

	c := newPersonCollection()
	defer c.Close()

	err = c.Session.Find(nil).Sort("-createdAt").All(&persons)
	return persons, err
}

// GetPerson Get a Person from database and returns
// a Person on success
func GetPerson(id string) (Person, error) {
	var (
		person Person
		err    error
	)

	c := newPersonCollection()
	defer c.Close()
	err = c.Session.Find(bson.M{"national_id": id}).One(&person)
	return person, err
}
