package main

import (
	boom "goboom/goboom"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	LEFTMOST_STARTX = 40
	LEFTMOST_ANGLE = 160
	RIGHTMOST_ANGLE = 200
)

func spawnMissiles(scene *boom.GameObject) {

	delay := 5_000 * time.Millisecond
	speed := 0.5
	ticker := time.NewTicker(delay)
	defer ticker.Stop()

	for range ticker.C {
		angle := boom.RandIntBetween(LEFTMOST_ANGLE, RIGHTMOST_ANGLE) // towards the base. (180 is straight down)
		posX := boom.RandIntBetween(LEFTMOST_STARTX, int(scene.GetGame().GetWidth() - LEFTMOST_STARTX))
		createMissile(scene, float32(posX), angle, float32(speed), rl.Red)

		// reduce delay by 1/2 second each time
		delay -= 250 * time.Millisecond
		if delay < 250 * time.Millisecond {
			delay = 250 * time.Millisecond
		}

		// increase speed by 0.1 each time
		speed += 0.1

		// Reset the timer with the new delay
		ticker.Reset(delay)
	
	}
}

func createMissile(scene *boom.GameObject, posX float32, angle int, speed float32, color rl.Color) {

	missile := boom.Circle(posX, 0, 3, color, rl.Blank, 2)
	missile.AddTags("missile")
	missile.SetOrigin(-0.5, -0.5)
	vel := boom.NewVelocityComp(0, 0)
	vel.SetVelocityByHeading(float32(angle), speed)
	collide := boom.NewCollideComp(boom.CollisionCircle{Radius: 2})

	collide.NewCollider("base", func(m, b *boom.GameObject) {
		m.SetLifespan(0)
		b.SetLifespan(0)
		createBurst(scene, m.GetGlobalX(), m.GetGlobalY(), rl.Red)
	})

	collide.NewCollider("earth", func(m, e *boom.GameObject) {
		createBurst(scene, m.GetGlobalX(), m.GetGlobalY(), rl.White)
		m.SetLifespan(0)
	})

	missile.AddComponents(vel, collide)

	missile.OnDraw = func() {
		rl.DrawLine(int32(posX-missile.GetX()), int32(-missile.GetY()), 3, 3, color)
	}
	scene.Add(missile)
}