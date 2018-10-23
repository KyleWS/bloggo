package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/KyleWS/chikkin-server/handlers"
	"github.com/KyleWS/chikkin-server/models"
	mgo "gopkg.in/mgo.v2"
)

/*
DEBUG INSTRUCTIONS:
To debug in vscode you need to do the following things
1. Set default addr to 2555 (not 443 because it is running locally)
2. Comment out the TLS variables
3. Hardcode TLS path or run program with ListenAndServe instead of ListenAndServeTLS
*/

const defaultAddr = ":443"
const defaultMongo = "localhost:27017"
const defaultMongoDBName = "chikkin-db"
const defaultMongoColName = "chikkin-locations-collection-v1"

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
	mongoDatabase := models.NewMongoStore(mongoSess, defaultMongoDBName, defaultMongoColName)
	handlerContext := handlers.NewHandlerContext(*mongoDatabase)

	mux := http.NewServeMux()
	mux.HandleFunc("/location", handlerContext.LocationRequestHandler)

	// Serving static files can be done with this format.
	//dir := http.Dir("/images") // This file must be created
	//fs := http.FileServer(dir)
	//mux.Handle("/static/", http.StripPrefix("/static/", fs)) // use the 'static' path in the URL followed by the resource name

	corsHandler := handlers.NewCORS(mux)
	fmt.Printf("server is listening at http://localhost[%s]...\n", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeypath, corsHandler))
}
