package day21

import log "github.com/sirupsen/logrus"

type Player struct {
	id       int
	score    int
	position int
}

func NewPlayer(id int, startingPosition int) *Player {
	return &Player{
		id:       id,
		position: startingPosition,
	}
}

func (player *Player) Score() int {
	return player.score
}

func (player *Player) TakeTurn(dice Die) {
	log.Debugf("Player %d is taking their turn.", player.id)
	log.Tracef("Starting Position = %d", player.position)
	log.Tracef("Starting Score = %d", player.score)

	roll := dice.RollN(3)

	if roll > 10 {
		roll = roll % 10
	}

	player.position = player.position + roll

	for player.position > 10 {
		player.position -= 10
	}

	player.score += player.position

	log.Debugf("Ending Position = %d", player.position)
	log.Debugf("Ending Score = %d", player.score)
}
