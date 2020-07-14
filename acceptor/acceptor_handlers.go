package acceptor

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// PrepareReceiveHandler ...
func PrepareReceiveHandler(w http.ResponseWriter, r *http.Request) {
	// Obtain the id from URL params
	id := mux.Vars(r)["id"]

	PrepareReceive(id)

	// err := PrepareReceive(id)
	// if err != nil {
	// 	log.WithFields(log.Fields{"error": err}).Error("failed to prepare")
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	log.WithFields(log.Fields{
		"id": id,
	}).Debug("successfull prepare")

	w.WriteHeader(http.StatusOK)
}

// AcceptReceiveHandler ...
func AcceptReceiveHandler(w http.ResponseWriter, r *http.Request) {
	// Obtain the id & value from URL params
	id := mux.Vars(r)["id"]
	value := mux.Vars(r)["value"]

	AcceptReceive(id)

	// err := AcceptReceive(id, value)
	// if err != nil {
	// 	log.WithFields(log.Fields{"error": err}).Error("failed to accept")
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	log.WithFields(log.Fields{
		"id":    id,
		"value": value,
	}).Debug("successfull accept")

	w.WriteHeader(http.StatusOK)
}

// LearnReceiveHandler ...
func LearnReceiveHandler(w http.ResponseWriter, r *http.Request) {
	// Obtain the value from URL params
	value := mux.Vars(r)["value"]

	// LearnReceive(value)

	// err := LearnReceive(value)
	// if err != nil {
	// 	log.WithFields(log.Fields{"error": err}).Error("failed to learn")
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	log.WithFields(log.Fields{
		"value": value,
	}).Debug("successfull learn")

	w.WriteHeader(http.StatusOK)
}