package managers

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zdsdd/asteroids/internal/constants"
	"github.com/zdsdd/asteroids/internal/gameobjects"
)

type AsteroidManager struct {
	asteroids     []*gameobjects.Asteroid
	lastSpawnTime time.Time
}

func (am *AsteroidManager) GetAsteroids() *[]*gameobjects.Asteroid {
	return &am.asteroids
}

func NewAsteroidManager() *AsteroidManager {
	asteroids := make([]*gameobjects.Asteroid, 0)
	return &AsteroidManager{
		asteroids:     asteroids,
		lastSpawnTime: time.Now(),
	}
}

func (am *AsteroidManager) Update() error {
	for _, v := range am.asteroids {
		if err := v.Update(); err != nil {
			return err
		}
	}
	timeSinceLastSpawn := time.Since(am.lastSpawnTime)
	if timeSinceLastSpawn.Seconds() >= constants.ASTEROID_SPAWN_RATE {
		am.SpawnAsteroid()
	}
	return nil
}

func (am *AsteroidManager) Draw(dest *ebiten.Image) {
	for _, v := range am.asteroids {
		v.Draw(dest)
	}
}

func (am *AsteroidManager) SpawnAsteroid() {
	onOutOfScreen := func(as *gameobjects.Asteroid) {
		for i, v := range am.asteroids {
			if v == as {
				am.asteroids = append(am.asteroids[:i], am.asteroids[i+1:]...)
			}
		}
	}

	asteroid := gameobjects.NewAsteroidTowardsWindow(onOutOfScreen)
	am.asteroids = append(am.asteroids, asteroid)
	am.lastSpawnTime = time.Now()
}
