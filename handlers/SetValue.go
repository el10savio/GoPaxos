package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// KVPair ...
type KVPair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// SetValue ...
func SetValue(w http.ResponseWriter, r *http.Request) {
	// json decode the KV pair
	var kv KVPair
	_ = json.NewDecoder(r.Body).Decode(&kv)

	err := Store.Set(kv.Key, kv.Value)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to set value")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.WithFields(log.Fields{
		"key":   kv.Key,
		"value": kv.Value,
	}).Debug("successfull SetValue")

	w.WriteHeader(http.StatusOK)
}
