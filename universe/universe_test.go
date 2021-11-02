package universe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniverseInitialization(t *testing.T){
	u:=CreateUniverse()

	assert.NotNil(t, u.sectors[0])
}

func TestSectorsAreConnected(t *testing.T){
	u:=CreateUniverse()

	assert.NotNil(t, u.sectors[0].neighbourhood[0])
}