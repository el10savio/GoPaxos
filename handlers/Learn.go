package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Learn is the HTTP handler to process incoming Learn requests
// It persists the agreed upon value to its local KV Store
func Learn(w http.ResponseWriter, r *http.Request) {
	// Obtain the key & value from URL params
	key := mux.Vars(r)["key"]
	value := mux.Vars(r)["value"]

	err := Store.Set(key, value)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to set value")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.WithFields(log.Fields{
		"key":   key,
		"value": value,
	}).Debug("successful learn")

	w.WriteHeader(http.StatusOK)
}
