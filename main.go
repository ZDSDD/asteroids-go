package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/zdsdd/asteroids/internal/gameobjects"
)

type Game struct {
	asteroids []*gameobjects.Asteroid
	player    *gameobjects.Player
}

func (g *Game) Update() error {
	// Check if the left mouse button was pressed
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		fmt.Printf("Mouse clicked at: X: %d, Y: %d\n", x, y)
		g.asteroids[0].CircleShape.X, g.asteroids[0].Y = float32(x), float32(y)
	}
	// deltaTime := 1.0 / ebiten.ActualTPS()
	for _, v := range g.asteroids {
		v.Update()
	}

	g.player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw a red rectangle
	g.asteroids[0].Draw(screen)
	g.player.Draw(screen)
}

// Draw a triangle using three points
func drawTriangle(screen *ebiten.Image, x1, y1, x2, y2, x3, y3 int, clr color.Color) {
	vector.StrokeLine(screen, float32(x1), float32(y1), float32(x2), float32(y2), 4, clr, false)
	vector.StrokeLine(screen, float32(x2), float32(y2), float32(x3), float32(y3), 4, clr, false)
	vector.StrokeLine(screen, float32(x3), float32(y3), float32(x1), float32(y1), 4, clr, false)
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480 // Set the game window size
}

func newGame() *Game {
	asteroid := &gameobjects.Asteroid{
		CircleShape: gameobjects.CircleShape{X: 100, Y: 100, Radius: 50},
		Velocity:    gameobjects.Vec2{X: 2, Y: 1}, // Move right and slightly down
	}

	asteroids := []*gameobjects.Asteroid{
		asteroid,
	}

	player := &gameobjects.Player{
		TriangleShape: gameobjects.TriangleShape{
			Position: gameobjects.Vec2{X: 320, Y: 240}, // Start position at the center
			Base:     40,
			Height:   60,
		},
		Speed: 2, // Player movement speed
	}
	return &Game{
		asteroids: asteroids,
		player:    player,
	}
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Mouse Click Example")
	ebiten.SetVsyncEnabled(true)

	game := newGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
