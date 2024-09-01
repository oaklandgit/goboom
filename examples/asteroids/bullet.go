package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func createBullet(x, y, angle float32) boom.Renderable {
	bullet := boom.NewCircle(x, y, 3, rl.Blank)
	bullet.SetFill(rl.White)
	bullet.SetOrigin(0.5, 0.5)
	bullet.SetVelocityByHeading(angle, 5)
	bullet.SetLifespan(1_000) // 1 second
	return bullet
}