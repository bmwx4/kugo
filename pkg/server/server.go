package server

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/facebookgo/httpdown"
	"github.com/gorilla/mux"
)

// Controller is an interface that register handlers with a http router
type Controller interface {
	Register(router *mux.Router)
}

// Server is the api server interface
type Server interface {
	http.Handler
	Register(ctrl Controller)
	ListenAndServe() error
}

type server struct {
	listenAddr string
	prefix     string
	router     *mux.Router
	https      bool
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s == nil {
		panic("nil server")
	}
	s.router.ServeHTTP(w, r)
}

//NewHTTPServer xxx
func NewHTTPServer() Server {
	return &server{
		listenAddr: "0.0.0.0:8080",
		prefix:     "",
		router:     mux.NewRouter(),
	}
}

//NewHTTPSServer xxx
func NewHTTPSServer() Server {
	return &server{
		listenAddr: "0.0.0.0:6443",
		prefix:     "",
		https:      true,
		router:     mux.NewRouter(),
	}
}

func (s *server) Register(ctrl Controller) {
	if s == nil || ctrl == nil {
		return
	}
	ctrl.Register(s.router)
}

func (s *server) ListenAndServe() error {
	if s == nil {
		return errors.New("nil server")
	}
	var (
		tlsConfig *tls.Config
	)

	if s.https {
		cer, err := tls.LoadX509KeyPair("server.crt", "server.key")
		if err != nil {
			log.Fatal(err)
			return err
		}

		tlsConfig = &tls.Config{Certificates: []tls.Certificate{cer}}
	}

	httpServer := &http.Server{
		Addr:      s.listenAddr,
		Handler:   s.router,
		TLSConfig: tlsConfig,
	}

	hd := &httpdown.HTTP{
		StopTimeout: time.Second,
		KillTimeout: time.Second,
	}

	if err := httpdown.ListenAndServe(httpServer, hd); err != nil {
		fmt.Printf("listen and serve failed: %s", err)
		return err
	}
	return nil
}

func loadCA(caFile string) *x509.CertPool {
	pool := x509.NewCertPool()

	if ca, e := ioutil.ReadFile(caFile); e != nil {
		log.Fatal("ReadFile: ", e)
	} else {
		pool.AppendCertsFromPEM(ca)
	}
	return pool
}
