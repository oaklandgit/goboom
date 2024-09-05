package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func createBrick(color rl.Color) *boom.GameObject {

	brick := boom.NewRectangle(0, 0, 40, 20, rl.Blank)
	brick.SetFill(color)
	brick.AddTags("brick")

	return brick
}