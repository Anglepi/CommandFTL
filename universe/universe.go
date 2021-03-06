package universe

import "strconv"

const TotalSectors = 9

type Universe struct {
	sectors [TotalSectors]Sector
}

func CreateUniverse() *Universe {
	u := new(Universe)

	for i := 0; i < TotalSectors; i++ {
		newSector := new(Sector)
		newSector.name = "SCT" + strconv.Itoa(i)
		u.sectors[i] = *newSector
	}

	connectSectors(u)

	return u
}

func connectSectors(u *Universe) {
	for i := 0; i < TotalSectors; i++ {
		var neighbours = [2]*Sector{&u.sectors[(i+1)%TotalSectors], &u.sectors[(i+2)%TotalSectors]}
		u.sectors[i].addNeighbours(neighbours[:])
		u.sectors[i].neighbourhood = append(u.sectors[i].neighbourhood, &u.sectors[(i+1)%TotalSectors], &u.sectors[(i+1)%TotalSectors])
	}
}

func (universe *Universe) CreateNewShip(shipName string) (bool, string) {
	var shipIntroduced = false
	var shipURL = ""

	if universe.isShipNameAvailable(shipName) {
		for i := 0; i < TotalSectors && !shipIntroduced; i++ {
			if len(universe.sectors[i].players) < 1 {
				//TBD generate API endpoint and asign to shipURL
				shipURL = universe.sectors[i].name + "/" + universe.sectors[i].AddNewShip(shipName) 
				shipIntroduced = true
			}
		}
	}

	return shipIntroduced, shipURL
}

func (universe *Universe) GetSectorByName(sectorName string) *Sector {
	var sector *Sector
	sectorFound := false

	for i := 0; i < TotalSectors && !sectorFound; i++ {
		sectorFound = sectorName == universe.sectors[i].name
		if sectorFound {
			sector = &universe.sectors[i]
		}
	}

	return sector
}

func (universe *Universe) HasShipInSector(shipName string, sectorName string) bool {
	sector := universe.GetSectorByName(sectorName)
	if sector != nil {
		return !sector.IsShipNameAvailable(shipName)
	} else {
		return false
	}
}

func (universe *Universe) isShipNameAvailable(shipName string) bool {
	nameAvailable := true

	for i := 0; i < TotalSectors && nameAvailable; i++ {
		nameAvailable = universe.sectors[i].IsShipNameAvailable(shipName)
	}

	return nameAvailable
}
