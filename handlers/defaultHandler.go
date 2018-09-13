package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/KyleWS/bloggo/models"
)

// Basic handler. Can be assigned to any path for testing
// GET: returns "Hello"
// POST: can be used to test database operations
func (ctx *Context) DefaultHandler(w http.ResponseWriter, r *http.Request) {
	all := r.URL.Query()["all"]
	id := r.URL.Query()["id"]
	switch r.Method {
	case "GET":
		if len(id) > 0 {
			if !bson.IsObjectIdHex(id[0]) {
				http.Error(w, fmt.Sprintf("error ID must be valid: %d", id[0]), http.StatusBadRequest)
				return
			}
			record, err := ctx.db.GetByID(bson.ObjectIdHex(id[0]))
			if err != nil {
				http.Error(w, fmt.Sprintf("error getting record by id: %v", err), http.StatusInternalServerError)
				return
			}
			w.Header().Add("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(record); err != nil {
				http.Error(w, fmt.Sprintf("error encoding response value to JSON: %v", err), http.StatusInternalServerError)
				return
			}
		}
		if all != nil {
			allData, err := ctx.db.GetAll()
			if err != nil {
				http.Error(w, fmt.Sprintf("error getting all records: %v", err), http.StatusInternalServerError)
				return
			}
			w.Header().Add("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(allData); err != nil {
				http.Error(w, fmt.Sprintf("error encoding response value to JSON: %v", err), http.StatusInternalServerError)
				return
			}
		} else {
			w.Write([]byte("Hello"))
		}
	case "POST":
		data := &models.GenericData{}
		if err := json.NewDecoder(r.Body).Decode(data); err != nil {
			http.Error(w, fmt.Sprintf("error decoding JSON: %v", err), http.StatusBadRequest)
			return
		}

		dataReturn, err := ctx.db.Insert(data)
		if err != nil {
			http.Error(w, fmt.Sprintf("error inserting task: %v", err), http.StatusInternalServerError)
			return
		}
		// set header to json so that it formats properly for the user who wants it
		w.Header().Add("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(dataReturn); err != nil {
			http.Error(w, fmt.Sprintf("error encoding response value to JSON: %v", err), http.StatusInternalServerError)
			return
		}
	case "DELETE":
		if len(id) > 0 {
			if !bson.IsObjectIdHex(id[0]) {
				http.Error(w, fmt.Sprintf("error ID must be valid: %d", id[0]), http.StatusBadRequest)
				return
			}
			if err := ctx.db.Delete(bson.ObjectIdHex(id[0])); err != nil {
				http.Error(w, fmt.Sprintf("error deleting record: %v", err), http.StatusInternalServerError)
				return
			}
		}

	case "PATCH":
		// need to wire this up but it seems sort of pointles atm
		// look in auth.go for a better idea
		return
	default:
		http.Error(w, "method must be GET or POST", http.StatusMethodNotAllowed)
		return
	}
}
