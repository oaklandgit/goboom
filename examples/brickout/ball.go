package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func createBall() *boom.GameObject {

	ball := boom.Create(
		0, 0,
		boom.NewCircleComp(6, rl.White, rl.Red, 2),
		boom.NewVelocityComp(-1, -1),
	)
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