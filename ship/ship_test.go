package ship

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShipInitialization(t *testing.T){
	shipName := "Awesome Name"
	s := CreateShip(shipName)

	assert.NotNil(t, s)
	assert.Equal(t, s.hullHitPoints, 10)
	assert.Equal(t, s.energyAvailable, 20)
	assert.Equal(t, s.name, shipName)
}

func TestShipShootWeapons(t *testing.T){
	s1 := CreateShip("Awesome Name")
	s2 := CreateShip("Daedalus")

	s1.ShootWeapon(s2)

	//Once I have weapons implemented, I will check the exact amount of damage dealt
	assert.Less(t, s2.hullHitPoints, 10)
}