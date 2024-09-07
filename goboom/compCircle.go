package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const EDGES = 12

type CircleComp struct {
	GameObject 		*GameObject
	Radius 			float32
	FillColor 		rl.Color
	StrokeColor 	rl.Color
	StrokeWeight 	float32
}


func Circle(x, y, r float32, fill, stroke rl.Color, strokeWeight float32) *GameObject {
	obj := NewGameObject()
	obj.X = x
	obj.Y = y

	comp := &CircleComp{
		Radius: r,
		FillColor: fill,
		StrokeColor: stroke,
		StrokeWeight: 1,
	}

	obj.AddComponent(comp)
	
	return obj
}


func (c *CircleComp) GetComponentId() string {
	return "circle"
}

func (c *CircleComp) SetGameObject(g *GameObject) {
	c.GameObject = g
}

func (c *CircleComp) OnInit() {}

func (c *CircleComp) OnUpdate() {}

func (c *CircleComp) OnDraw() {
	obj := c.GameObject
	centerX := obj.X + c.GetWidth() * obj.GetOriginX()
	centerY := obj.Y + c.GetHeight() * obj.GetOriginY()
	rl.PushMatrix()
	rl.Translatef(centerX, centerY, 0)
	rl.Scalef(obj.GetScaleX(), obj.GetScaleY(), 1)
	rl.Rotatef(obj.GetAngle(), 0, 0, 1)
	rl.Translatef(-centerX, -centerY, 0)
	
	
	if c.FillColor != rl.Blank {
		rl.DrawPoly(rl.Vector2{X: obj.X + c.Radius, Y: obj.Y + c.Radius},
			EDGES,
			c.Radius,
			0, // rotation handled by the GameObject
			c.FillColor)
	}

	if c.StrokeColor != rl.Blank {
		rl.DrawPolyLinesEx(
			rl.Vector2{X: obj.X + c.Radius, Y: obj.Y + c.Radius},
			EDGES,
			c.Radius,
			0, // rotation handled by the GameObject
			c.StrokeWeight,
			c.StrokeColor)
	}

	rl.PopMatrix()
}

func (c *CircleComp) GetWidth() float32 {
	return c.Radius * 2 * c.GameObject.GetScaleX()
}

func (c *CircleComp) GetHeight() float32 {
	return c.Radius * 2 * c.GameObject.GetScaleY()
}