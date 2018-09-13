package models

import (
	"gopkg.in/mgo.v2/bson"
)

// Simple example data structure.
// Screw interfaces i'll figure them out later.
type GenericData struct {
	// Add as many fields as u want
	// Default values, remember that
	ID bson.ObjectId `json:"id,omitempty" bson:"_id"`
}

type Updates struct {
	ID bson.ObjectId
}

func (data *GenericData) ApplyUpdates(updates *Updates) {
	// do this for every field in updates that you want to update
	data.ID = updates.ID
}
