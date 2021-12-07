package main

import (
	"github.com/cstria0106/gomduri/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
	"time"
)

const ScreenScale = 3

func main() {
	rand.Seed(time.Now().UnixMilli())
	ebiten.SetWindowSize(
		game.GameWidth*ScreenScale,
		game.GameHeight*ScreenScale,
	)
	if err := ebiten.RunGame(game.NewGame()); err != nil {
		panic(err)
	}
}
