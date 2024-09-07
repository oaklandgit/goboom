package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RectComp struct {
	GameObject 		*GameObject
	Width 			float32
	Height 			float32
	FillColor 		rl.Color
	StrokeColor 	rl.Color
	StrokeWeight 	float32
}


func Rectangle(x, y, w, h float32, fill, stroke rl.Color, strokeWeight float32) *GameObject {
	obj := NewGameObject()
	obj.X = x
	obj.Y = y

	comp := &RectComp{
		Width: w,
		Height: h,
		FillColor: fill,
		StrokeColor: stroke,
		StrokeWeight: strokeWeight,
	}

	obj.AddComponent(comp)
	
	return obj
}

func (c *RectComp) GetComponentId() string {
	return "circle"
}

func (c *RectComp) SetGameObject(g *GameObject) {
	c.GameObject = g
}

func (c *RectComp) OnInit() {}

func (c *RectComp) OnUpdate() {}

func (c *RectComp) OnDraw() {
	obj := c.GameObject
	centerX := obj.X + c.GetWidth() * obj.GetOriginX()
	centerY := obj.Y + c.GetHeight() * obj.GetOriginY()

	rl.PushMatrix()
	rl.Translatef(centerX, centerY, 0)
	rl.Scalef(obj.GetScaleX(), obj.GetScaleY(), 1)
	rl.Rotatef(obj.GetAngle(), 0, 0, 1)
	rl.Translatef(-centerX, -centerY, 0)
	
	
	rect := rl.Rectangle{X: obj.X, Y: obj.Y, Width: c.Width, Height: c.Height}

	rl.DrawRectanglePro(
		rect,
		rl.Vector2{X: 0, Y: 0},
		0,
		c.FillColor)

	rl.DrawRectangleLinesEx(
		rect,
		c.StrokeWeight,
		c.StrokeColor)

	rl.PopMatrix()
}

func (c *RectComp) GetWidth() float32 {
	return c.Width * c.GameObject.GetScaleX()
}

func (c *RectComp) GetHeight() float32 {
	return c.Height * c.GameObject.GetScaleY()
}