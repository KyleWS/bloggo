package models

import "gopkg.in/mgo.v2/bson"

// Simple example data structure.
// Screw interfaces i'll figure them out later.
type GenericData struct {
	// Add as many fields as u want
	// Default values, remember that
	ID bson.ObjectId `json:"id,omitempty" bson:"_id"`
}
