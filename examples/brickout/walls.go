package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func createWall(x, y, width, height float32) *boom.GameObject {

	wall := boom.Rectangle(x, y, width, height, rl.DarkBrown, rl.Blank, 2)
	wall.AddTags("wall")
	collision := boom.NewCollideComp(
		boom.CollisionRect{Width: width, Height: height}, "wall")

	wall.AddComponent(collision)

	return wall
}