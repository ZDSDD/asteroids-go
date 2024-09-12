package gameobjects

import (
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zdsdd/asteroids/internal/gamelogic"
)

const (
	maxTrailBubbles = 64
)

// Player represents an isosceles triangle-shaped player
type Player struct {
	shape           TriangleShape
	Velocity        Vec2 // Movement speed
	AcceleratePower float32
	DeceleratePower float32
	trailBubbles    []*trailBubble
}

func (p *Player) spawnTrail() {
	onDeath := func(tb *trailBubble) error {
		bubbles := make([]*trailBubble, 0, maxTrailBubbles)
		// Remove bubble from the list here
		for i, bubble := range p.trailBubbles {
			if bubble != tb {
				bubbles = append(bubbles, p.trailBubbles[i])
			}
		}
		p.trailBubbles = bubbles
		return nil
	}
	forwardX := float32(math.Sin(float64(p.shape.Rotation)))
	forwardY := -float32(math.Cos(float64(p.shape.Rotation)))
	vel := Vec2{
		X: -forwardX * float32(rand.Float32()*3),
		Y: -forwardY * float32(rand.Float32()*3),
	}
	bubble := NewTrailBubble(p.shape.Position.X, p.shape.Position.Y, vel, onDeath)
	if len(p.trailBubbles) == maxTrailBubbles {
		// Remove the oldest bubble and add the new one
		p.trailBubbles = append(p.trailBubbles[1:], bubble)
	} else {
		p.trailBubbles = append(p.trailBubbles, bubble)
	}
}

// NewPlayer creates and initializes a new Player object.
func NewPlayer(x, y, base, height, acceleratePower, deceleratePower float32, velocity Vec2) *Player {
	player := &Player{
		shape: TriangleShape{
			Position: Vec2{X: x, Y: y},
			Base:     base,
			Height:   height,
		},
		Velocity:        Vec2{X: 0, Y: 0},
		AcceleratePower: acceleratePower,
		DeceleratePower: deceleratePower,
		trailBubbles:    make([]*trailBubble, 0, maxTrailBubbles), // Initialize empty slice for trailBubbles
	}
	// Define the death event handler
	return player
}

// Update method for the Player to move and rotate based on input
func (p *Player) Update() error {
	p.handleMovement()
	for _, v := range p.trailBubbles {
		if err := v.Update(); err != nil {
			return err
		}
	}
	return nil
}

func (p *Player) handleMovement() {
	forwardX := float32(math.Sin(float64(p.shape.Rotation)))
	forwardY := -float32(math.Cos(float64(p.shape.Rotation)))

	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {

		p.Velocity.X += forwardX * p.AcceleratePower
		p.Velocity.Y += forwardY * p.AcceleratePower
		p.spawnTrail()
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {

		// Apply braking effect: gradually reduce velocity without reversing it
		p.Velocity.X *= 0.95 // Scale down the velocity
		p.Velocity.Y *= 0.95

	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {

		p.shape.Rotation -= 0.05
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {

		p.shape.Rotation += 0.05
	}

	p.shape.Position.X += p.Velocity.X
	p.shape.Position.Y += p.Velocity.Y

	bounceBack(&p.shape.Position, &p.Velocity, gamelogic.SCREEN_WIDTH, gamelogic.SCREEN_HEIGHT)
}

// Draw method for the Player (draws an isosceles triangle)
func (p *Player) Draw(screen *ebiten.Image) {
	p.shape.Draw(screen)
	for _, v := range p.trailBubbles {
		v.shape.Draw(screen)
	}
}

// Function to handle bouncing at the screen edges
func bounceBack(position *Vec2, velocity *Vec2, screenWidth, screenHeight float32) {
	if position.X <= 0 || position.X >= screenWidth {
		velocity.X = -velocity.X
		position.X = 0
	}

	if position.Y <= 0 || position.Y >= screenHeight {
		velocity.Y = -velocity.Y
		position.Y = 0
	}
}
