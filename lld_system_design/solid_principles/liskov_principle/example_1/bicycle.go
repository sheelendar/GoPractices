package example1

type BiCycle struct {
	speed int
}

func GetBiCycle(speed int) *BiCycle {
	return &BiCycle{speed: speed}
}

func (moterCycle *BiCycle) TurnOnEngine() {
	panic("there is not engine")
}

func (moterCycle *BiCycle) Accelerate() {
	moterCycle.speed += 10
}
