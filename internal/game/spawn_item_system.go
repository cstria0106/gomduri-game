package game

import (
	"github.com/cstria0106/gomduri/internal/game/engine"
	"math"
)

type SpawnItemSystem struct {
	engine.BaseObject

	frame int
}

func NewSpawnItemSystem() *SpawnItemSystem {
	return &SpawnItemSystem{
		frame: 0,
	}
}

func (s *SpawnItemSystem) Update(gameInterface interface{}) (bool, error) {
	game := gameInterface.(*Game)

	s.frame += 1

	if s.frame%20 == 0 {
		game.engine.Add(NewAPlus())
	}

	return false, nil
}

func (s *SpawnItemSystem) Priority() int {
	return math.MaxInt
}
