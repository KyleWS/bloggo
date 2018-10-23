package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Location struct {
	// Add as many fields as u want
	// Default values, remember that
	ID   bson.ObjectId `json:"id,omitempty" bson:"_id"`
	Name string        `json:"name" bson:"name"`
	Loc  string        `json:"loc" bson:"loc"`
	Time time.Time     `json:"time" bson:"time"`
}

type LocationPost struct {
	Name string `json:"name" bson:"name"`
	Loc  string `json:"loc" bson:"loc"`
}

func (lp *LocationPost) Received() *Location {
	loc := &Location{}
	loc.ID = bson.NewObjectId()
	loc.Time = time.Now()
	loc.Name = lp.Name
	loc.Loc = lp.Loc
	return loc
}

func NewLocationPost(name string, location string) *LocationPost {
	return &LocationPost{
		Name: name,
		Loc:  location,
	}
}
