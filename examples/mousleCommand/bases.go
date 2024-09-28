package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)


func NewBase() *boom.GameObject {

	window := func() *boom.GameObject {
		return boom.Rectangle(0, 0, 12, 12, rl.White, rl.Black, 1)
	}

	base := boom.GridArray(7, 6, 1, window)


	return base
}