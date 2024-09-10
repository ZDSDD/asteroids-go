package gameobjects

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type CircleShape struct {
	X, Y, Radius float32
}

func (cs *CircleShape) Draw(dest *ebiten.Image) {
	vector.StrokeCircle(dest, cs.X, cs.Y, cs.Radius, 10, color.RGBA{255, 0, 0, 255}, false)
}

type TriangleShape struct {
	Position Vec2    // Position of the player (center of the base of the triangle)
	Base     float32 // Length of the base of the triangle
	Height   float32 // Height of the triangle
}
