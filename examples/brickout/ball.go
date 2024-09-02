package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func createBall() boom.Renderable {

	ball := boom.NewCircle(0, 0, 8, rl.Blank)
	ball.SetFill(rl.White)
	ball.SetVelocityByHeading(-45, 5)

	return ball
}