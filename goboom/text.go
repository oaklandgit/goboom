package goboom

import rl "github.com/gen2brain/raylib-go/raylib"

type Text struct {
	GameObject
	Text  string
	FontSize float32
}

func NewText(x, y float32, txt string, color rl.Color) *Text {
	t := &Text{
		GameObject: NewGameObject(),
	}
	t.FillColor = color
	t.X = x
	t.Y = y
	t.Text = txt
	t.FontSize = 16
	return t
}

func (t *Text) SetText(txt string) *Text {
	t.Text = txt
	return t
}

func (t *Text) OnDraw() {

	rl.DrawText(
		t.Text,
		int32(t.X),
		int32(t.Y),
		int32(t.FontSize), t.FillColor)
}

func (t *Text) GetText() string {
	return t.Text
}

func (t *Text) GetWidth() float32 {
	return float32(rl.MeasureText(t.Text, int32(t.FontSize)))
}

func (t *Text) GetHeight() float32 {
	return float32(rl.MeasureText("M", int32(t.FontSize)))
}

