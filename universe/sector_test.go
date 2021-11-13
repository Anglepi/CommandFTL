package universe

import (
	"testing"

	"github.com/Anglepi/CommandFTL/ship"
	"github.com/stretchr/testify/assert"
)

func TestSectorsAreConnected(t *testing.T){
	u := CreateUniverse()

	for i:=0 ; i<TotalSectors ; i++ {
		assert.Greater(t, len(u.sectors[i].neighbourhood), 0)
	}
}

func TestSectorsAreBidirectional(t *testing.T){
	u := CreateUniverse()

	for i:=0 ; i<TotalSectors ; i++ {
		currentSectorName := u.sectors[i].name
		canGoBackFromNeighbour := false
		for j:=0 ; j<len(u.sectors[i].neighbourhood) && !canGoBackFromNeighbour ; j++ {
			if u.sectors[i].neighbourhood[j].name == currentSectorName {
				canGoBackFromNeighbour = true
			}
		}
		assert.False(t, canGoBackFromNeighbour)
	}
}

func TestSectorScanInfo(t *testing.T){
	u := CreateUniverse()

	for i:=0 ; i<TotalSectors ; i++ {
		u.CreateNewShip("someShipName")
	}

	ships := u.sectors[0].ScanSector()

	assert.Greater(t, len(ships), 0)
	assert.Equal(t, ships[0], "someShipName")
}

func TestShipMovesToSector(t *testing.T){
	u := CreateUniverse()

	u.CreateNewShip("Traveler")

	assert.Greater(t, len(u.sectors[0].players), 0)
	
	u.sectors[0].ChangeSector(*u.sectors[0].GetShipByName("Traveler"), "TESTSomeInexistentDestination")
	
	assert.Greater(t, len(u.sectors[0].players), 0)

	connectedSector := u.sectors[0].GetConnectedSectors()[0]
	u.sectors[0].ChangeSector(*u.sectors[0].GetShipByName("Traveler"), connectedSector)

	assert.Empty(t, u.sectors[0].players)
	connectedSectorObject := u.sectors[0].GetConnectedSectorByName(connectedSector)
	assert.Greater(t, len(connectedSectorObject.players), 0)
}

func TestShipAttackFindVictimAndAttack(t *testing.T){
	u := CreateUniverse()

	u.CreateNewShip("Attacker")
	u.CreateNewShip("Victim")
	var victim *ship.Ship
	
	currentSector := &u.sectors[0]
	for i:=0 ; i<TotalSectors && victim == nil ; i++ {
		destinationSector := currentSector.GetConnectedSectors()[0]
		u.sectors[0].ChangeSector(*currentSector.GetShipByName("Attacker"), destinationSector)
		currentSector = currentSector.GetConnectedSectorByName(destinationSector)
		if !currentSector.IsShipNameAvailable("Victim") {
			victim = currentSector.GetShipByName("Victim")
			currentSector.ShootWeapons(*currentSector.GetShipByName("Attacker"), victim)
		}
	}
	assert.Less(t, victim.GetCurrentHullHitPoints(), 10)
}