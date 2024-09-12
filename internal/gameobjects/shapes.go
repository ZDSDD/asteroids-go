package gameobjects

import (
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Shape struct {
	Position    Vec2
	StrokeWidth float32
	Color       color.RGBA
}

type CircleShape struct {
	Shape
	Radius float32
}

// TriangleShape represents an isosceles triangle shape
type TriangleShape struct {
	Shape
	Base     float32 // The length of the base of the triangle
	Height   float32 // The height of the triangle from base to apex
	Rotation float32 // The rotation of the triangle in radians
	Filled   bool    // Whether the triangle should be filled when drawn
}

// NewTriangleShape creates a new TriangleShape with the given parameters
func NewTriangleShape(x, y, base, height, rotation float32, color color.RGBA, filled bool) *TriangleShape {
	return &TriangleShape{
		Shape: Shape{
			Position:    Vec2{X: x, Y: y},
			StrokeWidth: 1, // Default stroke width
			Color:       color,
		},
		Base:     base,
		Height:   height,
		Rotation: rotation,
		Filled:   filled,
	}
}

// Scale scales the triangle by the given factor
func (ts *TriangleShape) Scale(factor float32) {
	ts.Base *= factor
	ts.Height *= factor
}

// ContainsPoint checks if the given point is inside the triangle
func (ts *TriangleShape) ContainsPoint(point Vec2) bool {
	// Implementation of point-in-triangle check
	// This is a placeholder and needs to be implemented
	return false
}

func (cs *CircleShape) Draw(dest *ebiten.Image) {
	vector.StrokeCircle(dest, cs.Position.X, cs.Position.Y, cs.Radius, cs.StrokeWidth, cs.Color, false)
}

var emptySubImage = ebiten.NewImage(3, 3).SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)

func (ts *TriangleShape) Draw(dest *ebiten.Image) {
	// Calculate the three points of the isosceles triangle
	halfBase := ts.Base / 2

	// Points of the triangle before rotation
	apex := Vec2{X: 0, Y: -ts.Height / 2}
	baseLeft := Vec2{X: -halfBase, Y: ts.Height / 2}
	baseRight := Vec2{X: halfBase, Y: ts.Height / 2}

	// Rotate the points
	sin, cos := float32(math.Sin(float64(ts.Rotation))), float32(math.Cos(float64(ts.Rotation)))
	rotatePoint := func(p Vec2) Vec2 {
		return Vec2{
			X: p.X*cos - p.Y*sin + ts.Position.X,
			Y: p.X*sin + p.Y*cos + ts.Position.Y,
		}
	}

	apexRotated := rotatePoint(apex)
	baseLeftRotated := rotatePoint(baseLeft)
	baseRightRotated := rotatePoint(baseRight)

	// Draw filled triangle if Filled is true
	if ts.Filled {
		var path vector.Path
		path.MoveTo(apexRotated.X, apexRotated.Y)
		path.LineTo(baseLeftRotated.X, baseLeftRotated.Y)
		path.LineTo(baseRightRotated.X, baseRightRotated.Y)
		path.Close()

		vs, is := path.AppendVerticesAndIndicesForFilling(nil, nil)
		op := &ebiten.DrawTrianglesOptions{
			FillRule: ebiten.EvenOdd,
		}
		dest.DrawTriangles(vs, is, emptySubImage, op)
	}

	// Draw the outline
	vector.StrokeLine(dest, apexRotated.X, apexRotated.Y, baseLeftRotated.X, baseLeftRotated.Y, ts.StrokeWidth, ts.Color, false)
	vector.StrokeLine(dest, baseLeftRotated.X, baseLeftRotated.Y, baseRightRotated.X, baseRightRotated.Y, ts.StrokeWidth, ts.Color, false)
	vector.StrokeLine(dest, baseRightRotated.X, baseRightRotated.Y, apexRotated.X, apexRotated.Y, ts.StrokeWidth, ts.Color, false)
}
