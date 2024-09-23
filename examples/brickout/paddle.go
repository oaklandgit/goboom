package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func createPaddle() *boom.GameObject {

	paddle := boom.Rectangle(0, 0, 80, 16, rl.White, rl.Blank, 0)
	paddle.AddTags("paddle")

	vel := boom.NewVelocityComp(0, 0)
	control := boom.NewInputComp()
	collision := boom.NewCollideComp(
		boom.CollisionRect{Width: 80, Height: 16}, "paddle")

	
	control.NewInput(rl.KeyLeft, boom.KeyDown, func() {
		vel.SetVelocity(-PADDLE_SPEED, 0)
	})
	
	control.NewInput(rl.KeyRight, boom.KeyDown, func() {
		vel.SetVelocity(PADDLE_SPEED, 0)
	})
	
	control.NewInput(rl.KeyLeft, boom.KeyReleased, func() {
		vel.SetVelocity(0, 0)
	})
	
	control.NewInput(rl.KeyRight, boom.KeyReleased, func() {
		vel.SetVelocity(0, 0)
	})
	
	paddle.SetId("the paddle")
	paddle.AddComponents(vel, control, collision)
	
	return paddle
}