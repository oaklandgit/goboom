package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const CIRCLE_EDGES = 24

type Circle struct {
	Shape
	Radius float32
}

func NewCircle(x, y, radius float32, strokeColor rl.Color) *Circle {
	c := &Circle{
		Shape: NewShape(),
	}
	c.X = x
	c.Y = y
	c.StrokeColor = strokeColor
	c.Radius = radius
	return c
}

func (c *Circle) SetRadius(radius float32) *Circle {
	c.Radius = radius
	return c
}

func (c *Circle) OnDraw() {
	// add the radius to the x and y to center the circle
	// to let the GameObject handle the pivot

	// FILL
	if c.FillColor != rl.Blank {
		rl.DrawPoly(rl.Vector2{X: c.X + c.Radius, Y: c.Y + c.Radius},
			CIRCLE_EDGES,
			c.Radius,
			0, // rotation handled by the GameObject
			c.FillColor)
	}

	// STROKE
	if c.StrokeColor == rl.Blank {
	rl.DrawPolyLinesEx(
		rl.Vector2{X: c.X + c.Radius, Y: c.Y + c.Radius},
		CIRCLE_EDGES,
		c.Radius,
		0, // rotation handled by the GameObject
		c.StrokeWeight,
		c.StrokeColor)
	}
	
}

func (c *Circle) GetWidth() float32{
	return c.Radius * 2
}

func (c *Circle) GetHeight() float32{
	return c.Radius * 2
}
