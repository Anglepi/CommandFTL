package universe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniverseInitialization(t *testing.T){
	u:=CreateUniverse()

	assert.NotNil(t, u)
	assert.NotNil(t, u.sectors[0])
}

func TestShipCreation(t *testing.T){
	u := CreateUniverse()

	for i:=0 ; i<TotalSectors ; i++ {
		var shipCreated, shipURL = u.CreateNewShip("someShipName")
		assert.True(t, shipCreated)
		assert.Contains(t, shipURL, "/")
	}
	
	var shipCreated, shipURL = u.CreateNewShip("someShipName")
	assert.False(t, shipCreated)
	assert.Empty(t, shipURL)
}