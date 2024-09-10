package gameobjects

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Drawable interface {
	Draw(*ebiten.Image)
}

type Updatable interface {
	Update()
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
