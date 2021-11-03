package universe

import "strconv"

const TotalSectors = 9;

type Universe struct {
	sectors  [TotalSectors]Sector
}

func CreateUniverse() *Universe {
	u := new(Universe)

	for i:=0 ; i<TotalSectors ; i++ {
		newSector := new(Sector)
		newSector.name = "SCT"+strconv.Itoa(i)
		u.sectors[i] = *newSector
	}
	connectSectors(u)
	return u
}

func connectSectors(u *Universe){
	for i:=0 ; i<TotalSectors ; i++{
		var neighbours = [2]Sector{u.sectors[(i+1)%TotalSectors], u.sectors[(i+1)%TotalSectors]}
		u.sectors[i].addNeighbours(neighbours[:])
		u.sectors[i].neighbourhood = append(u.sectors[i].neighbourhood, u.sectors[(i+1)%TotalSectors], u.sectors[(i+1)%TotalSectors])
	}
}

func (universe *Universe) CreateNewShip(shipName string) (bool, string) {
	var shipIntroduced = false
	var shipURL = ""
	//TBD check for ships with the same name within the universe
	for i:=0 ; i<TotalSectors && !shipIntroduced ; i++{
		if len(universe.sectors[i].players) < 1 {
			//TBD generate API endpoint and asign to shipURL
			shipURL = universe.sectors[i].name + "/" + universe.sectors[i].AddNewShip(shipName)
			shipIntroduced = true
		}
	}
	return shipIntroduced, shipURL
}
