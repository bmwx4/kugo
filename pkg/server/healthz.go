package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type healthz struct{}

// Healthz controller example
var Healthz Controller = &healthz{}

func (h *healthz) Register(router *mux.Router) {
	subrouter := router.Path("/healthz").Subrouter()
	subrouter.Methods("GET").HandlerFunc(h.check)
}

func (h *healthz) check(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}
