package main

// The following implements the main Go
// package starting up the paxos server

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"./handlers"
)

const (
	// PORT defines the port value
	// for the Paxos Server service
	PORT = "8080"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	r := handlers.Router()

	log.WithFields(log.Fields{
		"port": PORT,
	}).Info("starting paxos server")

	http.ListenAndServe(":"+PORT, r)
}
