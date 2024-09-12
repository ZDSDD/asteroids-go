package gameobjects

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type CircleShape struct {
	X, Y, Radius, StrokeWidth float32
	Color                     color.RGBA
}

func (cs *CircleShape) Draw(dest *ebiten.Image) {
	vector.StrokeCircle(dest, cs.X, cs.Y, cs.Radius, cs.StrokeWidth, cs.Color, false)
}

type TriangleShape struct {
	Position               Vec2 // Position of the player (center of the base of the triangle)
	Base, Height, Rotation float32
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
func (ts *TriangleShape) Draw(dest *ebiten.Image) {
	// Calculate the three points of the isosceles triangle
	halfBase := ts.Base / 2

	// Points of the triangle before rotation
	apex := Vec2{X: ts.Position.X, Y: ts.Position.Y - ts.Height}
	baseLeft := Vec2{X: ts.Position.X - halfBase, Y: ts.Position.Y}
	baseRight := Vec2{X: ts.Position.X + halfBase, Y: ts.Position.Y}

	// Calculate the center of mass (centroid) of the triangle
	// The centroid is 1/3rd of the way up from the base towards the apex
	centroidX := (apex.X + baseLeft.X + baseRight.X) / 3
	centroidY := (apex.Y + baseLeft.Y + baseRight.Y) / 3

	// Rotate the triangle points around the centroid (center of mass)
	apex.X, apex.Y = rotatePoint(apex.X, apex.Y, centroidX, centroidY, ts.Rotation)
	baseLeft.X, baseLeft.Y = rotatePoint(baseLeft.X, baseLeft.Y, centroidX, centroidY, ts.Rotation)
	baseRight.X, baseRight.Y = rotatePoint(baseRight.X, baseRight.Y, centroidX, centroidY, ts.Rotation)

	// Draw the triangle using three lines
	vector.StrokeLine(dest, apex.X, apex.Y, baseLeft.X, baseLeft.Y, 2, color.RGBA{0, 255, 0, 255}, false)           // Left side
	vector.StrokeLine(dest, apex.X, apex.Y, baseRight.X, baseRight.Y, 2, color.RGBA{0, 255, 0, 255}, false)         // Right side
	vector.StrokeLine(dest, baseLeft.X, baseLeft.Y, baseRight.X, baseRight.Y, 2, color.RGBA{0, 255, 0, 255}, false) // Base

}
