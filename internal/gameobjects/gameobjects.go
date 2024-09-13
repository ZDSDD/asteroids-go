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

type Collidable interface {
	GetPosition() Vec2
	GetRadius() float32
}
