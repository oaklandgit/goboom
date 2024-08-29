package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Shape struct {
	Identity
	Position
	Rotation
	Scale
	Visibility
	Size
	Alpha
	Pivot
	Color rl.Color
}

func NewShape() Shape {
	return Shape{
		Scale: Scale{
			ScaleX: 1,
			ScaleY: 1,
		},
		Visibility: Visibility{
			Visible: true,
		},
	}
}

func (s *Shape) OnInit() {}

func (s *Shape) OnUpdate() {}