package goboom

type Component interface {
	SetGameObject(g *GameObject)
	GetGameObject() *GameObject
	OnInit()
	OnUpdate(scene *GameObject)
	OnDraw(scene *GameObject)
	GetComponentId() string
}

func Create(x, y, w, h float32, c ...Component) *GameObject {	
	obj := NewGameObject()
	obj.X = x
	obj.Y = y
	obj.Width = w
	obj.Height = h
	for _, comp := range c {
		obj.AddComponent(comp)
	}
	return obj
}

func (g *GameObject) AddComponent(c Component) {
	c.SetGameObject(g)
	g.Components = append(g.Components, c)
}

func (g *GameObject) AddComponents(c ...Component) {
	for _, comp := range c {
		g.AddComponent(comp)
	}
}

func (g *GameObject) GetComponents() []Component {
	return g.Components
}

func (g *GameObject) GetComponent(id string) Component {
	for _, c := range g.Components {
		if c.GetComponentId() == id {
			return c
		}
	}
	return nil
}

func (g *GameObject) HasComponent(id string) bool {
	for _, c := range g.Components {
		if c.GetComponentId() == id {
			return true
		}
	}
	return false
}

func (g *GameObject) Get(id string) Component {
	for _, c := range g.Components {
		if c.GetComponentId() == id {
			return c
		}
	}
	return nil
}