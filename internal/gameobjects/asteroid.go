package gameobjects

import "github.com/hajimehoshi/ebiten/v2"

type Asteroid struct {
	CircleShape
	Velocity Vec2_32
}

func (ast *Asteroid) Draw(dest *ebiten.Image) {
	ast.CircleShape.Draw(dest)
}

func (a *Asteroid) Update() {
	a.CircleShape.X += a.Velocity.X
	a.CircleShape.Y += a.Velocity.Y
}
