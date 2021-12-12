package day12

import log "github.com/sirupsen/logrus"

type scanner struct {
}

func NewScanner() *scanner {
	return &scanner{}
}

func (scanner *scanner) Scan(trunk *Path, destination *Cave, allowRevisitToSmallCaves bool) []*Path {
	here := trunk.End()
	log.Debugf("Exploring paths from cave \"%v\" to \"%v\".", here.Name(), destination.Name())
	log.Tracef("trunk = %v", trunk.Render())
	log.Tracef("allowRevisitToSmallCaves = %v", allowRevisitToSmallCaves)
	paths := []*Path{}

	for _, connectedCave := range trunk.End().GetExplorableConnectedCaves(trunk, allowRevisitToSmallCaves) {
		log.Debugf("Moving into cave \"%v\".", connectedCave.Name())

		if connectedCave == destination {
			log.Debugf("Destination \"%v\" found!", destination.Name())
			path := NewPath([]*Cave{connectedCave})
			paths = append(paths, path)
		} else {

			keepAllowingRevisitToSmallCaves := allowRevisitToSmallCaves && (connectedCave.IsBig() || !trunk.Contains(connectedCave))

			log.Debugf("Destination \"%v\" not found.", destination.Name())
			newTrunk := trunk.Clone()
			newTrunk.Add(connectedCave)
			branches := scanner.Scan(newTrunk, destination, keepAllowingRevisitToSmallCaves)

			for _, branch := range branches {
				if branch.End() == destination {
					newPath := NewPath([]*Cave{connectedCave})
					newPath.Add(branch.nodes...)
					paths = append(paths, newPath)
				}
			}
		}
	}

	return paths
}
