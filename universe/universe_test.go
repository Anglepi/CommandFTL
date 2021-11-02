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

func TestSectorsAreConnected(t *testing.T){
	u := CreateUniverse()

	for i:=0 ; i<TotalSectors ; i++ {
		assert.Greater(t, len(u.sectors[i].neighbourhood), 0)
	}
}

func TestShipCreation(t *testing.T){
	u := CreateUniverse()

	for i:=0 ; i<TotalSectors ; i++ {
		var shipCreated, shipURL = CreateNewShip(u, "someShipName")
		assert.True(t, shipCreated)
		assert.Contains(t, shipURL, "/")
	}
	
	var shipCreated, shipURL = CreateNewShip(u, "someShipName")
	assert.False(t, shipCreated)
	assert.Empty(t, shipURL)
}