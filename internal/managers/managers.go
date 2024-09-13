package managers

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zdsdd/asteroids/internal/constants"
	"github.com/zdsdd/asteroids/internal/gameobjects"
	"github.com/zdsdd/asteroids/internal/sliceutils"
)

type AsteroidManager struct {
	asteroids     []*gameobjects.Asteroid
	lastSpawnTime time.Time
}

func (am *AsteroidManager) GetAsteroids() []*gameobjects.Asteroid {
	return am.asteroids
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
	onKillFunc := func(as *gameobjects.Asteroid) {
		fmt.Printf("Asteroid that was killed had %v radius. What should happen to it?\n", as.Radius)

		// Seed the random number generator
		rand.Seed(time.Now().UnixNano())

		if as.Radius/2 >= constants.ASTEROID_MIN_RADIUS {
			// Generate random angles for splitting
			randomAngle1 := rand.Float64() * math.Pi // random angle between 0 and Pi
			randomAngle2 := rand.Float64() * math.Pi // another random angle between 0 and Pi

			// Create two new asteroids with random rotations
			asteroid1 := gameobjects.NewAsteroid(
				as.OnKill,
				as.OnOutOfScrFunc,
				as.Velocity.Rotate(randomAngle1),
				as.Position,
				as.Radius/2,
			)
			asteroid2 := gameobjects.NewAsteroid(
				as.OnKill,
				as.OnOutOfScrFunc,
				as.Velocity.Rotate(-randomAngle2),
				as.Position,
				as.Radius/2,
			)

			// Add new asteroids to the manager
			am.asteroids = append(am.asteroids, asteroid1)
			am.asteroids = append(am.asteroids, asteroid2)
		}

		// Remove the original asteroid
		am.RemoveAsteroid(as)
	}

	asteroid := gameobjects.NewAsteroidTowardsWindow(onKillFunc, onOutOfScreen)
	am.asteroids = append(am.asteroids, asteroid)
	am.lastSpawnTime = time.Now()
}

func (am *AsteroidManager) RemoveAsteroid(a *gameobjects.Asteroid) {
	am.asteroids, _ = sliceutils.RemoveItem(am.asteroids, func(ast *gameobjects.Asteroid) bool { return ast == a })
}
