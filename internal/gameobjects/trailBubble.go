package gameobjects

import (
	"image/color"
	"time"
)

const (
	maxTrailBubbles = 100
)

type OnDeathFunc func(tb *trailBubble) error

type trailBubble struct {
	shape     *CircleShape
	lifeTime  time.Duration
	spawnTime time.Time
	Velocity  Vec2
	OnDeath   OnDeathFunc
}

func NewTrailBubble(x, y float32, onDeathFunc OnDeathFunc) *trailBubble {
	return &trailBubble{
		shape: &CircleShape{
			X:           x, // Initial position same as player
			Y:           y,
			Radius:      5, // Example radius
			StrokeWidth: 1,
			Color:       color.RGBA{255, 255, 255, 200},
		},
		spawnTime: time.Now(),
		lifeTime:  3 * time.Second,
		OnDeath:   onDeathFunc,
	}
}

func (tb *trailBubble) Update() error {
	if tb.OnDeath == nil {
		return nil
	}
	timeSince := time.Since(tb.spawnTime)
	// fmt.Printf("timeSince: %v, lifeTime:%v\n", timeSince.Seconds(), float64(tb.lifeTime))
	if timeSince >= tb.lifeTime {
		return tb.OnDeath(tb)
	}
	return nil
}
