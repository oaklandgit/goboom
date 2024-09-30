package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func createEarth(scene *boom.GameObject) *boom.GameObject {

	floor := scene.GetGame().Height - 100
	width := scene.GetGame().Width + 400
	earth := boom.Rectangle(-200, floor, width, 100, rl.Blue, rl.Black, 1)

	size := earth.GetBoundingBox()
	collide := boom.NewCollideComp(boom.CollisionRect{Width: size.Width, Height: size.Height})
	earth.AddComponent(collide)
	earth.AddTags("earth")

	return earth

}