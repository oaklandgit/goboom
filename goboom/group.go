package goboom

type Group struct {
	GameObject
	Children []Renderable
}

func NewGroup(x, y float32) *Group {
	g := &Group{
		GameObject: NewGameObject(),
	}
	g.X = x
	g.Y = y
	return g
}


func (g *Group) GetWidth() float32{
	// To do
	return 0
}

func (g *Group) GetHeight() float32{
	// To do
	return 0
}
