package goboom

import rl "github.com/gen2brain/raylib-go/raylib"

type Rectangle struct {
	Shape
	Width  float32
	Height float32
}

func NewRectangle(x, y, width, height float32, strokeColor rl.Color) *Rectangle {
	r := &Rectangle{
		Shape: NewShape(),
	}
	r.X = x
	r.Y = y
	r.Width = width
	r.Height = height
	r.StrokeColor = strokeColor
	return r
}

func (r *Rectangle) OnDraw() {

	rect := rl.Rectangle{X: r.X, Y: r.Y, Width: r.Width, Height: r.Height}

	rl.DrawRectanglePro(
		rect,
		rl.Vector2{X: 0, Y: 0},
		0,
		r.FillColor)

	rl.DrawRectangleLinesEx(
		rect,
		r.StrokeWeight,
		r.StrokeColor)
}

func (r *Rectangle) GetWidth() float32 {
	return r.Width
}

func (r *Rectangle) GetHeight() float32 {
	return r.Height
}
