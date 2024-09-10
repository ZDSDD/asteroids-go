package gameobjects

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Player represents an isosceles triangle-shaped player
type Player struct {
	Position Vec2_64 // Position of the player (center of the base of the triangle)
	Base     float64 // Length of the base of the triangle
	Height   float64 // Height of the triangle
	Speed    float64 // Movement speed
}

// Update method for the Player to move it based on keyboard input
func (p *Player) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.Position.Y -= p.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.Position.Y += p.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.Position.X -= p.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.Position.X += p.Speed
	}
}

// Draw method for the Player (draws an isosceles triangle)
func (p *Player) Draw(screen *ebiten.Image) {
	// Calculate the three points of the isosceles triangle
	halfBase := p.Base / 2

	// Points of the triangle
	apex := Vec2_64{X: p.Position.X, Y: p.Position.Y - p.Height}
	baseLeft := Vec2_64{X: p.Position.X - halfBase, Y: p.Position.Y}
	baseRight := Vec2_64{X: p.Position.X + halfBase, Y: p.Position.Y}

	// Draw the triangle using three lines
	ebitenutil.DrawLine(screen, apex.X, apex.Y, baseLeft.X, baseLeft.Y, color.RGBA{0, 255, 0, 255})           // Left side
	ebitenutil.DrawLine(screen, apex.X, apex.Y, baseRight.X, baseRight.Y, color.RGBA{0, 255, 0, 255})         // Right side
	ebitenutil.DrawLine(screen, baseLeft.X, baseLeft.Y, baseRight.X, baseRight.Y, color.RGBA{0, 255, 0, 255}) // Base
}
