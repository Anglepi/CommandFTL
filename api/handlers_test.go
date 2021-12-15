package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var server *Server = CreateServer()

const StatusOK string = "OK"


func executeRequest(request *http.Request) *httptest.ResponseRecorder{
	recorder := httptest.NewRecorder()
	server.Server.Handler.ServeHTTP(recorder, request)

	return recorder
}

func TestCreateNewPlayer(t *testing.T){
	req, _ := http.NewRequest("POST", "/newGame", bytes.NewBuffer([]byte(`{ "shipName": "testShip"}`)))
	
	response := executeRequest(req)

	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, "application/json", response.HeaderMap.Get("Content-Type"))

	var data map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &data)

	assert.Equal(t, data["status"], StatusOK)
	assert.Equal(t, data["msg"], "SCT0/testShip")
}