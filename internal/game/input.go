package game

import (
	"github.com/cstria0106/gomduri/internal/game/engine"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"math"
)

type Input struct {
	engine.BaseObject

	Left         bool
	LeftPressed  bool
	Right        bool
	RightPressed bool
	Jump         bool
	JumpPressed  bool
}

func (i *Input) Update(_ interface{}) (bool, error) {
	i.Left = ebiten.IsKeyPressed(ebiten.KeyLeft)
	i.LeftPressed = inpututil.IsKeyJustPressed(ebiten.KeyLeft)

	i.Right = ebiten.IsKeyPressed(ebiten.KeyRight)
	i.RightPressed = inpututil.IsKeyJustPressed(ebiten.KeyRight)

	i.Jump = ebiten.IsKeyPressed(ebiten.KeyZ)
	i.JumpPressed = inpututil.IsKeyJustPressed(ebiten.KeyZ)

	return false, nil
}

func (i *Input) Priority() int {
	return math.MaxInt
}
