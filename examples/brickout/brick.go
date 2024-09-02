package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func createBrick(color rl.Color) boom.Renderable {

	brick := boom.NewRectangle(0, 0, 40, 20, rl.Blank)
	brick.SetFill(color)

	return brick
}