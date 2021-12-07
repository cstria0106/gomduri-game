package game

import (
	"github.com/Tarliton/collision2d"
	"github.com/cstria0106/gomduri/internal/game/engine"
	"github.com/cstria0106/gomduri/internal/util"
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
)

const FallingItemMinSpeed = 1
const FallingItemMaxSpeed = 4

type FallingItem struct {
	engine.BaseObject

	speed float64
	image *ebiten.Image
	x, y  float64

	collider       *collision2d.Polygon
	colliderAdjust collision2d.Vector

	debug bool
}

func NewFallingItem(image *ebiten.Image, box collision2d.Box) *FallingItem {
	imageW, imageH := image.Size()

	x := rand.Float64() * float64(GameWidth-imageW)
	y := float64(-imageH)

	colliderAdjust := box.Pos.Clone()
	shape := box.ToPolygon().SetOffset(colliderAdjust.Add(collision2d.Vector{X: x, Y: y}))

	return &FallingItem{
		image:          image,
		speed:          FallingItemMinSpeed + rand.Float64()*(FallingItemMaxSpeed-FallingItemMinSpeed),
		x:              x,
		y:              y,
		collider:       &shape,
		colliderAdjust: colliderAdjust,
		debug:          false,
	}
}

func (f *FallingItem) Update(gameInterface interface{}) (earned bool, err error) {
	game := gameInterface.(*Game)

	f.y += f.speed

	f.collider.SetOffset(
		collision2d.Vector{
			X: f.x,
			Y: f.y,
		}.Add(f.colliderAdjust),
	)

	if colliding, _ := collision2d.TestPolygonPolygon(*f.collider, *game.gomduri.collider); colliding {
		return true, nil
	}

	if f.y > GameHeight {
		imageW, imageH := f.image.Size()

		x := rand.Float64() * float64(GameWidth-imageW)
		y := float64(-imageH)
		f.x = x
		f.y = y
	}

	return false, nil
}

func (f *FallingItem) Priority() int {
	return -1
}

func (f *FallingItem) DrawOn(image *ebiten.Image) {
	o := &ebiten.DrawImageOptions{}
	o.GeoM.Translate(f.x, f.y)
	image.DrawImage(f.image, o)

	if f.debug {
		util.DrawCollider(image, f.collider)
	}
}

func (f *FallingItem) Layer() int {
	return -1
}

func (f *FallingItem) String() string {
	return "FallingItem"
}
