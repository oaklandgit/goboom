package goboom

import rl "github.com/gen2brain/raylib-go/raylib"

type Rectangle struct {
	GameObject
	Width  float32
	Height float32
}

func NewRectangle(x, y, width, height float32, strokeColor rl.Color) *Rectangle {
	r := &Rectangle{
		GameObject: NewGameObject(),
	}
	r.X = x
	r.Y = y
	r.Width = width
	r.Height = height
	r.StrokeColor = strokeColor

	r.OnDraw = func() {
		drawRectangle(r)
	}

	return r
}

func drawRectangle(r *Rectangle) {

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
	return r.Width * r.ScaleX
}

func (r *Rectangle) GetHeight() float32 {
	return r.Height * r.ScaleY
}
