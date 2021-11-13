package universe

import (
	"testing"

	"strconv"

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
		var shipCreated, shipURL = u.CreateNewShip("someShipName"+strconv.Itoa(i))
		assert.True(t, shipCreated)
		assert.Contains(t, shipURL, "/")
	}
	
	var shipCreated, shipURL = u.CreateNewShip("someShipName")
	assert.False(t, shipCreated)
	assert.Empty(t, shipURL)
}

func TestNoDuplicateShipNames(t *testing.T){
	u := CreateUniverse()

	u.CreateNewShip("ShipA")
	secondShipCreated, _ := u.CreateNewShip("ShipA")

	assert.False(t, secondShipCreated)
}

func TestConnectionBetweenSectors(t *testing.T){
	u := CreateUniverse()

	for i:=0 ; i<len(u.sectors) ; i++{
		currentSector := u.sectors[i]
		for j:=0 ; j<len(currentSector.neighbourhood) ; j++{
			assert.False(t, currentSector.name == currentSector.neighbourhood[j].name)
		}
	}
}