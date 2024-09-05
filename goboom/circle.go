package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const CIRCLE_EDGES = 24

func NewCircle(x, y, radius float32, strokeColor rl.Color) *GameObject {

	c := NewGameObject()

	c.X = x
	c.Y = y
	c.StrokeColor = strokeColor

	c.OnDraw = func () {
		// FILL
		if c.FillColor != rl.Blank {
			rl.DrawPoly(rl.Vector2{X: c.X + radius, Y: c.Y + radius},
				CIRCLE_EDGES,
				radius,
				0, // rotation handled by the GameObject
				c.FillColor)
		}

		// STROKE
		if c.StrokeColor != rl.Blank {
			rl.DrawPolyLinesEx(
				rl.Vector2{X: c.X + radius, Y: c.Y + radius},
				CIRCLE_EDGES,
				radius,
				0, // rotation handled by the GameObject
				c.StrokeWeight,
				c.StrokeColor)
		}
	}

	c.GetWidth = func() float32 {
		return radius * 2 * c.ScaleX
	}

	c.GetHeight = func() float32 {
		return radius * 2 * c.ScaleY
	}
	return c
}