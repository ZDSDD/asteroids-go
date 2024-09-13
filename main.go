package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zdsdd/asteroids/internal/constants"
	"github.com/zdsdd/asteroids/internal/gameobjects"
	"github.com/zdsdd/asteroids/internal/managers"
)

type Game struct {
	gameObjects []gameobjects.GameObject
	updatable   []gameobjects.Updatable
}

func (g *Game) Update() error {
	for _, v := range g.gameObjects {
		err := v.Update()
		if err != nil {
			fmt.Printf("There was an error in updating gameobjects\n%v\n%v\n", err, v)
			return err
		}
	}
	for _, v := range g.updatable {
		err := v.Update()
		if err != nil {
			fmt.Printf("There was an error in updateables\n%v\n%v\n", err, v)
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

	onAstroidBulletHit := func(b *gameobjects.Bullet, a *gameobjects.Asteroid) {
		player.RemoveBullet(b)
		a.Brake()
	}
	asteroidManager := managers.NewAsteroidManager()

	onPlayerAndAsteroidCollision := func(p *gameobjects.Player, a *gameobjects.Asteroid) {
		player.Velocity.X = a.Velocity.X * 1.5
		player.Velocity.Y = a.Velocity.Y * 3
	}

	gameManager := managers.NewGameManager(player, asteroidManager, onPlayerAndAsteroidCollision, onAstroidBulletHit)

	return &Game{
		gameObjects: []gameobjects.GameObject{
			player, asteroidManager,
		},
		updatable: []gameobjects.Updatable{gameManager},
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
