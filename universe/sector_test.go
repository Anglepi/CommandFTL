package universe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSectorsAreConnected(t *testing.T){
	u := CreateUniverse()

	for i:=0 ; i<TotalSectors ; i++ {
		assert.Greater(t, len(u.sectors[i].neighbourhood), 0)
	}
}

func TestSectorsAreBidirectional(t *testing.T){
	u := CreateUniverse()

	for i:=0 ; i<TotalSectors ; i++ {
		currentSectorName := u.sectors[i].name
		canGoBackFromNeighbour := false
		for j:=0 ; j<len(u.sectors[i].neighbourhood) && !canGoBackFromNeighbour ; j++ {
			if u.sectors[i].neighbourhood[j].name == currentSectorName {
				canGoBackFromNeighbour = true
			}
		}
		assert.False(t, canGoBackFromNeighbour);
	}
}
