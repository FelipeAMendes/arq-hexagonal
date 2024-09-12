package server

import (
	"github.com/felipeamendes/arq-hexagonal/adapters/web/handler"
	"github.com/felipeamendes/arq-hexagonal/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

type Webserver struct {
	Service application.ProductServiceInterface
}

func MakeNewWebserver() *Webserver {
	return &Webserver{}
}

func (w Webserver) Serve() {
	router := mux.NewRouter()
	negroni := negroni.New(
		negroni.NewLogger(),
	)
	handler.MakeProductHandlers(router, negroni, w.Service)
	http.Handle("/", router)
	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
