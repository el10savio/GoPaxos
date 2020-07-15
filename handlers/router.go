package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"../acceptor"
	"../kvstore"
	"../proposer"
)

var (
	// Store ...
	Store kvstore.Store
)

func init() {
	Store = kvstore.Initialize()
}

// Route defines the Mux
// router individual route
type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

// Routes is a collection
// of individual Routes
var Routes = []Route{
	Route{"/", "GET", Index},
	Route{"/store/get/{key}", "GET", GetValue},
	Route{"/store/set/{key}/{value}", "GET", proposer.PaxosSetValue},
	Route{"/prepare/{id}", "GET", acceptor.PrepareReceiveHandler},
	Route{"/accept/{id}", "GET", acceptor.AcceptReceiveHandler},
	Route{"/learn/{key}/{value}", "GET", Learn},
}

// Index is the handler for the path "/"
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World Paxos Server\n")
}

// Logger is the middleware to
// log the incoming request
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"path":   r.URL,
			"method": r.Method,
		}).Info("incoming request")

		next.ServeHTTP(w, r)
	})
}

// Router returns a mux router
func Router() *mux.Router {
	router := mux.NewRouter()

	for _, route := range Routes {
		router.HandleFunc(
			route.Path,
			route.Handler,
		).Methods(route.Method)
	}

	router.Use(Logger)

	return router
}
