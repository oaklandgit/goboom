package goboom

import (
	"fmt"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PolyComp struct {
	GameObject 		*GameObject
	Points 			[]rl.Vector2
	Stroke
	Fill
}

func Polygon(x, y float32, path string, closed bool, fill, stroke rl.Color, strokeWeight float32) *GameObject {
	obj := NewGameObject()
	obj.X = x
	obj.Y = y

	comp := NewPolyComp(path, closed, fill, stroke, strokeWeight)

	obj.AddComponent(comp)
	return obj
}

func NewPolyComp(path string, closed bool, fill, stroke rl.Color, strokeWeight float32) *PolyComp {
	comp := &PolyComp{
		Stroke: Stroke{
			StrokeWeight: strokeWeight,
			StrokeColor: stroke,
		},
		Fill: Fill{
			FillColor: fill,
		},
	}

	comp.Points, _ = StringToVectors(path)
	if closed {
		comp.Points = append(comp.Points, comp.Points[0])
	}

	return comp
}


func (c *PolyComp) GetComponentId() string {
	return "polygon"
}

func (c *PolyComp) SetGameObject(g *GameObject) {
	c.GameObject = g
}

func (c *PolyComp) GetGameObject() *GameObject {
	return c.GameObject
}

func (c *PolyComp) OnInit() {}

func (c *PolyComp) OnUpdate(scene *GameObject) {}

func (c *PolyComp) OnDraw(scene *GameObject) {
	obj := c.GameObject

	var scaleX, scaleY float32 = 1, 1

	if (obj.Width != 0) {
		scaleX = obj.Width/c.GetNaturalWidth()
	}

	if (obj.Height != 0) {
		scaleY = obj.Height/c.GetNaturalHeight()
	}

	rl.Scalef(scaleX, scaleY, 1)
	
	rl.DrawPixel(0, 0, rl.Red)
	
	rl.DrawSplineLinear(c.Points, c.StrokeWeight, c.StrokeColor)
	// TO DO: raylib doesn't have a fill function for irregular polygons
	// need to implement a custom fill function
	
}

func (c *PolyComp) GetNaturalWidth() float32 {
	return GetMaxX(c.Points) - GetMinX(c.Points)
}

func (c *PolyComp) GetNaturalHeight() float32 {
	return GetMaxY(c.Points) - GetMinY(c.Points)
}

func StringToVectors(input string) ([]rl.Vector2, error) {
	parts := strings.Fields(input) // Split by spaces
	if len(parts)%2 != 0 {
		// just ignore the last point if it's not a pair
		parts = parts[:len(parts)-1]
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