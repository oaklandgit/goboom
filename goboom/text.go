package goboom

import rl "github.com/gen2brain/raylib-go/raylib"

const FONT_SIZE = 16


func NewText(x, y float32, txt *string, color rl.Color) *GameObject {
	
	t := NewGameObject()

	// t.FillColor = color
	// t.X = x
	// t.Y = y
	// t.OnDraw = func () {
	// 	rl.DrawText(
	// 		*txt,
	// 		int32(t.X),
	// 		int32(t.Y),
	// 		int32(FONT_SIZE), t.FillColor)
	// }
	// t.GetWidth = func() float32 {
	// 	return float32(rl.MeasureText(*txt, FONT_SIZE))
	// }
	// t.GetHeight = func() float32 {
	// 	return float32(rl.MeasureText("M", FONT_SIZE))
	// }

	return t
}