package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zdsdd/asteroids/internal/gamelogic"
	"github.com/zdsdd/asteroids/internal/gameobjects"
)

type Game struct {
	asteroids []*gameobjects.Asteroid
	player    *gameobjects.Player
}

func (g *Game) Update() error {
	for _, v := range g.asteroids {
		v.Update()
	}
	g.player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, v := range g.asteroids {
		v.Draw(screen)
	}
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return gamelogic.SCREEN_WIDTH, gamelogic.SCREEN_HEIGHT // Set the game window size
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
			Rotation: 180,
		},
		Speed: 2, // Player movement speed
	}
	return &Game{
		asteroids: asteroids,
		player:    player,
	}
}

func main() {
	ebiten.SetWindowSize(gamelogic.SCREEN_WIDTH, gamelogic.SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Asteroid")
	ebiten.SetVsyncEnabled(true)

	game := newGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
