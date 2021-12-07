package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Object interface {
	Update(game interface{}) (shouldRemove bool, err error)
	Priority() int
	DrawOn(image *ebiten.Image)
	Layer() int
}

type BaseObject struct{}

func (b BaseObject) Update(game interface{}) (shouldRemove bool, err error) {
	return false, nil
}

func (b BaseObject) Priority() int {
	return 0
}

func (b BaseObject) DrawOn(_ *ebiten.Image) {}

func (b BaseObject) Layer() int {
	return 0
}
