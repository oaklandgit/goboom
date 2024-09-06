package goboom

type Component interface {
	Init()
	Update()
	Draw()
}

func (g *GameObject) AddComponent(c Component) {
	g.Components = append(g.Components, c)
}


// TEST COMPONENT

type SayHello struct {
	Hello string
}