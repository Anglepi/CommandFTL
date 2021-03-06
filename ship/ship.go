package ship

type Ship struct {
	name            string
	hullHitPoints   int
	energyAvailable int
	systems         []SystemBasics
}

func CreateShip(shipName string) *Ship {
	return &Ship{name: shipName, hullHitPoints: 10, energyAvailable: 20}
}

func (ship *Ship) GetName() string {
	return ship.name
}

func (ship *Ship) GetCurrentHullHitPoints() int {
	return ship.hullHitPoints
}

func (ship *Ship) ShootWeapon(target *Ship) {
	//TBD weapon precission, weapon characteristics
	target.hullHitPoints--
}

func (ship *Ship) IsDestroyed() bool {
	return ship.hullHitPoints <= 0
}
