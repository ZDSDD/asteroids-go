package gameobjects

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// Player represents an isosceles triangle-shaped player
type Player struct {
	TriangleShape
	Speed float32 // Movement speed
}

// Update method for the Player to move and rotate based on input
func (p *Player) Update() error {
	// The forward vector should point towards the player's head (the apex of the triangle)
	// Calculate forward direction using the rotation angle
	// The forward vector is aligned with the player's "head" point (apex of the triangle)
	forwardX := float32(math.Sin(float64(p.Rotation)))  // Use sine for the x-direction
	forwardY := -float32(math.Cos(float64(p.Rotation))) // Use cosine for the y-direction, negative because y-axis is inverted

	// Movement logic (WASD or arrow keys for movement)
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		// Move forward towards the player's head (apex) along the forward vector
		p.Position.X += forwardX * p.Speed
		p.Position.Y += forwardY * p.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		// Move backward (reverse) along the forward vector
		p.Position.X -= forwardX * p.Speed
		p.Position.Y -= forwardY * p.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.Rotation -= 0.05 // Rotate left (counterclockwise)
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.Rotation += 0.05 // Rotate right (clockwise)
	}
	return nil
}

// Draw method for the Player (draws an isosceles triangle)
func (p *Player) Draw(screen *ebiten.Image) {
	p.TriangleShape.Draw(screen)
}
