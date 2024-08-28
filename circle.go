package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Circle struct {
	Shape
	Radius float32
}

func NewCircle(x, y, radius float32, color rl.Color) *Circle {
	c := &Circle{}
	c.X = x
	c.Y = y
	c.Radius = radius
	c.Color = color
	return c
}

func (c *Circle) SetRadius(radius float32) *Circle {
	c.Radius = radius
	return c
}

func (c *Circle) OnDraw() {
	rl.DrawCircleLines(int32(c.X), int32(c.Y), c.Radius, c.Color)
}

func (c *Circle) GetWidth() float32{
	return c.Radius * 2
}

func (c *Circle) GetHeight() float32{
	return c.Radius * 2
}
