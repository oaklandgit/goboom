package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Rectangle struct {
	Shape
	Width  float32
	Height float32
}

func NewRectangle(x, y, width, height float32, color rl.Color) *Rectangle {
	r := &Rectangle{
		Shape: NewShape(),
	}
	r.X = x
	r.Y = y
	r.Width = width
	r.Height = height
	r.Color = color
	return r
}

func (r *Rectangle) OnDraw() {
	rl.DrawRectangleLines(int32(r.X), int32(r.Y), int32(r.Width), int32(r.Height), r.Color)
}

func (r *Rectangle) GetWidth() float32 {
	return r.Width
}

func (r *Rectangle) GetHeight() float32 {
	return r.Height
}
