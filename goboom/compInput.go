package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type InputComp struct {
	GameObject 	*GameObject
	Inputs 		[]Input
}


type ButtonMode int

const (
	KeyPressed = iota
	KeyReleased
	KeyDown
)

type Input struct {
	Key   int32
	Mode ButtonMode
	Action func()
}

func NewInputComp(inputs ...Input) *InputComp {
	comp := &InputComp{
		Inputs: inputs,
	}
	return comp
}

func (c *InputComp) NewInput(key int32, mode ButtonMode, action func()) Input {

	input := Input{
		Key:   key,
		Mode: mode,
		Action: action,
	}

	c.Inputs = append(c.Inputs, input)

	return input
}


func (c *InputComp) GetComponentId() string {
	return "input"
}

func (c *InputComp) SetGameObject(g *GameObject) {
	c.GameObject = g
}

func (c *InputComp) OnInit() {}

func (c *InputComp) OnUpdate() {

	
	for _, input := range c.Inputs {

		if input.Mode == KeyPressed && rl.IsKeyPressed(input.Key) {
			input.Action()
		}

		if input.Mode == KeyDown && rl.IsKeyDown(input.Key) {
			input.Action()
		}

		if input.Mode == KeyReleased && rl.IsKeyReleased(input.Key) {
			input.Action()
		}
	}
}

func (c *InputComp) OnDraw() {}
