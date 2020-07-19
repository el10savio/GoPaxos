package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"../acceptor"
)

// PrepareReceiveHandler is the HTTP handler
// to process incoming Prepare Requests
func PrepareReceiveHandler(w http.ResponseWriter, r *http.Request) {
	// Obtain the id from URL params
	id := mux.Vars(r)["id"]

	prepared, err := acceptor.PrepareReceive(id)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to prepare")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Debug("successful prepare")

	if !prepared {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
