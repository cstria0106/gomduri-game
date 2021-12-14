package game

import (
	"fmt"
	"github.com/cstria0106/gomduri/internal/asset"
	"github.com/cstria0106/gomduri/internal/game/engine"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"log"
)

type Score struct {
	engine.BaseObject
	score    int
	fontFace font.Face
}

func NewScore() *Score {
	face, err := opentype.NewFace(
		asset.GalmuriFont,
		&opentype.FaceOptions{
			Size:    16,
			DPI:     72,
			Hinting: font.HintingNone,
		},
	)

	if err != nil {
		log.Println("failed to create face from font:", err)
	}

	return &Score{
		score:    0,
		fontFace: face,
	}
}

func (d *Score) Add(score int) {
	d.score += score
}

func (d *Score) DrawOn(image *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.Filter = ebiten.FilterNearest

	scoreString := fmt.Sprintf("점수: %d", d.score)
	bound := text.BoundString(d.fontFace, scoreString)

	options.GeoM.Translate(float64(GameWidth/2-bound.Dx()/2), 15)
	options.GeoM.Translate(-1, 0)
	text.DrawWithOptions(image, scoreString, d.fontFace, options)
	options.GeoM.Translate(1, 1)
	text.DrawWithOptions(image, scoreString, d.fontFace, options)
	options.GeoM.Translate(0, -2)
	text.DrawWithOptions(image, scoreString, d.fontFace, options)
	options.GeoM.Translate(1, 1)
	text.DrawWithOptions(image, scoreString, d.fontFace, options)

	options.ColorM.Scale(0, 0, 0, 1)
	options.GeoM.Translate(-1, 0)
	text.DrawWithOptions(image, scoreString, d.fontFace, options)
}

func (d *Score) Layer() int {
	return -10
}
