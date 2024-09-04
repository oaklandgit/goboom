package goboom

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewPolygon(x, y float32, strokeColor rl.Color, closed bool, points ...float32) *GameObject {
	p := NewGameObject()
	// p.X = GetTopLeftX(points)
	// p.Y = GetTopLeftY(points)
	p.X = x
	p.Y = y
	p.StrokeColor = strokeColor

	p.OnDraw = func() {
		vectors, err := PointsToVectors(points)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		rl.PushMatrix()
		rl.Translatef(p.X, p.Y, 0)
		rl.DrawSplineLinear(vectors, p.StrokeWeight, p.StrokeColor) 
		rl.PopMatrix()
	}

	p.GetWidth = func() float32 {
		var minX, maxX float32
		minX = points[0]
		maxX = points[0]
	
		for i := 0; i < len(points); i += 2 {
			if points[i] < minX {
				minX = points[i]
			}
			if points[i] > maxX {
				maxX = points[i]
			}
		}
	
		return maxX - minX
	}

	p.GetHeight = func() float32 {
		var minY, maxY float32
		minY = points[1]
		maxY = points[1]

		for i := 1; i < len(points); i += 2 {
			if points[i] < minY {
				minY = points[i]
			}
			if points[i] > maxY {
				maxY = points[i]
			}
		}

		return maxY - minY
	}

	return p
}

func GetTopLeftX(points []float32) float32 {
	minX := points[0]
	for i := 0; i < len(points); i += 2 {
		if points[i] < minX {
			minX = points[i]
		}
	}
	return minX
}

func GetTopLeftY(points []float32) float32 {
	minY := points[1]
	for i := 1; i < len(points); i += 2 {
		if points[i] < minY {
			minY = points[i]
		}
	}
	return minY
}


func PointsToVectors(pts []float32) ([]rl.Vector2, error)  {
	if len(pts)%2 != 0 {
		return nil, fmt.Errorf("points must be in pairs")
	}

	var vecs []rl.Vector2
	for i := 0; i < len(pts); i += 2 {
		vecs = append(vecs, rl.Vector2{X: pts[i], Y: pts[i+1]})
	}
	return vecs, nil
}
