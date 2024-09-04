package goboom

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)


func NewGroup(x, y float32, children ...*GameObject) *GameObject {
	g := NewGameObject()
	g.X = x
	g.Y = y
	g.Add(children...)
    g.OnDraw = func() {}
    g.GetWidth = func() float32 {
        return CalculateGroupBounds(g.Children).Width
    }
    g.GetHeight = func() float32 {
        return CalculateGroupBounds(g.Children).Height
    }
	return g
}

func CalculateGroupBounds(children []*GameObject) rl.Rectangle {
    // Initialize bounds
    minX, minY := float32(math.MaxFloat32), float32(math.MaxFloat32)
    maxX, maxY := float32(-math.MaxFloat32), float32(-math.MaxFloat32)

    // Iterate over children to find the bounds
    for _, child := range children {
        // Calculate the child's bounding box
        childX := GetGlobalX(child)
        childY := GetGlobalY(child)
        childWidth := child.GetWidth()
        childHeight := child.GetHeight()

        // Update min and max bounds
        if childX < minX {
            minX = childX
        }
        if childY < minY {
            minY = childY
        }
        if childX+childWidth > maxX {
            maxX = childX + childWidth
        }
        if childY+childHeight > maxY {
            maxY = childY + childHeight
        }
    }

    // Calculate the group's bounding box
    groupWidth := maxX - minX
    groupHeight := maxY - minY

    // Return the group's bounding rectangle
    return rl.Rectangle{
        X:      minX,
        Y:      minY,
        Width:  groupWidth,
        Height: groupHeight,
    }
}