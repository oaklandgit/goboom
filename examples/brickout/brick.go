package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func createBrick(color rl.Color) *boom.GameObject {

	
	brick := boom.Rectangle(0, 0, 40, 20, color, rl.White, 2)
	collision := boom.NewCollideComp(boom.CollisionRect{Width: 40, Height: 20}, "brick")

	brick.AddTags("brick")
	
	brick.AddComponent(collision)
	return brick
}