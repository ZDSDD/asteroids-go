package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zdsdd/asteroids/internal/constants"
	"github.com/zdsdd/asteroids/internal/gameobjects"
	"github.com/zdsdd/asteroids/internal/managers"
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
	return constants.SCREEN_WIDTH, constants.SCREEN_HEIGHT // Set the game window size
}

func newGame() *Game {
	player := gameobjects.NewPlayer(320, 240, 40, 60, 0.04, 0.02, gameobjects.Vec2{X: 0, Y: 0})
	return &Game{
		gameObjects: []gameobjects.GameObject{
			player, managers.NewAsteroidManager(),
		},
	}
}

func main() {
	ebiten.SetWindowSize(constants.SCREEN_WIDTH, constants.SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Asteroid")
	ebiten.SetVsyncEnabled(true)

	game := newGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
