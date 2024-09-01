package goboom

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Polygon struct {
	GameObject
	Points []float32
	Closed bool
}

func NewPolygon(x, y float32, strokeColor rl.Color, closed bool, points ...float32) *Polygon {
	p := &Polygon{
		GameObject: NewGameObject(),
	}
	p.Points = points
	p.X = x
	p.Y = y
	p.StrokeColor = strokeColor
	p.Closed = closed

	if p.Closed {
		p.Points = append(p.Points, p.Points[0], p.Points[1])
	}

	p.OnDraw = func() {
		drawPolygon(p)
	}

	return p
}


func (p *Polygon) PointsToVectors() ([]rl.Vector2, error)  {
	if len(p.Points)%2 != 0 {
        return nil, fmt.Errorf("points must be in pairs")
    }

	var vectors []rl.Vector2
	for i := 0; i < len(p.Points); i += 2 {
		vectors = append(vectors, rl.Vector2{X: p.Points[i], Y: p.Points[i+1]})
	}
	return vectors, nil
}

func drawPolygon(p *Polygon) {

	vectors, err := p.PointsToVectors()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	rl.PushMatrix()
	rl.Translatef(p.X, p.Y, 0)
	rl.DrawSplineLinear(vectors, p.StrokeWeight, p.StrokeColor) 
	rl.PopMatrix()
}

func (p *Polygon) GetWidth() float32{
	var minX, maxX float32
	minX = p.Points[0]
	maxX = p.Points[0]

	for i := 0; i < len(p.Points); i += 2 {
		if p.Points[i] < minX {
			minX = p.Points[i]
		}
		if p.Points[i] > maxX {
			maxX = p.Points[i]
		}
	}

	return maxX - minX

}

func (p *Polygon) GetHeight() float32{
	var minY, maxY float32
	minY = p.Points[1]
	maxY = p.Points[1]

	for i := 1; i < len(p.Points); i += 2 {
		if p.Points[i] < minY {
			minY = p.Points[i]
		}
		if p.Points[i] > maxY {
			maxY = p.Points[i]
		}
	}

	return maxY - minY
}
