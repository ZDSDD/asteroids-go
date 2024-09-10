package gameobjects

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Player represents an isosceles triangle-shaped player
type Player struct {
	TriangleShape
	Speed    float32 // Movement speed
	Rotation float32
}

// Helper function to rotate a point (x, y) around the origin (cx, cy)
func rotatePoint(x, y, cx, cy, angle float32) (float32, float32) {
	sin, cos := float32(math.Sin(float64(angle))), float32(math.Cos(float64(angle)))

	x -= cx
	y -= cy

	xNew := x*cos - y*sin
	yNew := x*sin + y*cos

	xNew += cx
	yNew += cy

	return xNew, yNew
}

// Update method for the Player to move and rotate based on input
// Update method for the Player to move and rotate based on input
func (p *Player) Update() {
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
}

// Draw method for the Player (draws an isosceles triangle)
func (p *Player) Draw(screen *ebiten.Image) {
	// Calculate the three points of the isosceles triangle
	halfBase := p.Base / 2

	// Points of the triangle before rotation
	apex := Vec2{X: p.Position.X, Y: p.Position.Y - p.Height}
	baseLeft := Vec2{X: p.Position.X - halfBase, Y: p.Position.Y}
	baseRight := Vec2{X: p.Position.X + halfBase, Y: p.Position.Y}

	// Calculate the center of mass (centroid) of the triangle
	// The centroid is 1/3rd of the way up from the base towards the apex
	centroidX := (apex.X + baseLeft.X + baseRight.X) / 3
	centroidY := (apex.Y + baseLeft.Y + baseRight.Y) / 3

	// Rotate the triangle points around the centroid (center of mass)
	apex.X, apex.Y = rotatePoint(apex.X, apex.Y, centroidX, centroidY, p.Rotation)
	baseLeft.X, baseLeft.Y = rotatePoint(baseLeft.X, baseLeft.Y, centroidX, centroidY, p.Rotation)
	baseRight.X, baseRight.Y = rotatePoint(baseRight.X, baseRight.Y, centroidX, centroidY, p.Rotation)

	// Draw the triangle using three lines
	vector.StrokeLine(screen, apex.X, apex.Y, baseLeft.X, baseLeft.Y, 2, color.RGBA{0, 255, 0, 255}, false)           // Left side
	vector.StrokeLine(screen, apex.X, apex.Y, baseRight.X, baseRight.Y, 2, color.RGBA{0, 255, 0, 255}, false)         // Right side
	vector.StrokeLine(screen, baseLeft.X, baseLeft.Y, baseRight.X, baseRight.Y, 2, color.RGBA{0, 255, 0, 255}, false) // Base
}
