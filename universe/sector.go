package universe

import (
	"github.com/Anglepi/CommandFTL/ship"
)

type Sector struct {
	name          string
	players       []ship.Ship
	neighbourhood []*Sector
}

func (sector *Sector) addNeighbours(neighbours []*Sector) {
	sector.neighbourhood = append(sector.neighbourhood, neighbours...)

	for _, neighbour := range neighbours {
		neighbour.neighbourhood = append(neighbour.neighbourhood, sector)
	}
}

func (sector *Sector) ScanSector() []string {
	totalShipsInSector := len(sector.players)
	var shipNames []string

	for i := 0; i < totalShipsInSector; i++ {
		shipNames = append(shipNames, sector.players[i].GetName())
	}

	return shipNames
}

func (sector *Sector) ShootWeapons(attacker ship.Ship, victim *ship.Ship) {
	attacker.ShootWeapon(victim)
}

func (sector *Sector) ChangeSector(travelingShip ship.Ship, destinationSectorName string) bool {
	sectorExists := false
	var neighbour *Sector

	for i := 0; i < len(sector.neighbourhood) && !sectorExists; i++ {
		sectorExists = sector.neighbourhood[i].name == destinationSectorName
		if sectorExists {
			neighbour = sector.neighbourhood[i]
		}
	}

	if sectorExists {
		sector.removeShip(travelingShip)
		neighbour.addExistingShip(travelingShip)
	}
	return sectorExists
}

func (sector *Sector) addExistingShip(ship ship.Ship) {
	sector.players = append(sector.players, ship)
}

func (sector *Sector) AddNewShip(shipName string) string {
	newShip := ship.CreateShip(shipName)
	sector.players = append(sector.players, *newShip)
	return newShip.GetName()
}

func (sector *Sector) removeShip(ship ship.Ship) {
	shipFound := false

	for i := 0; i < len(sector.players) && !shipFound; i++ {
		shipFound = sector.players[i].GetName() == ship.GetName()
		if shipFound {
			sector.players[i] = sector.players[len(sector.players)-1]
			sector.players = sector.players[:len(sector.players)-1]
		}
	}
}

func (sector *Sector) IsShipNameAvailable(shipName string) bool {
	nameAvailable := true

	for i := 0; i < len(sector.players) && nameAvailable; i++ {
		nameAvailable = sector.players[i].GetName() != shipName
	}

	return nameAvailable
}

func (sector *Sector) GetConnectedSectors() []string {
	var connectedSectors []string

	for i := 0; i < len(sector.neighbourhood); i++ {
		connectedSectors = append(connectedSectors, sector.neighbourhood[i].name)
	}

	return connectedSectors
}

func (sector *Sector) GetShipByName(shipName string) *ship.Ship {
	var ship *ship.Ship

	for i := 0; i < len(sector.players); i++ {
		if sector.players[i].GetName() == shipName {
			ship = &sector.players[i]
		}
	}

	return ship
}

func (sector *Sector) GetConnectedSectorByName(sectorName string) *Sector {
	var neighbour *Sector

	for i := 0; i < len(sector.neighbourhood) && neighbour == nil; i++ {
		if sector.neighbourhood[i].name == sectorName {
			neighbour = sector.neighbourhood[i]
		}
	}

	return neighbour
}
