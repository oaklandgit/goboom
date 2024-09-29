package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func createEarth(scene *boom.GameObject) *boom.GameObject {

	floor := scene.GetGame().Height - 100
	width := scene.GetGame().Width
	earth := boom.Rectangle(-200, floor, width + 200, 100, rl.Blue, rl.Black, 1)
	collide := boom.NewCollideComp(boom.CollisionRect{Width: 1_000, Height: 100})
	earth.AddComponent(collide)
	earth.AddTags("earth")

	return earth

}