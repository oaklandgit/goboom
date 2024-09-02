package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func createPaddle() boom.Renderable {

	paddle := boom.NewRectangle(0, 0, 80, 16, rl.Blank)
	paddle.SetFill(rl.White)

	return paddle
}