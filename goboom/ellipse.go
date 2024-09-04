package goboom

import rl "github.com/gen2brain/raylib-go/raylib"

func NewEllipse(x, y, radiusH, radiusV float32, strokeColor rl.Color) *GameObject {
	e := NewGameObject()
	e.X = x
	e.Y = y
	e.StrokeColor = strokeColor

	e.OnDraw = func() {
		rl.DrawEllipse(
			0,
			0,
			radiusH, radiusV, e.FillColor)
	
		// simulate stroke weight
		for i := e.StrokeWeight; i > 0; i-- {
			rl.DrawEllipseLines(
				0,
				0,
				radiusH + i, radiusV + i, e.StrokeColor)
		}
	}

	return e
}
