package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type EllipseComp struct {
	GameObject 		*GameObject
	Width 			float32
	Height 			float32
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

func (e *EllipseComp) SetGameObject(g *GameObject) {
	e.GameObject = g
}

func (e *EllipseComp) GetGameObject() *GameObject {
	return e.GameObject
}

func (e *EllipseComp) OnInit() {}

func (e *EllipseComp) OnUpdate(scene *GameObject) {}

func (e *EllipseComp) OnDraw(scene *GameObject) {

	obj := e.GameObject
	// centerX := obj.X + obj.GetWidth() * obj.GetOriginX()
	// centerY := obj.Y + obj.GetHeight() * obj.GetOriginY()

	// rl.PushMatrix()
	// rl.Translatef(centerX, centerY, 0)
	// rl.Scalef(obj.GetScaleX(), obj.GetScaleY(), 1)
	// rl.Rotatef(obj.GetAngle(), 0, 0, 1)
	// rl.Translatef(-centerX, -centerY, 0)

	rl.DrawEllipse(int32(obj.X), int32(obj.Y), obj.Height, obj.Width, e.FillColor)
	rl.DrawEllipseLines(int32(obj.X), int32(obj.Y), obj.Height, obj.Width, e.StrokeColor)
	
	// rl.PopMatrix()
}

// func (e *EllipseComp) GetWidth() float32 {
// 	return e.Width
// }

// func (e *EllipseComp) GetHeight() float32 {
// 	return e.Height
// }