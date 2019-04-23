package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Controller struct {
	Prefix  string
	Version string
}

// New will create a new network api controller
func New() *Controller {
	return &Controller{
		Prefix:  "",
		Version: "v1",
	}
}

//Register XX
func (c *Controller) Register(router *mux.Router) {
	version := func(uri string) string {
		return fmt.Sprintf("%s/%s%s", c.Prefix, c.Version, uri)
	}
	router.Methods("GET").Path(version("/")).HandlerFunc(c.Home)
}

//Home  xxx
func (c *Controller) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
	return
}
