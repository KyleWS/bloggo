package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KyleWS/bloggo/models"
)

// Basic handler. Can be assigned to any path for testing
// GET: returns "Hello"
// POST: can be used to test database operations
func (ctx *Context) DefaultHandler(w http.ResponseWriter, r *http.Request) {
	op := r.URL.Query()["all"]
	switch r.Method {
	case "GET":
		if len(op) > 0 {
			allData, err := ctx.db.GetAll()
			if err != nil {
				http.Error(w, fmt.Sprintf("error getting all records: %v", err), http.StatusInternalServerError)
			}
			w.Header().Add("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(allData); err != nil {
				http.Error(w, fmt.Sprintf("error encoding response value to JSON: %v", err), http.StatusInternalServerError)
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
		}
	default:
		http.Error(w, "method must be GET or POST", http.StatusMethodNotAllowed)
		return
	}
}
