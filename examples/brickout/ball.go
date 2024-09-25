package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func bounce (b *boom.GameObject, v rl.Vector2) {
	vel := b.GetComponent("velocity").(*boom.VelocityComp)
	currVel := vel.GetVelVector()
	vel.SetVelocity(rl.Vector2Reflect(currVel, v))
}

func createBall(velx, velY float32) *boom.GameObject {

	// define components
	vel := boom.NewVelocityComp(velx, velY)
	collision := boom.NewCollideComp(boom.CollisionCircle{Radius: 6}, "ball")

	
	collision.NewCollider("brick", func(b, brick *boom.GameObject) {
		brick.SetLifespan(0)
		bounce(b, rl.NewVector2(0, 1))
	})

	collision.NewCollider("wall", func(b, wall *boom.GameObject) {
		bounce(b, rl.NewVector2(1, 0))
	})

	collision.NewCollider("paddle", func(b, paddle *boom.GameObject) {
		bounce(b, rl.NewVector2(0, 1))
	})

	ball := boom.Ellipse(0, 0, 6, 6, rl.White, rl.Red, 2)
	ball.AddTags("ball")
	ball.SetId("the ball")

	ball.AddComponents(vel, collision)


	return ball
}