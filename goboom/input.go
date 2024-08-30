package goboom

import rl "github.com/gen2brain/raylib-go/raylib"

type ButtonPress int

const (
	KeyPressed = iota
	KeyReleased
	KeyDown
)

type Input struct {
	Key   int32
	Press ButtonPress
	Action func()
}

func (g *Game) NewInput(key int32, press ButtonPress, action func()) {

	input := Input{
		Key:   key,
		Press: press,
		Action: action,
	}

	g.Inputs = append(g.Inputs, input)
}

func CheckInput(input Input) {

	if input.Press == KeyPressed && rl.IsKeyPressed(input.Key) {
		input.Action()
		return
	}

	if input.Press == KeyDown && rl.IsKeyDown(input.Key) {
		input.Action()
		return
	}

	if input.Press == KeyReleased && rl.IsKeyReleased(input.Key) {
		input.Action()
		return
	}

}
