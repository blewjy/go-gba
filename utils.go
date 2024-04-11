package main

import (
	"image/color"
)

func colorToBytes(c color.RGBA) []byte {
	return []byte{c.R, c.G, c.B, c.A}
}
