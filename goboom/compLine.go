package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type LineComp struct {
	CompSetGet
	End 			rl.Vector2 // start is determined by the GameObject's x, y
	StrokeColor 	rl.Color
	StrokeWeight 	float32
}

func Line(x1, y1, x2, y2 float32, color rl.Color, weight float32) *GameObject {
	obj := NewGameObject()
	obj.SetXY(x1, y1)

	comp := NewLineComp(rl.NewVector2(x2, y2), color, weight)

	obj.AddComponent(comp)
	return obj
}

func NewLineComp(end rl.Vector2, color rl.Color, weight float32) *LineComp {
	comp := &LineComp{
		End: end,
		StrokeColor: color,
		StrokeWeight: weight,
	}

	return comp
}


func (c *LineComp) GetComponentId() string {
	return "line"
}

func (c *LineComp) OnInit() {}

func (c *LineComp) OnUpdate(scene *GameObject) {}

func (c *LineComp) OnDraw(scene *GameObject) {
	
	obj := c.GameObject
	start := rl.NewVector2(obj.GetX(), obj.GetY())
	rl.DrawLineEx(start, c.End, c.StrokeWeight, c.StrokeColor)
	
}
