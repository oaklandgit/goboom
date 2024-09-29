package main

import (
	"fmt"
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const BALL_RADIUS = 12

func bounce (b *boom.GameObject, v rl.Vector2) {
	vel := b.GetComponent("velocity").(*boom.VelocityComp)
	currVel := vel.GetVelVector()
	vel.SetVelocity(rl.Vector2Reflect(currVel, v))
}

func createBall(scene *boom.GameObject, velx, velY float32) *boom.GameObject {

	// define components
	vel := boom.NewVelocityComp(velx, velY)
	collision := boom.NewCollideComp(boom.CollisionCircle{Radius: BALL_RADIUS}, "ball")

	
	collision.NewCollider("brick", func(b, brick *boom.GameObject) {
		brick.SetLifespan(0)
		bounce(b, rl.NewVector2(0, 1))

		scoreboard := scene.GetById("scoreboard")
		score++
		scoreboard.GetComponent("text").(*boom.TextComp).SetText(fmt.Sprintf("Score: %d", score))
	})

	collision.NewCollider("wall", func(b, wall *boom.GameObject) {
		bounce(b, rl.NewVector2(1, 0))
	})

	collision.NewCollider("ceiling", func(b, wall *boom.GameObject) {
		bounce(b, rl.NewVector2(0, -1))
	})

	collision.NewCollider("paddle", func(b, paddle *boom.GameObject) {
		bounce(b, rl.NewVector2(0, 1))
	})

	ball := boom.Ellipse(0, 0, BALL_RADIUS, BALL_RADIUS, rl.White, rl.Red, 2)
	ball.AddTags("ball")
	ball.SetId("the ball")

	ball.AddComponents(vel, collision)

	return ball
}