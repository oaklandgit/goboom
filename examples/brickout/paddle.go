package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func createPaddle() *boom.GameObject {

	paddle := boom.Rectangle(0, 0, 80, 16, rl.White, rl.Blank, 0)
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