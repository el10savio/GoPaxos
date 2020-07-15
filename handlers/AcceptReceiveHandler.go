package acceptor

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// AcceptReceiveHandler ...
func AcceptReceiveHandler(w http.ResponseWriter, r *http.Request) {
	// Obtain the id from URL params
	id := mux.Vars(r)["id"]

	accepted, err := AcceptReceive(id)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to prepare")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Debug("successfull accept")

	if !accepted {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
