package gameobjects

import "github.com/hajimehoshi/ebiten/v2"

type Asteroid struct {
	CircleShape
	Velocity Vec2
}

func (ast *Asteroid) Draw(dest *ebiten.Image) {
	ast.CircleShape.Draw(dest)
}

func (a *Asteroid) Update() error {
	a.CircleShape.X += a.Velocity.X
	a.CircleShape.Y += a.Velocity.Y
	return nil
}
