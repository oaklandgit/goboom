package goboom

// a group is nothing but a game object that contains other game objects
func NewGroup(x, y float32, children ...*GameObject) *GameObject {
	g := NewGameObject()
    g.SetXY(x, y)
	g.SetOrigin(0.5, 0.5)
    g.Add(children...)
	return g
}
