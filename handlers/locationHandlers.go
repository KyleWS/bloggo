package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KyleWS/chikkin-server/models"
)

func (ctx *Context) LocationRequestHandler(w http.ResponseWriter, r *http.Request) {
	tokenParam := r.URL.Query()["token"]
	if tokenParam == nil {
		http.Error(w, fmt.Sprintf("error token query parameter must be present"), http.StatusBadRequest)
		return
	}
	token := tokenParam[0]
	if len(token) == 0 {
		http.Error(w, fmt.Sprintf("error token value must be assigned"), http.StatusBadRequest)
		return
	}
	if token != "chikkin" {
		http.Error(w, fmt.Sprintf("error bad token "), http.StatusBadRequest)
		return
	}
	switch r.Method {
	case "GET":
		// Called when user is polling endpoint for location data
		recentLocs, err := ctx.db.GetRecent()
		if err != nil {
			http.Error(w, fmt.Sprintf("error getting all records: %v", err), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(recentLocs); err != nil {
			http.Error(w, fmt.Sprintf("error encoding response value to JSON: %v", err), http.StatusInternalServerError)
			return
		}

	case "POST":
		// Call when user is posting new location in the form:
		// {"name":"data", "loc": "location data stringified"}
		locationPost := &models.LocationPost{}
		if err := json.NewDecoder(r.Body).Decode(locationPost); err != nil {
			http.Error(w, fmt.Sprintf("error decoding JSON: %v", err), http.StatusBadRequest)
			return
		}
		locationReturn, err := ctx.db.InsertLocation(locationPost)
		if err != nil {
			http.Error(w, fmt.Sprintf("error inserting task: %v", err), http.StatusInternalServerError)
			return
		}
		// set header to json so that it formats properly for the user who wants it
		w.Header().Add("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(locationReturn); err != nil {
			http.Error(w, fmt.Sprintf("error encoding response value to JSON: %v", err), http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, "method must be GET or POST", http.StatusMethodNotAllowed)
		return
	}
}
