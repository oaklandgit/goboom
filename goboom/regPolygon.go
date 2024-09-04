package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)


func NewRegPoly(x, y float32, sides int32, radius float32, strokeColor rl.Color) *GameObject {

	rp := NewGameObject()
	rp.X = x
	rp.Y = y
	rp.StrokeColor = strokeColor

	rp.OnDraw = func() {
		center := rl.Vector2{X: rp.X + radius, Y: rp.Y + radius}
		rl.DrawPoly(center, sides, radius, 0, rp.FillColor)
		rl.DrawPolyLinesEx(center, sides, radius, 0, rp.StrokeWeight, rp.StrokeColor); 
	}
	rp.GetWidth = func() float32 {
		return radius * 2
	}
	rp.GetHeight = func() float32 {
		return radius * 2
	}
	return rp
}
