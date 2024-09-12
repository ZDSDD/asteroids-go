package gameobjects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zdsdd/asteroids/internal/gamelogic"
)

type OnOutOfScreen func(*Asteroid) error

type Asteroid struct {
	CircleShape
	Velocity       Vec2
	OnOutOfScrFunc OnOutOfScreen
}

func (ast *Asteroid) Draw(dest *ebiten.Image) {
	ast.CircleShape.Draw(dest)
}

func (a *Asteroid) Update() error {
	a.CircleShape.Shape.Position.X += a.Velocity.X
	a.CircleShape.Shape.Position.Y += a.Velocity.Y
	if a.CircleShape.Shape.Position.X < 0 || a.CircleShape.Shape.Position.X > gamelogic.SCREEN_WIDTH {
		a.OnOutOfScrFunc(a)
	}
	return nil
}
