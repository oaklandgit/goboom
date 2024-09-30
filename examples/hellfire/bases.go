package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewBase() *boom.GameObject {

	window := func() *boom.GameObject {
		w := boom.Rectangle(0, 0, 12, 12, rl.White, rl.Gray, 1)
		w.AddTags("window")
		collide := boom.NewCollideComp(boom.CollisionRect{Width: 12, Height: 12})
		w.AddComponent(collide)
		return w
	}

	base := boom.GridArray(7, 6, 0, window)
	// size := base.GetBoundingBox()
	// base.AddComponent(collide)
	base.AddTags("base")

	return base
}