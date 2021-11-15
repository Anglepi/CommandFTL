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

	//HandleFTL
	router.HandleFunc("/action/ftl/{sector}/{shipName}/{destination}", func(w http.ResponseWriter, r *http.Request) {
		// Params

		vars := mux.Vars(r)
		sector := server.Universe.GetSectorByName(vars["sector"])
		ship := sector.GetShipByName(vars["shipName"])

		if server.Universe.HasShipInSector(vars["shipName"], vars["sector"]) {
			changeSuccessful := sector.ChangeSector(*ship, vars["destination"])
			if changeSuccessful {
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(Response{"OK", "You are now at sector " + vars["destination"]})
			} else {
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(Response{"ERROR", "Destination sector is not within range from your location"})
			}

		} else {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{"ERROR", "Ship name or current sector incorrect"})
		}

	}).Methods("POST")

	//HandleShootWeapons
	router.HandleFunc("/action/shoot/{sector}/{shipName}/{target}", func(w http.ResponseWriter, r *http.Request) {
		// Params
		//vars := mux.Vars(r)
	})
}
