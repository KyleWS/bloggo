package models

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

// Mongo represents the connection to a mongodb
// server. The Collection attribute is a shorthand
// useful for calling operations on the given db/colname
type Mongo struct {
	Session    *mgo.Session
	Dbname     string
	Colname    string
	Collection *mgo.Collection
}

// Creates a new mongo database struct, to be used when
// wanting to store information persistently across application
// runtime.
func NewMongoStore(sess *mgo.Session, dbName string, collectionName string) *Mongo {
	if sess == nil {
		panic("NIL pointer passed for session. Please make sure the server is up")
	}

	return &Mongo{
		Session:    sess,
		Dbname:     dbName,
		Colname:    collectionName,
		Collection: sess.DB(dbName).C(collectionName),
	}
}

// Insert location string for given person's name
func (mongo *Mongo) InsertLocation(locationPost *LocationPost) (*Location, error) {
	location := locationPost.Received()
	if err := mongo.Collection.Insert(location); err != nil {
		return nil, fmt.Errorf("error inserting struct of type generic: %v\n%v", location, err)
	}
	return location, nil
}

// Returns the most recent location chikbits
func (mongo *Mongo) GetRecent() ([]*Location, error) {
	locations := []*Location{}
	// Remember, .All(&) takes an address. Must be address here  V V V
	// Even though it wont show a red error.
	if err := mongo.Collection.Find(nil).Sort("-time").Limit(20).All(&locations); err != nil {
		return nil, fmt.Errorf("error getting all data: %v", err)
	}
	return locations, nil
}
