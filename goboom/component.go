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
	ModifyWidth(float32) float32
	ModifyHeight(float32) float32
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

// Alias
func (g *GameObject) With(c Component) {
	g.AddComponent(c)
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