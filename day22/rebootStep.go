package day22

type RebootStep struct {
	cuboid *Cuboid
	on     bool
}

func NewRebootStep(cuboid *Cuboid, on bool) *RebootStep {
	return &RebootStep{
		cuboid: cuboid,
		on:     on,
	}
}
