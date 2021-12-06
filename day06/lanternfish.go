package day06

type lanternfish struct {
	timer int
}

func (fish *lanternfish) Init(timer int) {
	fish.timer = timer
}

func (fish *lanternfish) Tick() *lanternfish {
	if fish.timer == 0 {
		fish.timer = 6
		return &lanternfish{
			timer: 8,
		}
	} else {
		fish.timer--
		return nil
	}
}
