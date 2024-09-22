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


	collision := boom.NewCollideComp(boom.CollisionCircle{Radius: 6}, "ball")
	
	collision.NewCollider("brick", func(b, brick *boom.GameObject) {
		brick.SetLifespan(0)
		b.SetVelocityVect(
			rl.Vector2Reflect(b.GetVelocityVect(), rl.NewVector2(0, 1)))

	})

	collision.NewCollider("wall", func(b, wall *boom.GameObject) {
		b.SetVelocityVect(
			rl.Vector2Reflect(b.GetVelocityVect(), rl.NewVector2(1, 0)))
	})

	collision.NewCollider("paddle", func(b, paddle *boom.GameObject) {
		b.SetVelocityVect(
			rl.Vector2Reflect(b.GetVelocityVect(), rl.NewVector2(0, 1)))
	})

	ball.AddComponent(collision)

	return ball
}