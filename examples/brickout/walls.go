package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func createWall(x, y, width, height float32) *boom.GameObject {
	wall := boom.NewRectangle(x, y, width, height, rl.Blank)
	wall.SetFill(rl.Red)
	wall.AddTags("wall")

	return wall
}