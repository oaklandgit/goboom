package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func createBullet(x, y, angle, thrust float32) *boom.GameObject {
	bullet := boom.NewRectangle(x, y, 1, 20, rl.Red)
	bullet.SetOrigin(0, 0)
	bullet.SetVelocityByHeading(angle, thrust)
	bullet.SetAngle(angle)
	bullet.SetLifespan(1_000) // 1 second
	return bullet
}