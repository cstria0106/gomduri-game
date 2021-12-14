package game

import (
	"github.com/cstria0106/gomduri/internal/game/engine"
	"math"
)

type Spawner struct {
	engine.BaseObject

	frame int
}

func NewSpawner() *Spawner {
	return &Spawner{
		frame: 0,
	}
}

func (s *Spawner) Update(gameInterface interface{}) (bool, error) {
	game := gameInterface.(*Game)

	s.frame += 1

	if s.frame%20 == 0 {
		game.engine.Add(NewAPlus())
	}

	return false, nil
}

func (s *Spawner) Priority() int {
	return math.MaxInt
}
