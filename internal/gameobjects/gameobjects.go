package gameobjects

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Drawable interface {
	Draw(*ebiten.Image)
}

type Updateable interface {
	Update()
}

type GameObject interface {
	Drawable
	Updateable
}

type Vec2_64 struct {
	X, Y float64
}

type Vec2 struct {
	X, Y float32
}
