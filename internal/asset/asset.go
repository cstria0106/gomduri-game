package asset

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font/opentype"
	"image"
	_ "image/png"
)

//go:embed images/gomduri.png
var gomduriImageFile []byte
var GomduriImage *ebiten.Image

//go:embed images/a_plus.png
var aPlusImageFile []byte
var APlusImage *ebiten.Image

//go:embed images/cat.png
var catImageFile []byte
var CatImage *ebiten.Image

//go:embed images/f.png
var fImageFile []byte
var FImage *ebiten.Image

//go:embed images/plus_one.png
var plusOneImageFile []byte
var PlusOneImage *ebiten.Image

//go:embed images/background.png
var backgroundImageFile []byte
var BackgroundImage *ebiten.Image

//go:embed fonts/galmuri9.ttf
var galmuriFontFile []byte
var GalmuriFont *opentype.Font

func getImage(source []byte) *ebiten.Image {
	i, _, err := image.Decode(bytes.NewReader(source))
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(i)
}

func getFont(source []byte) *opentype.Font {
	f, err := opentype.Parse(source)
	if err != nil {
		panic(err)
	}

	return f
}

func init() {
	GomduriImage = getImage(gomduriImageFile)
	APlusImage = getImage(aPlusImageFile)
	CatImage = getImage(catImageFile)
	FImage = getImage(fImageFile)
	PlusOneImage = getImage(plusOneImageFile)
	BackgroundImage = getImage(backgroundImageFile)

	GalmuriFont = getFont(galmuriFontFile)
}
