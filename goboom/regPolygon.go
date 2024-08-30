package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RegPoly struct {
	Shape
	Sides int32
	Radius float32
}

func NewRegPoly(x, y float32, sides int32, radius float32, strokeColor rl.Color) *RegPoly {
	rp := &RegPoly{
		Shape: NewShape(),
	}
	rp.Sides = sides
	rp.X = x
	rp.Y = y
	rp.Radius = radius
	rp.StrokeColor = strokeColor
	return rp
}

func (rp *RegPoly) SetRadius(radius float32) *RegPoly {
	rp.Radius = radius
	return rp
}

func (rp *RegPoly) OnDraw() {
	// add the radius to the x and y to center the polygon
	// to let the GameObject handle the pivot
	center := rl.Vector2{X: rp.X + rp.Radius, Y: rp.Y + rp.Radius}
	rl.DrawPoly(center, rp.Sides, rp.Radius, 0, rp.FillColor)
	rl.DrawPolyLinesEx(center, rp.Sides,rp.Radius, 0, rp.StrokeWeight, rp.StrokeColor); 
}

func (rp *RegPoly) GetWidth() float32{
	return rp.Radius * 2
}

func (rp *RegPoly) GetHeight() float32{
	return rp.Radius * 2
}
