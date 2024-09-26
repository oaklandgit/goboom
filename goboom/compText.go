package goboom

import rl "github.com/gen2brain/raylib-go/raylib"

type TextComp struct {
	GameObject *GameObject
	Color	  rl.Color
	Text	  string
	FontSize  int32
}

func Text(x, y float32, text string, size int32, color rl.Color) *GameObject {
	obj := NewGameObject()
	obj.X = x
	obj.Y = y

	comp := NewTextComp(text, size, color)
	obj.AddComponent(comp)
	
	return obj
}

func NewTextComp(text string, size int32, color rl.Color) *TextComp {
	comp := &TextComp{
		Text: text,
		FontSize: size,
		Color: color,
	}
	return comp
}

func (t *TextComp) GetComponentId() string {
	return "text"
}

func (t *TextComp) SetGameObject(g *GameObject) {
	t.GameObject = g
}

func (t *TextComp) GetGameObject() *GameObject {
	return t.GameObject
}

func (t *TextComp) OnInit() {
	obj := t.GameObject
	obj.Width = float32(rl.MeasureText(t.Text, t.FontSize))
	obj.Height = float32(rl.MeasureText("M", t.FontSize))
}

func (t *TextComp) OnUpdate(scene *GameObject) {}

func (t *TextComp) OnDraw(scene *GameObject) {
	rl.DrawText(t.Text, 0, 0, t.FontSize, t.Color)
}