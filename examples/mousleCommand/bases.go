package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewBase() *boom.GameObject {

	window := func() *boom.GameObject {
		return boom.Rectangle(0, 0, 12, 12, rl.White, rl.Gray, 1)
	}

	base := boom.GridArray(7, 6, 0, window)
	collide := boom.NewCollideComp(boom.CollisionRect{Width: 100, Height: 100})
	base.AddComponent(collide)
	base.AddTags("base")

	return base
}