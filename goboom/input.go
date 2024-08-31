package goboom

import rl "github.com/gen2brain/raylib-go/raylib"

type ButtonMode int

const (
	KeyPressed = iota
	KeyReleased
	KeyDown
)

type InputManager struct {
	Inputs []Input
}

type Input struct {
	Key   int32
	Mode ButtonMode
	Action func()
}

func (im *InputManager) NewInput(key int32, mode ButtonMode, action func()) {

	input := Input{
		Key:   key,
		Mode: mode,
		Action: action,
	}

	im.Inputs = append(im.Inputs, input)
}

func (im *InputManager) OnInput() {

	for _, input := range im.Inputs {

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
