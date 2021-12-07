package game

import (
	"fmt"
	"github.com/Tarliton/collision2d"
	"github.com/cstria0106/gomduri/internal/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"math"
)

const GomduriSpeed = 2.5
const GomduriAccelerationSpeed = 0.4
const GomduriSpeedMultiplyOnJump = 0.5
const GomduriWidth = 32
const GomduriHeight = 64
const GomduriGroundY = GameHeight - 24
const GomduriInitialJumpForce = 2.5
const GomduriJumpAccelerationForce = 0.3
const GomduriGravity = 0.2
const GomduriMaxJumpFrame = 30

type Gomduri struct {
	sourceImage *ebiten.Image

	// position
	speedX    float64
	speedY    float64
	x         float64
	y         float64
	lookingAt int

	// movement
	moveFrame int
	jumpFrame int
	longJump  bool

	// sprite/animation
	sprites     []*ebiten.Image
	spriteIndex int

	// collision
	collider       *collision2d.Polygon
	colliderAdjust collision2d.Vector

	// debug
	debug bool
}

func NewGomduri(sourceImage *ebiten.Image) *Gomduri {
	var sprites []*ebiten.Image

	for i := 0; i < 12; i++ {
		sprites = append(
			sprites,
			sourceImage.SubImage(image.Rect(
				GomduriWidth*i,
				0,
				GomduriWidth*(i+1),
				GomduriHeight,
			)).(*ebiten.Image),
		)
	}

	var x float64 = GameWidth/2 - GomduriWidth/2
	var y float64 = GomduriGroundY

	adjust := collision2d.Vector{X: -GomduriWidth / 2, Y: -48}
	collider := collision2d.NewBox(collision2d.Vector{X: 0, Y: 0}, GomduriWidth, 32).
		ToPolygon().
		SetOffset(collision2d.Vector{X: x, Y: y}.Add(adjust))

	gomduri := &Gomduri{
		sourceImage:    sourceImage,
		speedX:         0,
		speedY:         0,
		x:              x,
		y:              y,
		lookingAt:      1,
		moveFrame:      0,
		jumpFrame:      0,
		longJump:       false,
		sprites:        sprites,
		spriteIndex:    0,
		collider:       &collider,
		colliderAdjust: adjust,
		debug:          false,
	}

	return gomduri
}

func (g *Gomduri) Update(gameInterface interface{}) (bool, error) {
	game := gameInterface.(*Game)

	// get direction
	direction := 0

	if game.input.Left {
		direction -= 1
	}

	if game.input.Right {
		direction += 1
	}

	// jump
	if g.y == GomduriGroundY && game.input.JumpPressed {
		g.speedY = -GomduriInitialJumpForce
		g.longJump = true
	}

	if g.y+g.speedY < GomduriGroundY && game.input.Jump && g.longJump {
		g.jumpFrame += 1
		if g.jumpFrame < GomduriMaxJumpFrame {
			g.speedY += -(GomduriJumpAccelerationForce +
				(GomduriGravity / GomduriMaxJumpFrame)) *
				(float64(GomduriMaxJumpFrame-g.jumpFrame) / GomduriMaxJumpFrame)
		}
	} else {
		g.longJump = false
	}

	// horizontal move
	multiply := 1.0
	if g.speedY != 0 {
		multiply = GomduriSpeedMultiplyOnJump
	}

	if direction != 0 {
		g.speedX += float64(direction) * GomduriAccelerationSpeed * multiply
	} else {
		if g.speedX > 0 {
			g.speedX -= GomduriAccelerationSpeed
			if g.speedX < 0 {
				g.speedX = 0
			}
		} else if g.speedX < 0 {
			g.speedX += GomduriAccelerationSpeed
			if g.speedX > 0 {
				g.speedX = 0
			}
		}
	}

	if g.x+g.speedX < GomduriWidth/2 {
		g.x = GomduriWidth / 2
		g.speedX = 0
		direction = 0
	}

	if g.x+g.speedX > GameWidth-GomduriWidth/2 {
		g.x = GameWidth - GomduriWidth/2
		g.speedX = 0
		direction = 0
	}

	if g.speedX > GomduriSpeed {
		g.speedX = GomduriSpeed
	} else if g.speedX < -GomduriSpeed {
		g.speedX = -GomduriSpeed
	}

	// actual move
	g.x += g.speedX
	g.y += g.speedY

	// clamp y
	if g.y >= GomduriGroundY {
		g.y = GomduriGroundY
		g.speedY = 0
		g.jumpFrame = 0
	} else {
		g.speedY += GomduriGravity
	}

	// animation
	if g.speedY < 0 {
		g.moveFrame += 1

		g.spriteIndex = 8 + (g.moveFrame/4)%2
	} else if g.speedY > 0 {
		g.moveFrame += 1

		g.spriteIndex = 10 + (g.moveFrame/4)%2
	} else if g.speedX == 0 {
		g.spriteIndex = 1
		g.moveFrame = 0
	} else {
		if direction != 0 {
			g.lookingAt = direction
		}

		g.moveFrame += 1

		g.spriteIndex = 3 + (g.moveFrame/5)%4
	}

	// adjust collider
	g.collider.SetOffset(
		collision2d.Vector{
			X: math.Floor(g.x),
			Y: math.Floor(g.y),
		}.Add(g.colliderAdjust),
	)

	return false, nil
}

func (g *Gomduri) Priority() int {
	return -2
}

func (g *Gomduri) DrawOn(image *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}

	options.GeoM.Scale(float64(-g.lookingAt), 1)
	options.GeoM.Translate(
		math.Floor(g.x+GomduriWidth/2*float64(g.lookingAt)),
		math.Floor(g.y-GomduriHeight),
	)

	image.DrawImage(
		g.sprites[g.spriteIndex], options,
	)

	if g.debug {
		ebitenutil.DebugPrint(
			image,
			fmt.Sprintf(
				"x: %f\n"+
					"y: %f\n"+
					"speedX: %f\n"+
					"speedY: %f\n"+
					"lookingAt: %d\n"+
					"longJump: %v",
				g.x,
				g.y,
				g.speedX,
				g.speedY,
				g.lookingAt,
				g.longJump,
			),
		)

		util.DrawCollider(image, g.collider)
	}

}

func (g *Gomduri) Layer() int {
	return 0
}

func (g *Gomduri) String() string {
	return "Gomduri"
}
