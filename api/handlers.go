package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

func InitializeHandlers(router *mux.Router, server *Server) {
	//HandleShipName
	router.HandleFunc("/newGame", func(w http.ResponseWriter, r *http.Request) {
		// Params
		type Request struct {
			ShipName string `json:"shipName"`
		}
		var request Request

		err := json.NewDecoder(r.Body).Decode(&request)
		shipName := request.ShipName

		if err != nil || shipName == "" {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			var success, address = server.Universe.CreateNewShip(shipName)
			w.Header().Set("Content-Type", "application/json")

			if !success {
				w.WriteHeader(http.StatusConflict)
				json.NewEncoder(w).Encode(Response{"ERROR", "A ship already exists with that name"})
			} else {
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(Response{"OK", address})
			}
		}
	}).Methods("POST")

	//HandleFTL
	router.HandleFunc("/action/ftl/{shipName}/{sector}", func(w http.ResponseWriter, r *http.Request) {
		type Request struct {
			Destination string `json:"destination"`
		}
		var request Request

		vars := mux.Vars(r)
		fmt.Println(vars)
		sector := server.Universe.GetSectorByName(vars["sector"])
		if sector != nil {
			ship := sector.GetShipByName(vars["shipName"])
			err := json.NewDecoder(r.Body).Decode(&request)
			destination := request.Destination

			if server.Universe.HasShipInSector(vars["shipName"], vars["sector"]) {
				if err != nil || destination == "" {
					w.WriteHeader(http.StatusBadRequest)
				} else {
					changeSuccessful := sector.ChangeSector(*ship, destination)
					if changeSuccessful {
						w.WriteHeader(http.StatusOK)
						json.NewEncoder(w).Encode(Response{"OK", ship.GetName() + "/" + destination})
					} else {
						w.WriteHeader(http.StatusUnprocessableEntity)
						json.NewEncoder(w).Encode(Response{"ERROR", "Destination sector is not within range from your location"})
					}
				}

			} else {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(Response{"ERROR", "Ship name or current sector incorrect"})
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}

	}).Methods("PUT")

	//HandleShootWeapons
	router.HandleFunc("/action/shoot/{shipName}/{sector}", func(w http.ResponseWriter, r *http.Request) {
		// Params
		type Request struct {
			Target string `json:"target"`
		}
		var request Request
		err := json.NewDecoder(r.Body).Decode(&request)

		vars := mux.Vars(r)
		sector := server.Universe.GetSectorByName(vars["sector"])
		fmt.Println("something")
		if sector != nil {
			ship := sector.GetShipByName(vars["shipName"])
			if err != nil || request.Target == "" {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				target := sector.GetShipByName(request.Target)

				if server.Universe.HasShipInSector(vars["shipName"], vars["sector"]) {
					if server.Universe.HasShipInSector(request.Target, vars["sector"]) {
						sector.ShootWeapons(*ship, target)
						w.WriteHeader(http.StatusOK)
						json.NewEncoder(w).Encode(Response{"OK", "You shoot your weapons at " + request.Target})
					} else {
						w.WriteHeader(http.StatusUnprocessableEntity)
						json.NewEncoder(w).Encode(Response{"ERROR", "Your target does not exist"})
					}
				} else {
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode(Response{"ERROR", "Ship name or current sector incorrect"})
				}
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}).Methods("PUT")
}
