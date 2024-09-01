package goboom

type Group struct {
	GameObject
}

func NewGroup(x, y float32, children ...Renderable) *Group {
	g := &Group{
		GameObject: NewGameObject(),
	}
	g.X = x
	g.Y = y
	g.Add(children...)
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
