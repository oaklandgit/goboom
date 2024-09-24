package goboom

// a group is nothing but a game object that contains other game objects
func NewGroup(x, y float32, children ...*GameObject) *GameObject {
	g := NewGameObject()
    g.SetXY(x, y)
    g.Add(children...)
	return g
}
