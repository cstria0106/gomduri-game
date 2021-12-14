package game

import (
	"github.com/cstria0106/gomduri/internal/asset"
	"github.com/cstria0106/gomduri/internal/game/engine"
	"github.com/hajimehoshi/ebiten/v2"
)

type Background struct {
	engine.BaseObject
}

func NewBackground() *Background {
	return &Background{}
}

func (b *Background) DrawOn(image *ebiten.Image) {
	image.DrawImage(asset.BackgroundImage, &ebiten.DrawImageOptions{})
}

func (b *Background) Layer() int {
	return 100
}
