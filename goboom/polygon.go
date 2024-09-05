package goboom

import (
	"fmt"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewPolygon(x, y float32, strokeColor rl.Color, closed bool, path string) *GameObject {

	vectors, err := StringToVectors(path)

	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	p := NewGameObject()
	p.X = x
	p.Y = y
	p.StrokeColor = strokeColor

	p.OnDraw = func() {

		if closed {
			vectors = append(vectors, vectors[0])
		}

		rl.PushMatrix()
		rl.Translatef(p.X + p.GetWidth()/2, p.Y + p.GetHeight()/2, 0)
		rl.DrawSplineLinear(vectors, p.StrokeWeight, p.StrokeColor) 
		rl.PopMatrix()
	}

	p.GetWidth = func() float32 {
		return GetMaxX(vectors) - GetMinX(vectors)
	}

	p.GetHeight = func() float32 {
		return GetMaxY(vectors) - GetMinY(vectors)
	}

	return p
}

func GetMinX(vectors []rl.Vector2) float32 {
    if len(vectors) == 0 {
        return 0
    }
    minX := vectors[0].X
    for _, point := range vectors {
        if point.X < minX {
            minX = point.X
        }
    }
    return minX
}

func GetMinY(vectors []rl.Vector2) float32 {
    if len(vectors) == 0 {
        return 0
    }
    minY := vectors[0].Y
    for _, point := range vectors {
        if point.Y < minY {
            minY = point.Y
        }
    }
    return minY
}

func GetMaxX(vectors []rl.Vector2) float32 {
    if len(vectors) == 0 {
        return 0
    }
    maxX := vectors[0].X
    for _, point := range vectors {
        if point.X > maxX {
            maxX = point.X
        }
    }
    return maxX
}

func GetMaxY(vectors []rl.Vector2) float32 {
    if len(vectors) == 0 {
        return 0
    }
    maxY := vectors[0].Y
    for _, point := range vectors {
        if point.Y > maxY {
            maxY = point.Y
        }
    }
    return maxY
}

func StringToVectors(input string) ([]rl.Vector2, error) {
	parts := strings.Fields(input) // Split by spaces
	if len(parts)%2 != 0 {
		return nil, fmt.Errorf("input does not have an even number of coordinates")
	}

	vectors := []rl.Vector2{}

	for i := 0; i < len(parts); i += 2 {
		x, err := strconv.ParseFloat(parts[i], 32)
		if err != nil {
			return nil, fmt.Errorf("error parsing X value: %v", err)
		}
		y, err := strconv.ParseFloat(parts[i+1], 32)
		if err != nil {
			return nil, fmt.Errorf("error parsing Y value: %v", err)
		}
		vectors = append(vectors, rl.Vector2{X: float32(x), Y: float32(y)})
	}

	return vectors, nil
}

