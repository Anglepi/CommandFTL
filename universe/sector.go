package universe

import (
	"github.com/Anglepi/CommandFTL/ship"
)

type Sector struct {
	name string
	players []ship.Ship
	neighbourhood []Sector
}

func (sector *Sector) addNeighbours(neighbours []Sector) {
	sector.neighbourhood = append(sector.neighbourhood, neighbours...)

	for _, neighbour := range neighbours {
		neighbour.neighbourhood = append(neighbour.neighbourhood, *sector)
	}
}

func (sector *Sector) ScanSector(sectorName string) {

}

func (sector *Sector) ShotWeapons(attacker ship.Ship, victim ship.Ship) {

}

func (sector *Sector) ChangeSector(traveling ship.Ship, destinationSectorName string) {

}

func (sector *Sector) AddNewShip(shipName string) string {
	newShip := ship.CreateShip(shipName)
	sector.players = append(sector.players, *newShip)
	return newShip.GetName()
}

func (sector *Sector) RemoveShip(){

}

func (sector *Sector) AmountOfShips(){

}

