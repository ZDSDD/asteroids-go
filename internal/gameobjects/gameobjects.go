package gameobjects

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Drawable interface {
	Draw(*ebiten.Image)
}

type Updatable interface {
	Update() error
}

type GameObject interface {
	Drawable
	Updatable
}

type Vec2_64 struct {
	X, Y float64
}

type Vec2 struct {
	X, Y float32
}

// Function to calculate distance between two Vec2 points
func Distance(p1, p2 Vec2) float32 {
	return float32(math.Sqrt(float64((p2.X-p1.X)*(p2.X-p1.X) + (p2.Y-p1.Y)*(p2.Y-p1.Y))))
}

// Rotate rotates the vector by a given angle in radians
func (v *Vec2) Rotate(angle float64) Vec2 {
	// Calculate the new x and y using the rotation matrix
	cosAngle := float32(math.Cos(angle))
	sinAngle := float32(math.Sin(angle))
	xNew := float32(v.X*cosAngle - v.Y*sinAngle)
	yNew := float32(v.X*sinAngle + v.Y*cosAngle)

	return Vec2{X: xNew, Y: yNew}
}

type Collidable interface {
	GetPosition() Vec2
	GetRadius() float32
}
