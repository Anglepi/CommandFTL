package ship

type SystemCapabilities interface {
	executeCommand()
	passiveEffect()
}

type SystemBasics struct {
	name           string
	maxEnergy      int
	assignedEnergy int
	damagedEnergy  int
	capabilities   SystemCapabilities
}