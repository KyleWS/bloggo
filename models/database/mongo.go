package database

import (
	"fmt"

	"github.com/KyleWS/bloggo/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
		panic("NIL pointer passed for session. Please make sure the ")
	}

	return &Mongo{
		Session:    sess,
		Dbname:     dbName,
		Colname:    collectionName,
		Collection: sess.DB(dbName).C(collectionName),
	}
}

// Insert receiver on mongo database to allow insertion.
func (mongo *Mongo) Insert(data *models.GenericData) (*models.GenericData, error) {
	data.ID = bson.NewObjectId()
	if err := mongo.Collection.Insert(data); err != nil {
		return nil, fmt.Errorf("error inserting struct of type generic: %v\n%v", data, err)
	}
	return data, nil
}

// GetByID rreturns object with given ID
func (mongo *Mongo) GetByID(id bson.ObjectId) (*models.GenericData, error) {
	result := &models.GenericData{}
	if err := mongo.Collection.Find(bson.M{"_id": id}).One(&result); err != nil {
		return nil, fmt.Errorf("error in gettign record in GetByID %v", err)
	}
	return result, nil
}

// See models/genericModel.go for how updates are represented in a struct.

// Returns updated record if it worked
func (mongo *Mongo) Update(id bson.ObjectId, updates *models.Updates) (*models.GenericData, error) {
	recordToUpdate, err := mongo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("error getting record to be updated: %v", err)
	}
	change := mgo.Change{
		Update:    bson.M{"$set": recordToUpdate},
		ReturnNew: true,
	}
	result := &models.GenericData{}
	if _, err := mongo.Collection.FindId(id).Apply(change, result); err != nil {
		return nil, fmt.Errorf("error updating record: %v", err)
	}
	return result, nil
}

// Delete record
func (mongo *Mongo) Delete(id bson.ObjectId) error {
	if err := mongo.Collection.RemoveId(id); err != nil {
		return fmt.Errorf("error deleting record: %v", err)
	}
	return nil
}

// GetAll returns all the data in the database
func (mongo *Mongo) GetAll() ([]*models.GenericData, error) {
	data := []*models.GenericData{}
	if err := mongo.Collection.Find(nil).Limit(10).All(&data); err != nil {
		return nil, fmt.Errorf("error getting all data: %v", err)
	}
	return data, nil
}
