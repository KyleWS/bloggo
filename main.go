package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/KyleWS/bloggo/handlers"
	"github.com/KyleWS/bloggo/models/database"

	mgo "gopkg.in/mgo.v2"
)

const defaultAddr = ":443"
const defaultMongo = "localhost:27017"
const defaultMongoDBName = "DefaultDB"
const defaultMongoColName = "DefaultCol"

// See context.go for basic context management
// See cors.go for basic middleware management

func main() {
	addr := os.Getenv("SERVER_ADDRESS")
	if len(addr) == 0 {
		addr = defaultAddr
	}

	tlsKeypath := os.Getenv("TLSKEY")
	tlsCertPath := os.Getenv("TLSCERT")

	mongoAddr := os.Getenv("DATABASE_ADDRESS")
	//default to "localhost"
	if len(mongoAddr) == 0 {
		mongoAddr = defaultMongo
	}
	// Dial mongo database
	mongoSess, err := mgo.Dial(mongoAddr)
	if err != nil {
		log.Fatalf("error dialing mongo database. check that the provided URL is reachable from this program: %v", err)
	}
	mongoDatabase := database.NewMongoStore(mongoSess, defaultMongoDBName, defaultMongoColName)
	handlerContext := handlers.NewHandlerContext(*mongoDatabase)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlerContext.DefaultHandler)

	// Serving static files can be done with this format.
	//dir := http.Dir("/images") // This file must be created
	//fs := http.FileServer(dir)
	//mux.Handle("/static/", http.StripPrefix("/static/", fs)) // use the 'static' path in the URL followed by the resource name

	corsHandler := handlers.NewCORS(mux)
	fmt.Printf("server is listening at http://localhost[%s]...\n", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeypath, corsHandler))
}
