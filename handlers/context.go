package handlers

import "github.com/KyleWS/bloggo/models/database"

// Default Context and CORS structs/handlers to facilitate
// bare requirements for request handling.
type Context struct {
	db database.Mongo
}

func NewHandlerContext(mongo database.Mongo) *Context {
	return &Context{
		db: mongo,
	}
}
