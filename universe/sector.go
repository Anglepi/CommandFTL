package universe

import (
	"github.com/Anglepi/CommandFTL/ship"
)

type Sector struct {
	name string
	players []ship.Ship
	neighbourhood []Sector
}

func (sector *Sector) ScanSector(sectorName string) {

}

func (sector *Sector) ShotWeapons(attacker ship.Ship, victim ship.Ship) {

}

func (sector *Sector) ChangeSector(traveling ship.Ship, destinationSectorName string) {

}

func (sector *Sector) AddNewShip() {

}

func (sector *Sector) RemoveShip(){

}

func (sector *Sector) AmountOfShips(){

}

