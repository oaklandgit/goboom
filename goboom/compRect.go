package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RectComp struct {
	CompSetGet
	FillColor 		rl.Color
	StrokeColor 	rl.Color
	StrokeWeight 	float32
}


func Rectangle(x, y, w, h float32, fill, stroke rl.Color, strokeWeight float32) *GameObject {
	obj := NewGameObject()
	obj.X = x
	obj.Y = y
	obj.Width = w
	obj.Height = h

	comp := NewRectComp(fill, stroke, strokeWeight)
	obj.AddComponent(comp)
	
	return obj
}

func NewRectComp(fill, stroke rl.Color, strokeWeight float32) *RectComp {
	comp := &RectComp{
		FillColor: fill,
		StrokeColor: stroke,
		StrokeWeight: strokeWeight,
	}
	return comp
}

func (c *RectComp) GetComponentId() string {
	return "circle"
}

func (c *RectComp) OnInit() {}

func (c *RectComp) OnUpdate(scene *GameObject) {}

func (c *RectComp) OnDraw(scene *GameObject) {

	obj := c.GameObject

	
	rect := rl.Rectangle{X: 0, Y: 0, Width: obj.Width, Height: obj.Height}

	rl.DrawRectanglePro(
		rect,
		rl.Vector2{X: 0, Y: 0},
		0,
		c.FillColor)

	rl.DrawRectangleLinesEx(
		rect,
		c.StrokeWeight,
		c.StrokeColor)

}