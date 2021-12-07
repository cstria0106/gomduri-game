package game

import (
	"github.com/cstria0106/gomduri/internal/asset"
	"github.com/cstria0106/gomduri/internal/game/engine"
	"github.com/hajimehoshi/ebiten/v2"
)

type PlusOne struct {
	engine.BaseObject

	x, y  float64
	frame int
}

func NewPlusOne(x, y float64) *PlusOne {
	return &PlusOne{x: x, y: y}
}

func (p *PlusOne) Update(_ interface{}) (shouldRemove bool, err error) {
	p.y -= 1
	p.frame += 1

	return p.frame == 60, nil
}

func (p *PlusOne) DrawOn(image *ebiten.Image) {
	o := &ebiten.DrawImageOptions{}
	o.GeoM.Translate(p.x, p.y)
	image.DrawImage(asset.PlusOneImage, o)
}

func (p *PlusOne) Layer() int {
	return 5
}
