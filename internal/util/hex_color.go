package util

import (
	"image/color"
)

func HexColor(v uint32) color.Color {
	return color.RGBA{
		R: uint8((v >> 24) & 0xff),
		G: uint8((v >> 16) & 0xff),
		B: uint8((v >> 8) & 0xff),
		A: uint8(v & 0xff),
	}
}
