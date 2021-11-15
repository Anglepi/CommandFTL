package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

func InitializeHandlers(router *mux.Router, server *Server) {
	//HandleShipName
	router.HandleFunc("/new/{shipName}", func(w http.ResponseWriter, r *http.Request) {
		// Params
		vars := mux.Vars(r)
		shipName := vars["shipName"]

		if shipName == "" {
			w.WriteHeader(http.StatusUnprocessableEntity)
		} else {
			var success, address = server.Universe.CreateNewShip(shipName)
			w.Header().Set("Content-Type", "application/json")

			if !success {
				w.WriteHeader(http.StatusConflict)
				json.NewEncoder(w).Encode(Response{"ERROR", "A ship already exists with that name"})
			} else {
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(Response{"OK", address})
			}
		}
	}).Methods("POST")

}
