package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type EllipseComp struct {
	CompSetGet
	FillColor 		rl.Color
	StrokeColor 	rl.Color
	StrokeWeight 	float32
}


func Ellipse(x, y, w, h float32, fill, stroke rl.Color, strokeWeight float32) *GameObject {
	obj := NewGameObject()
	obj.X = x
	obj.Y = y
	obj.Width = w
	obj.Height = h

	comp := NewEllipseComp(fill, stroke, strokeWeight)
	obj.AddComponent(comp)
	
	return obj
}

// convenience function for creating a circle
func Circle(x, y, r float32, fill, stroke rl.Color, strokeWeight float32) *GameObject {
	return Ellipse(x, y, r*2, r*2, fill, stroke, strokeWeight)
}

func NewEllipseComp(fill, stroke rl.Color, strokeWeight float32) *EllipseComp {
	comp := &EllipseComp{
		FillColor: fill,
		StrokeColor: stroke,
		StrokeWeight: strokeWeight,
	}
	return comp
}

func (e *EllipseComp) GetComponentId() string {
	return "ellipse"
}

func (e *EllipseComp) OnInit() {}

func (e *EllipseComp) OnUpdate(scene *GameObject) {}

func (e *EllipseComp) OnDraw(scene *GameObject) {

	obj := e.GameObject

	rl.DrawEllipse(int32(obj.Width/2), int32(obj.Height/2), obj.Height/2, obj.Width/2, e.FillColor)
	rl.DrawEllipseLines(int32(obj.Width/2), int32(obj.Height/2), obj.Height/2, obj.Width/2, e.StrokeColor)
	
	// centerpoint for debugging
	// rl.DrawCircle(int32(obj.Width/2), int32(obj.Height/2), 2, rl.Black)
	// top left corner for debugging
	// rl.DrawCircle(0, 0, 2, rl.Black)
}