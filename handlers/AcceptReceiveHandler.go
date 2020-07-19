package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"../acceptor"
)

// AcceptReceiveHandler is the HTTP handler to process incoming Accept requests
func AcceptReceiveHandler(w http.ResponseWriter, r *http.Request) {
	// Obtain the id from URL params
	id := mux.Vars(r)["id"]

	accepted, err := acceptor.AcceptReceive(id)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to accept")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Debug("successful accept")

	if !accepted {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
