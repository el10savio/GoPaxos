package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// KVPair ...
// type KVPair struct {
// 	Key   string `json:"key"`
// 	Value string `json:"value"`
// }

// SetValue ...
func SetValue(w http.ResponseWriter, r *http.Request) {
	// json decode the KV pair
	// var kv KVPair
	// _ = json.NewDecoder(r.Body).Decode(&kv)

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
	}).Debug("successfull learn")

	w.WriteHeader(http.StatusOK)
}
