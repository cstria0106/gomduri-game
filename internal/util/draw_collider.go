package util

import (
	"github.com/Tarliton/collision2d"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"math"
)

func DrawCollider(image *ebiten.Image, collider *collision2d.Polygon) {
	minX, minY, maxX, maxY := math.MaxFloat64, math.MaxFloat64, 0.0, 0.0
	for _, point := range collider.CalcPoints {
		if point.X < minX {
			minX = point.X
		}

		if point.X > maxX {
			maxX = point.X
		}

		if point.Y < minY {
			minY = point.Y
		}

		if point.Y > maxY {
			maxY = point.Y
		}
	}
	ebitenutil.DrawRect(image, minX, minY, maxX-minX, maxY-minY, HexColor(0xff000099))
}
