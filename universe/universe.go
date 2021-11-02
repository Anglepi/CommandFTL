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
	connectSectors(u)
	return u
}

func connectSectors(u *Universe){
	for i:=0 ; i<TotalSectors ; i++{
		u.sectors[i].neighbourhood = append(u.sectors[i].neighbourhood, u.sectors[(i+1)%TotalSectors], u.sectors[(i+1)%TotalSectors])
	}
}

func CreateNewShip() {

}
