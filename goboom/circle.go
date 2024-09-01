package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const CIRCLE_EDGES = 24

type Circle struct {
	GameObject
	Radius float32
}

func NewCircle(x, y, radius float32, strokeColor rl.Color) *Circle {
	c := &Circle{
		GameObject: NewGameObject(),
	}
	c.X = x
	c.Y = y
	c.StrokeColor = strokeColor
	c.Radius = radius

	c.OnDraw = func () {
		drawCircle(c)
	}
	return c
}

func (c *Circle) SetRadius(radius float32) *Circle {
	c.Radius = radius
	return c
}

func drawCircle(c *Circle) {

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
