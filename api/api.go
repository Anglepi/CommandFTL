package api

import (
	"fmt"
	"net/http"

	"github.com/Anglepi/CommandFTL/universe"
	"github.com/gorilla/mux"
)

type Server struct {
	Server   *http.Server
	Universe *universe.Universe
}

func CreateServer() *Server {
	server := new(Server)
	server.Universe = universe.CreateUniverse()

	router := mux.NewRouter()

	InitializeHandlers(router, server)

	// Log middleware
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.Method + " " + r.RequestURI + " from " + r.RemoteAddr)
			next.ServeHTTP(w, r)
		})
	})

	server.Server = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,
	}

	return server
}

func (server *Server) Start() {
	server.Server.ListenAndServe()
}
