package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func createBall() *boom.GameObject {

	ball := boom.NewCircle(0, 0, 8, rl.Blank)
	ball.SetFill(rl.White)
	ball.SetVelocityByHeading(-45, 5)
	ball.AddTags("ball")

	ball.AddCollider("brick", func(b, brick *boom.GameObject) {
		brick.SetLifespan(0)
		b.SetVelocityVect(
			rl.Vector2Reflect(b.GetVelocityVect(), rl.NewVector2(0, 1)))
	})

	ball.AddCollider("wall", func(b, wall *boom.GameObject) {
		b.SetVelocityVect(
			rl.Vector2Reflect(b.GetVelocityVect(), rl.NewVector2(1, 0)))
	})

	ball.AddCollider("paddle", func(b, paddle *boom.GameObject) {
		b.SetVelocityVect(
			rl.Vector2Reflect(b.GetVelocityVect(), rl.NewVector2(0, 1)))
	})

	return ball
}