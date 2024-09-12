package gameobjects

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zdsdd/asteroids/internal/constants"
)

const (
	maxTrailBubbles = 64
)

// Player represents an isosceles triangle-shaped player
type Player struct {
	AcceleratePower float32
	DeceleratePower float32
	shape           TriangleShape
	Velocity        Vec2 // Movement speed
	trailBubbles    []*trailBubble
	Collider        CircleShape
	Bullets         []*Bullet
	LastTimeShoot   time.Time
}

func (p *Player) spawnTrail() {
	onDeath := func(tb *trailBubble) error {
		bubbles := make([]*trailBubble, 0, maxTrailBubbles)
		for i, bubble := range p.trailBubbles {
			if bubble != tb {
				bubbles = append(bubbles, p.trailBubbles[i])
			}
		}
		p.trailBubbles = bubbles
		return nil
	}
	// Set up a random generator, if not done already

	// Generate a random value between -0.5 and 0.5
	randomOffset := (rand.Float32() - 0.5) * p.shape.Base

	// Calculate the base center of the triangle (unchanged)
	baseCenter := Vec2{
		X: p.shape.Position.X - float32(math.Sin(float64(p.shape.Rotation)))*(p.shape.Height/2),
		Y: p.shape.Position.Y + float32(math.Cos(float64(p.shape.Rotation)))*(p.shape.Height/2),
	}

	// Calculate the random point along the base using trigonometric rotation
	spawnPoint := Vec2{
		X: baseCenter.X + float32(math.Cos(float64(p.shape.Rotation)))*randomOffset,
		Y: baseCenter.Y + float32(math.Sin(float64(p.shape.Rotation)))*randomOffset,
	}

	// Calculate velocity perpendicular to the base
	perpX := -float32(math.Sin(float64(p.shape.Rotation)) * 3)
	perpY := float32(math.Cos(float64(p.shape.Rotation)) * 3)

	vel := Vec2{
		X: perpX,
		Y: perpY,
	}

	bubble := NewTrailBubble(spawnPoint.X, spawnPoint.Y, vel, onDeath)

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
			Shape: Shape{
				Position: Vec2{
					X: x,
					Y: y,
				},
				StrokeWidth: 1,
				Color:       color.RGBA{255, 255, 255, 255},
			},
			Base:   base,
			Height: height,
		},
		Velocity:        Vec2{X: 0, Y: 0},
		AcceleratePower: acceleratePower,
		DeceleratePower: deceleratePower,
		trailBubbles:    make([]*trailBubble, 0, maxTrailBubbles), // Initialize empty slice for trailBubbles
		Collider: CircleShape{
			Shape: Shape{
				Position: Vec2{
					X: x,
					Y: y,
				},
				StrokeWidth: 2,
				Color:       color.RGBA{0, 255, 255, 255},
			},
			Radius: height / 2,
		},
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

	for _, v := range p.Bullets {
		if err := v.Update(); err != nil {
			return err
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return fmt.Errorf("escape key pressed")
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if time.Since(p.LastTimeShoot).Seconds() >= constants.SHOOTING_COOLDOWN {
			p.Bullets = append(p.Bullets, p.spawnBullet())
		}
	}

	return nil
}

func (p *Player) getForwardVector() Vec2 {

	forwardX := float32(math.Sin(float64(p.shape.Rotation)))
	forwardY := -float32(math.Cos(float64(p.shape.Rotation)))
	return Vec2{X: forwardX, Y: forwardY}
}

func (p *Player) handleMovement() {

	forwardVector := p.getForwardVector()

	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {

		p.Velocity.X += forwardVector.X * p.AcceleratePower
		p.Velocity.Y += forwardVector.Y * p.AcceleratePower
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
	p.Collider.Position.X = p.shape.Position.X
	p.Collider.Position.Y = p.shape.Position.Y

	bounceBack(&p.shape.Position, &p.Velocity, constants.SCREEN_WIDTH, constants.SCREEN_HEIGHT)
}

// Draw method for the Player (draws an isosceles triangle)
func (p *Player) Draw(screen *ebiten.Image) {
	p.shape.Draw(screen)
	for _, v := range p.trailBubbles {
		v.shape.Draw(screen)
	}
	p.Collider.Draw(screen)
	for _, v := range p.Bullets {
		v.Draw(screen)
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

type Bullet = Asteroid

func (p *Player) spawnBullet() *Bullet {
	fmt.Println("Spawned bulled!!")
	p.LastTimeShoot = time.Now()
	shotVector := p.getForwardVector()
	shotVector.X *= constants.BULLET_SPEED
	shotVector.Y *= constants.BULLET_SPEED

	return &Bullet{
		CircleShape: CircleShape{
			Shape: Shape{
				Position: Vec2{
					X: p.shape.Position.X,
					Y: p.shape.Position.Y,
				},
				StrokeWidth: 1,
				Color:       color.RGBA{100, 10, 200, 255},
			},
			Radius: 5,
		},
		Velocity: shotVector,
	}
}

func (p *Player) GetBullets() []*Bullet {
	return p.Bullets
}
