package gameobjects

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zdsdd/asteroids/internal/constants"
)

type OnOutOfScreen func(*Asteroid)

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

	// Check if asteroid is within the Kill Zone (beyond the screen + Safe Zone)
	if a.CircleShape.Shape.Position.X < -constants.ASTEROID_MAX_RADIUS*2 ||
		a.CircleShape.Shape.Position.X > constants.SCREEN_WIDTH+constants.ASTEROID_MAX_RADIUS*2 ||
		a.CircleShape.Shape.Position.Y < -constants.ASTEROID_MAX_RADIUS*2 ||
		a.CircleShape.Shape.Position.Y > constants.SCREEN_HEIGHT+constants.ASTEROID_MAX_RADIUS*2 {

		// If asteroid is beyond the Kill Zone, call the onOutOfScreen function
		if a.OnOutOfScrFunc != nil {
			a.OnOutOfScrFunc(a)
		}
	}

	return nil
}

func NewAsteroidTowardsWindow(onOutOfScreen OnOutOfScreen) *Asteroid {
	var position Vec2
	var velocity Vec2

	// Randomize spawn location in Safe Zone
	side := rand.Intn(4) // 0 = left, 1 = right, 2 = top, 3 = bottom

	switch side {
	case 0: // Left side (Safe Zone on the left)
		position = Vec2{
			X: -constants.ASTEROID_MAX_RADIUS,
			Y: rand.Float32() * (constants.SCREEN_HEIGHT + 2*constants.ASTEROID_MAX_RADIUS),
		}
		velocity = Vec2{
			X: rand.Float32() * 2,   // Move right toward the screen
			Y: rand.Float32()*2 - 1, // Random Y
		}
	case 1: // Right side (Safe Zone on the right)
		position = Vec2{
			X: constants.SCREEN_WIDTH + constants.ASTEROID_MAX_RADIUS,
			Y: rand.Float32() * (constants.SCREEN_HEIGHT + 2*constants.ASTEROID_MAX_RADIUS),
		}
		velocity = Vec2{
			X: -rand.Float32() * 2,  // Move left toward the screen
			Y: rand.Float32()*2 - 1, // Random Y
		}
	case 2: // Top side (Safe Zone on the top)
		position = Vec2{
			X: rand.Float32() * (constants.SCREEN_WIDTH + 2*constants.ASTEROID_MAX_RADIUS),
			Y: -constants.ASTEROID_MAX_RADIUS,
		}
		velocity = Vec2{
			X: rand.Float32()*2 - 1, // Random X
			Y: rand.Float32() * 2,   // Move down toward the screen
		}
	case 3: // Bottom side (Safe Zone on the bottom)
		position = Vec2{
			X: rand.Float32() * (constants.SCREEN_WIDTH + 2*constants.ASTEROID_MAX_RADIUS),
			Y: constants.SCREEN_HEIGHT + constants.ASTEROID_MAX_RADIUS,
		}
		velocity = Vec2{
			X: rand.Float32()*2 - 1, // Random X
			Y: -rand.Float32() * 2,  // Move up toward the screen
		}
	}

	// Generate a random radius for the asteroid
	radius := rand.Float32()*(constants.ASTEROID_MAX_RADIUS-constants.ASTEROID_MIN_RADIUS) + constants.ASTEROID_MIN_RADIUS

	// Return the new asteroid with the calculated position and velocity
	return &Asteroid{
		CircleShape: CircleShape{
			Shape: Shape{
				Position:    position,
				StrokeWidth: 1,
				Color:       color.RGBA{255, 255, 255, 255},
			},
			Radius: radius,
		},
		Velocity:       velocity,
		OnOutOfScrFunc: onOutOfScreen,
	}
}
