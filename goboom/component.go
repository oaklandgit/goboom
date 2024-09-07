package goboom

import "fmt"

type Component interface {
	SetGameObject(g *GameObject)
	OnInit()
	OnUpdate()
	OnDraw()
	GetComponentId() string
}

func Create(x, y float32, c ...Component) *GameObject {	
	obj := NewGameObject()
	obj.X = x
	obj.Y = y
	for _, comp := range c {
		obj.AddComponent(comp)
	}
	return obj
}

func (g *GameObject) AddComponent(c Component) {
	c.SetGameObject(g)
	g.Components = append(g.Components, c)
}

func (g *GameObject) GetComponents() []Component {
	return g.Components
}




// TEST COMPONENT

type SayMessage struct {
	GameObject *GameObject
	Message string
}


func NewSayMessage(message string) *SayMessage {
	return &SayMessage{
		Message: message,
	}
}

func (c *SayMessage) GetComponentId() string {
	return "SayMessage"
}

func (c *SayMessage) SetGameObject(g *GameObject) {
	c.GameObject = g
}

func (s *SayMessage) OnInit() {
	fmt.Println("Initializing message")
}

func (s *SayMessage) OnUpdate() {
	
	tags := s.GameObject.GetTags()
	fmt.Println(tags)
}

func (s *SayMessage) OnDraw() {
	fmt.Println(s.Message)
}