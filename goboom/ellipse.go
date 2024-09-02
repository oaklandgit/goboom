package goboom

import rl "github.com/gen2brain/raylib-go/raylib"

type Ellipse struct {
	GameObject
	RadiusH float32
	RadiusV float32
}

func NewEllipse(x, y, radiusH, radiusV float32, strokeColor rl.Color) *Ellipse {
	e := &Ellipse{
		GameObject: NewGameObject(),
	}
	e.X = x
	e.Y = y
	e.RadiusH = radiusH
	e.RadiusV = radiusV
	e.StrokeColor = strokeColor

	e.OnDraw = func() {
		drawEllipse(e)
	}

	return e
}

func drawEllipse(e *Ellipse) {

	rl.DrawEllipse(
		0,
		0,
		e.RadiusH, e.RadiusV, e.FillColor)

	rl.DrawEllipseLines(
			0,
			0,
			e.RadiusH, e.RadiusV, e.StrokeColor)
}

func (e *Ellipse) GetWidth() float32 {
	return e.RadiusH * e.ScaleX
}

func (e *Ellipse) GetHeight() float32 {
	return e.RadiusH * e.ScaleY
}
