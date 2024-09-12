package gameobjects

import (
	"image/color"
	"math/rand"
	"time"
)

type OnDeathFunc func(tb *trailBubble) error

type trailBubble struct {
	shape     *CircleShape
	lifeTime  time.Duration
	spawnTime time.Time
	Velocity  Vec2
	OnDeath   OnDeathFunc
}

func NewTrailBubble(x, y float32, vel Vec2, onDeathFunc OnDeathFunc) *trailBubble {
	return &trailBubble{
		shape: &CircleShape{
			X:           x, // Initial position same as player
			Y:           y,
			Radius:      5 + float32(rand.Int()%3), // Example radius
			StrokeWidth: 1,
			Color:       color.RGBA{255, 255, 255, 200},
		},
		Velocity:  vel,
		spawnTime: time.Now(),
		lifeTime:  3 * time.Second,
		OnDeath:   onDeathFunc,
	}
}

func (tb *trailBubble) Update() error {
	tb.shape.X += tb.Velocity.X
	tb.shape.Y += tb.Velocity.Y
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
