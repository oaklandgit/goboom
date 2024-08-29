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
	}
}

func (s *Shape) OnInit() {}

func (s *Shape) OnUpdate() {}


// func (s *Shape) SetColor(color rl.Color) *Shape {
// 	s.Color = color
// 	return s
// }

// func (s *Shape) OnInit() {}

// func (s *Shape) OnDraw() {}

// func (s *Shape) OnUpdate() {}