package day12

type Cave struct {
	name           string
	connectedCaves []*Cave
}

func NewCave(name string) *Cave {
	return &Cave{
		name: name,
	}
}

func (cave *Cave) Name() string {
	return cave.name
}

func (cave *Cave) ConnectsTo(other *Cave) bool {
	found := false

	for _, connectedCave := range cave.connectedCaves {
		if connectedCave == other {
			found = true
			break
		}
	}

	return found
}

func (cave *Cave) Connect(other *Cave) {
	if !cave.ConnectsTo(other) {
		cave.connectedCaves = append(cave.connectedCaves, other)
	}
}

func (cave *Cave) ConnectedCaves() []*Cave {
	return cave.connectedCaves
}

func (cave *Cave) IsBig() bool {
	return cave.name[0] >= 65 && cave.name[0] <= 90
}

func (cave *Cave) GetExplorableConnectedCaves(trunk *Path, allowRevisitToSmallCaves bool) []*Cave {
	branches := []*Cave{}

	for _, connectedCave := range cave.ConnectedCaves() {
		if connectedCave != cave && (allowRevisitToSmallCaves || !trunk.Contains(connectedCave) || connectedCave.IsBig()) {
			branches = append(branches, connectedCave)
		}
	}

	return branches
}
