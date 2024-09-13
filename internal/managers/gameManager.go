package managers

import (
	"github.com/zdsdd/asteroids/internal/gameobjects"
)

// OnPlayerAsteroidCollisionFunc is a type for player-asteroid collision callback functions
type OnPlayerAsteroidCollisionFunc func(p *gameobjects.Player, a *gameobjects.Asteroid)

// OnBulletAsteroidCollisionFunc is a type for bullet-asteroid collision callback functions
type OnBulletAsteroidCollisionFunc func(b *gameobjects.Bullet, a *gameobjects.Asteroid)

// GameManager manages the game state and interactions
type GameManager struct {
	player            *gameobjects.Player
	asteroidManager   *AsteroidManager
	OnPlayerCollision OnPlayerAsteroidCollisionFunc
	OnBulletCollision OnBulletAsteroidCollisionFunc
}

// Update handles the game logic for each frame
func (gm *GameManager) Update() error {
	gm.checkPlayerCollisions()
	gm.checkBulletCollisions()
	return nil
}

// checkPlayerCollisions checks for collisions between the player and asteroids
func (gm *GameManager) checkPlayerCollisions() {
	for _, asteroid := range gm.asteroidManager.GetAsteroids() {
		if gm.checkCollision(gm.player, asteroid) {
			if gm.OnPlayerCollision != nil {
				gm.OnPlayerCollision(gm.player, asteroid)
			}
		}
	}
}

// checkBulletCollisions checks for collisions between bullets and asteroids
func (gm *GameManager) checkBulletCollisions() {
	for _, bullet := range gm.player.GetBullets() {
		for _, asteroid := range gm.asteroidManager.GetAsteroids() {
			if gm.checkCollision(bullet, asteroid) {
				if gm.OnBulletCollision != nil {
					gm.OnBulletCollision(bullet, asteroid)
				}
			}
		}
	}
}

// checkCollision checks if two objects are colliding
func (gm *GameManager) checkCollision(obj1, obj2 gameobjects.Collidable) bool {
	distance := gameobjects.Distance(obj1.GetPosition(), obj2.GetPosition())
	sumOfRadii := obj1.GetRadius() + obj2.GetRadius()
	return distance < sumOfRadii
}

// NewGameManager creates a new GameManager instance
func NewGameManager(
	p *gameobjects.Player,
	am *AsteroidManager,
	onPlayerCollFunc OnPlayerAsteroidCollisionFunc,
	onBulletCollFunc OnBulletAsteroidCollisionFunc) *GameManager {
	return &GameManager{
		player:            p,
		asteroidManager:   am,
		OnPlayerCollision: onPlayerCollFunc,
		OnBulletCollision: onBulletCollFunc,
	}
}
