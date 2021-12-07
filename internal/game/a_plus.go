package game

import (
	"github.com/Tarliton/collision2d"
	"github.com/cstria0106/gomduri/internal/asset"
)

type APlus struct {
	*FallingItem
}

func NewAPlus() *APlus {
	box := collision2d.NewBox(collision2d.Vector{X: 4, Y: 0}, 19, 16)
	return &APlus{
		FallingItem: NewFallingItem(
			asset.APlusImage,
			box,
		),
	}
}

func (p *APlus) Update(gameInterface interface{}) (bool, error) {
	game := gameInterface.(*Game)

	if earned, err := p.FallingItem.Update(game); err != nil {
		return true, err
	} else if earned {
		game.engine.Add(NewPlusOne(p.x, p.y))
		return true, nil
	}

	return false, nil
}
