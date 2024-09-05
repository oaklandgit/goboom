package goboom

import rl "github.com/gen2brain/raylib-go/raylib"


func NewGroup(x, y float32, children ...*GameObject) *GameObject {
	g := NewGameObject()
	g.X = x
	g.Y = y
	g.Add(children...)
    g.OnDraw = func() {
        rl.DrawRectangle(int32(g.GetX()), int32(g.GetY()), int32(g.GetWidth()), int32(g.GetHeight()), rl.Yellow)
    }
    g.GetWidth = func() float32 {
        minX := float32(0)
        maxX := float32(0)
        for _, child := range g.Children {
          
            if child.GetX() < minX {
                minX = child.GetX()
            }
            if child.GetX() + child.GetWidth() > maxX {
                maxX = child.GetX() + child.GetWidth()
            }
        }
        return maxX - minX
    }
    g.GetHeight = func() float32 {
        minY := float32(0)
        maxY := float32(0)
        for _, child := range g.Children {
            if child.GetY() < minY {
                minY = child.GetY()
            }
            if child.GetY() + child.GetHeight() > maxY {
                maxY = child.GetY() + child.GetHeight()
            }
        }
        return maxY - minY
    }
	return g
}
