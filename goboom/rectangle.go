package goboom

import rl "github.com/gen2brain/raylib-go/raylib"

func NewRectangle(x, y, width, height float32, strokeColor rl.Color) *GameObject {

	r := NewGameObject()
	r.X = x
	r.Y = y
	r.StrokeColor = strokeColor

	r.OnDraw = func() {
		rect := rl.Rectangle{X: r.X, Y: r.Y, Width: width, Height: height}

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

	r.GetWidth = func() float32 {
		return width * r.ScaleX
	}

	r.GetHeight = func() float32 {
		return height * r.ScaleY
	}

	return r
}

// func drawRectangle(r *GameObject) {

// 	rect := rl.Rectangle{X: r.X, Y: r.Y, Width: width, Height: height}

// 	rl.DrawRectanglePro(
// 		rect,
// 		rl.Vector2{X: 0, Y: 0},
// 		0,
// 		r.FillColor)

// 	rl.DrawRectangleLinesEx(
// 		rect,
// 		r.StrokeWeight,
// 		r.StrokeColor)
// }
