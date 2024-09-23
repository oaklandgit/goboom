package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RegPolyComp struct {
	GameObject 		*GameObject
	Sides 			int32
	Radius 			float32
	Fill
	Stroke
}


func RegPolygon(x, y, r float32, sides int32, fill, stroke rl.Color, strokeWeight float32) *GameObject {
	obj := NewGameObject()
	obj.X = x
	obj.Y = y

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

func (c *RegPolyComp) SetGameObject(g *GameObject) {
	c.GameObject = g
}

func (c *RegPolyComp) GetGameObject() *GameObject {
	return c.GameObject
}

func (c *RegPolyComp) OnInit() {}

func (c *RegPolyComp) OnUpdate(scene *GameObject) {}

func (c *RegPolyComp) OnDraw(scene *GameObject) {
	obj := c.GameObject
	centerX := obj.X + obj.GetWidth() * obj.GetOriginX()
	centerY := obj.Y + obj.GetHeight() * obj.GetOriginY()
	rl.PushMatrix()
	rl.Translatef(centerX, centerY, 0)
	rl.Scalef(obj.GetScaleX(), obj.GetScaleY(), 1)
	rl.Rotatef(obj.GetAngle(), 0, 0, 1)
	rl.Translatef(-centerX, -centerY, 0)
	
	
	if c.FillColor != rl.Blank {
		rl.DrawPoly(rl.Vector2{X: obj.X + c.Radius, Y: obj.Y + c.Radius},
			c.Sides,
			c.Radius,
			0, // rotation handled by the GameObject
			c.FillColor)
	}

	if c.StrokeColor != rl.Blank {
		rl.DrawPolyLinesEx(
			rl.Vector2{X: obj.X + c.Radius, Y: obj.Y + c.Radius},
			c.Sides,
			c.Radius,
			0, // rotation handled by the GameObject
			c.StrokeWeight,
			c.StrokeColor)
	}

	rl.PopMatrix()
}

func (c *RegPolyComp) ModifyWidth(w float32) float32 {
	if c.Radius * 2 > w {
		return c.Radius * 2
	}
	return w
}

func (c *RegPolyComp) ModifyHeight(h float32) float32 {
	if c.Radius * 2 > h {
		return c.Radius * 2
	}
	return h
}