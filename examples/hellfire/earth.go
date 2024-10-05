package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func createEarth(scene *boom.GameObject) *boom.GameObject {

	floor := scene.GetGame().Height
	width := scene.GetGame().Width
	earth := boom.Rectangle(0, floor, width, 100, rl.Blue, rl.Black, 1)

	size := earth.GetBoundingBox()
	collide := boom.NewCollideComp(boom.CollisionRect{Width: size.Width, Height: size.Height})
	earth.AddComponent(collide)
	earth.AddTags("earth")

	return earth

}