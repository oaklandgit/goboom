package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RegPolyComp struct {
	CompSetGet
	Sides 			int32
	Radius 			float32
	Fill
	Stroke
}


func RegPolygon(x, y, r float32, sides int32, fill, stroke rl.Color, strokeWeight float32) *GameObject {
	obj := NewGameObject()
	obj.X = x
	obj.Y = y
	obj.Width = r * 2
	obj.Height = r * 2

	comp := &RegPolyComp{
		Radius: r,
		Sides: sides,
		Fill: Fill{
			FillColor: fill,
		},
		Stroke: Stroke{
			StrokeColor: stroke,
			StrokeWeight: strokeWeight,
		},
	}

	obj.AddComponent(comp)
	
	return obj
}


func (c *RegPolyComp) GetComponentId() string {
	return "regpolygon"
}

func (c *RegPolyComp) OnInit() {}

func (c *RegPolyComp) OnUpdate(scene *GameObject) {}

func (c *RegPolyComp) OnDraw(scene *GameObject) {
	if c.FillColor != rl.Blank {
		rl.DrawPoly(rl.Vector2{X: c.Radius, Y: c.Radius},
			c.Sides,
			c.Radius,
			0, // rotation handled by the GameObject
			c.FillColor)
	}

	if c.StrokeColor != rl.Blank {
		rl.DrawPolyLinesEx(
			rl.Vector2{X: c.Radius, Y: c.Radius},
			c.Sides,
			c.Radius,
			0, // rotation handled by the GameObject
			c.StrokeWeight,
			c.StrokeColor)
	}

}

func (c *RegPolyComp) GetWidth() float32 {
	return c.Radius * 2
}

func (c *RegPolyComp) GetHeight() float32 {
	return c.Radius * 2
}