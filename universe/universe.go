package universe

import "strconv"

const TotalSectors int = 9;

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
	return u
}

func CreateNewShip() {

}
