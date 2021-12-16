package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var server *Server

const StatusOK string = "OK"
const StatusERR string = "ERROR"

func executeRequest(request *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	server.Server.Handler.ServeHTTP(recorder, request)

	return recorder
}

func createShip(shipName string) string {
	req, _ := http.NewRequest("POST", "/newGame", bytes.NewBuffer([]byte(`{ "shipName": "`+shipName+`"}`)))

	response := executeRequest(req)
	var data map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &data)
	return data["msg"].(string)
}

func moveShip(traveler string, destination string) string {
	req, _ := http.NewRequest("PUT", "/"+traveler+"/engines", bytes.NewBuffer([]byte(`{"destination":"`+destination+`"}`)))

	response := executeRequest(req)

	var data map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &data)

	return data["msg"].(string)
}

func TestCreateNewPlayer(t *testing.T) {
	server = CreateServer()
	req, _ := http.NewRequest("POST", "/newGame", bytes.NewBuffer([]byte(`{ "shipName": "testShip"}`)))

	response := executeRequest(req)

	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, "application/json", response.Result().Header.Get("Content-Type"))

	var data map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &data)

	assert.Equal(t, data["status"], StatusOK)
	assert.Equal(t, data["msg"], "SCT0/testShip")
}

func TestNewPlayerNameAlreadyChosen(t *testing.T) {
	server = CreateServer()
	req, _ := http.NewRequest("POST", "/newGame", bytes.NewBuffer([]byte(`{ "shipName": "TestNewPlayerNameAlreadyChosen"}`)))
	req2, _ := http.NewRequest("POST", "/newGame", bytes.NewBuffer([]byte(`{ "shipName": "TestNewPlayerNameAlreadyChosen"}`)))

	executeRequest(req)
	response := executeRequest(req2)

	assert.Equal(t, http.StatusConflict, response.Code)
	assert.Equal(t, "application/json", response.Result().Header.Get("Content-Type"))

	var data map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &data)

	assert.Equal(t, data["status"], StatusERR)
	assert.Equal(t, data["msg"], "A ship already exists with that name")
}

func TestPlayerCanTravel(t *testing.T) {
	server = CreateServer()

	newShip := createShip("TestPlayerCanTravel")

	travelEndpoint := "/" + newShip + "/engines"
	req, _ := http.NewRequest("PUT", travelEndpoint, bytes.NewBuffer([]byte(`{"destination":"SCT1"}`)))

	response := executeRequest(req)

	assert.Equal(t, http.StatusOK, response.Code, "could not travel to "+travelEndpoint)
	assert.Equal(t, "application/json", response.Result().Header.Get("Content-Type"))

	var data map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &data)

	assert.Equal(t, data["status"], StatusOK)
	assert.Equal(t, data["msg"], "SCT1/TestPlayerCanTravel")
}

func TestShipCanShoot(t *testing.T) {
	server = CreateServer()

	attacker := createShip("attacker")
	createShip("victim")

	attacker = moveShip(attacker, "SCT1")

	req, _ := http.NewRequest("PUT", "/"+attacker+"/weapons", bytes.NewBuffer([]byte(`{"target":"victim"}`)))

	response := executeRequest(req)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "application/json", response.Result().Header.Get("Content-Type"))

	var data map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &data)

	assert.Equal(t, data["status"], StatusOK)
	assert.Equal(t, data["msg"], "You shoot your weapons at victim")
}
