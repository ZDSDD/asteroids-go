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
		lifeTime:  1 * time.Second,
		OnDeath:   onDeathFunc,
	}
}

func (tb *trailBubble) Update() error {
	timeSince := time.Since(tb.spawnTime)

	// Calculate the remaining lifetime as a fraction
	remainingLifeFraction := 1 - float64(timeSince)/float64(tb.lifeTime)

	// Ensure the fraction is between 0 and 1
	if remainingLifeFraction < 0 {
		remainingLifeFraction = 0
	} else if remainingLifeFraction > 1 {
		remainingLifeFraction = 1
	}

	// Calculate alpha value (255 when new, approaching 0 as it ages)
	alpha := uint8(255 * remainingLifeFraction)

	// Update the color with the new alpha value
	tb.shape.Color = color.RGBA{alpha, alpha, alpha, 255}

	// Update position
	tb.shape.X += tb.Velocity.X
	tb.shape.Y += tb.Velocity.Y

	// Check if the bubble should be removed
	if timeSince >= tb.lifeTime {
		if tb.OnDeath != nil {
			return tb.OnDeath(tb)
		}
		return nil
	}

	return nil
}
