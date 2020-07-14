package proposer

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// SetValue ...
func SetValue(w http.ResponseWriter, r *http.Request) {
	// Obtain the key & value from URL params
	key := mux.Vars(r)["key"]
	value := mux.Vars(r)["value"]

	err := Prepare(value)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to set value")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.WithFields(log.Fields{
		"key":   key,
		"value": value,
	}).Debug("successfull set value")

	w.WriteHeader(http.StatusOK)
}
