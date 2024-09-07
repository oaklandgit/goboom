package goboom

import (
	"fmt"
	"reflect"
)

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

func (g *GameObject) Get(id string) Component {
	for _, c := range g.Components {
		if c.GetComponentId() == id {
			return c
		}
	}
	return nil
}

func (obj *GameObject) Do(componentID, methodName string, args ...interface{}) {
    comp := obj.Get(componentID)
    if comp == nil {
        fmt.Printf("Error: component %s not found\n", componentID)
        return
    }

    compValue := reflect.ValueOf(comp)
    method := compValue.MethodByName(methodName)
    if !method.IsValid() {
        fmt.Printf("Error: method %s not found on component %s\n", methodName, componentID)
        return
    }

    methodArgs := make([]reflect.Value, len(args))
    for i, arg := range args {
        methodArgs[i] = reflect.ValueOf(arg)
    }

    method.Call(methodArgs)
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