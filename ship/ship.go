package ship

type Ship struct {
	name            string
	hullHitPoints   int
	energyAvailable int
	systems         []SystemBasics
}

func CreateShip() *Ship {
	return &Ship{hullHitPoints: 10, energyAvailable: 20}
}

func (ship *Ship) ShootWeapon(target Ship) {

}

func (ship *Ship) CanMoveToSector(destination string) {

}
