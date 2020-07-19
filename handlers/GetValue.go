package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// GetValue  is the HTTP handler to process incoming KV Get requests
// It gets the value from the in memory KV store
func GetValue(w http.ResponseWriter, r *http.Request) {
	// Obtain the key from URL params
	key := mux.Vars(r)["key"]

	value, err := Store.Get(key)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to get value")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.WithFields(log.Fields{
		"key":   key,
		"value": value,
	}).Debug("successful getvalue")

	// json encode response value
	json.NewEncoder(w).Encode(value)
}
