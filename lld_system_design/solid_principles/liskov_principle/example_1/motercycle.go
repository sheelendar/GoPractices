package example1

type MoterCycle struct {
	isEngineOn bool
	speed      int
}

func GetMoterCycle(isEngineOn bool, speed int) *MoterCycle {
	return &MoterCycle{isEngineOn: isEngineOn, speed: speed}
}

func (moterCycle *MoterCycle) TurnOnEngine() {
	moterCycle.isEngineOn = true
}

func (moterCycle *MoterCycle) Accelerate() {
	moterCycle.speed += 10
}
