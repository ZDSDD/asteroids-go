package managers

import (
	"github.com/zdsdd/asteroids/internal/gameobjects"
)

type OnPlayerColl func(asteroid *gameobjects.Asteroid)

type GameManager struct {
	player            *gameobjects.Player
	asteroids         *[]*gameobjects.Asteroid
	OnPlayerCollision OnPlayerColl
}

func (gm *GameManager) Update() error {
	for _, v := range *gm.asteroids {
		distanceBetweenCenters := gameobjects.Distance(gm.player.Collider.Position, v.Collider.Position)
		sumOfRadiuses := gm.player.Collider.Radius + v.Collider.Radius
		if distanceBetweenCenters < sumOfRadiuses {
			if gm.OnPlayerCollision != nil {
				gm.OnPlayerCollision(v)
			}
		}
	}
	return nil
}

func NewGameManager(
	p *gameobjects.Player,
	ast *[]*gameobjects.Asteroid,
	onPlayerCollFunc OnPlayerColl) *GameManager {
	return &GameManager{
		player:            p,
		asteroids:         ast,
		OnPlayerCollision: onPlayerCollFunc,
	}
}
