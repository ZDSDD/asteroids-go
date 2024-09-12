package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zdsdd/asteroids/internal/gamelogic"
	"github.com/zdsdd/asteroids/internal/gameobjects"
)

type Game struct {
	gameObjects []gameobjects.GameObject
}

func (g *Game) Update() error {
	for _, v := range g.gameObjects {
		err := v.Update()
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, v := range g.gameObjects {
		v.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return gamelogic.SCREEN_WIDTH, gamelogic.SCREEN_HEIGHT // Set the game window size
}

func newGame() *Game {
	asteroid := &gameobjects.Asteroid{
		CircleShape: gameobjects.CircleShape{X: 100, Y: 100, Radius: 50, StrokeWidth: 1, Color: color.RGBA{255, 255, 255, 255}},
		Velocity:    gameobjects.Vec2{X: 2, Y: 1}, // Move right and slightly down
	}

	player := gameobjects.NewPlayer(320, 240, 40, 60, 0.04, 0.02, gameobjects.Vec2{X: 0, Y: 0})
	return &Game{
		gameObjects: []gameobjects.GameObject{
			asteroid, player,
		},
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
