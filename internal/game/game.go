package game

import (
	"github.com/cstria0106/gomduri/internal/asset"
	"github.com/cstria0106/gomduri/internal/game/engine"
	"github.com/hajimehoshi/ebiten/v2"
)

const GameWidth = 320
const GameHeight = 240

type Game struct {
	engine *engine.Engine

	input   *Input
	gomduri *Gomduri
	items   []engine.Object
}

func NewGame() *Game {
	game := &Game{
		engine:  engine.NewEngine(),
		input:   &Input{},
		gomduri: nil,
		items:   []engine.Object{},
	}

	game.gomduri = NewGomduri(asset.GomduriImage)

	game.engine.Add(
		game.input,
		game.gomduri,
		NewSpawnItemSystem(),
	)

	return game
}

func (g *Game) Update() error {
	var err error

	if err = g.engine.Update(g); err != nil {
		return err
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.engine.Draw(screen)
}

func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return GameWidth, GameHeight
}
