package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func createPaddle() *boom.GameObject {

	paddle := boom.NewRectangle(0, 0, 80, 16, rl.Blank)
	paddle.SetFill(rl.White)
	paddle.AddTags("paddle")

	paddle.AddInput(rl.KeyLeft, boom.KeyDown, func() {
		paddle.SetVelocityByHeading(-90, PADDLE_SPEED)
	})

	paddle.AddInput(rl.KeyRight, boom.KeyDown, func() {
		paddle.SetVelocityByHeading(90, PADDLE_SPEED)
	})

	paddle.AddInput(rl.KeyLeft, boom.KeyReleased, func() {
		paddle.SetVelocity(0, 0)
	})

	paddle.AddInput(rl.KeyRight, boom.KeyReleased, func() {
		paddle.SetVelocity(0, 0)
	})
	
	return paddle
}