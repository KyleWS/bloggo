package handlers

import (
	"github.com/KyleWS/chikkin-server/models"
)

// Default Context and CORS structs/handlers to facilitate
// bare requirements for request handling.
type Context struct {
	db models.Mongo
}

func NewHandlerContext(mongo models.Mongo) *Context {
	return &Context{
		db: mongo,
	}
}
