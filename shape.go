package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Shape struct {
	Identity
	Lifecycle
	Position
	Size
	Hierarchy
	Pivot
	Color rl.Color
}

func (s *Shape) SetColor(color rl.Color) *Shape {
	s.Color = color
	return s
}

func (s *Shape) OnInit() {}

func (s *Shape) OnUpdate() {}