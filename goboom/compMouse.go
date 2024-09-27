package goboom

import rl "github.com/gen2brain/raylib-go/raylib"


type MouseComp struct {
	CompSetGet
	Move 		func(x, y float32)
	Click 		func()
	Release 	func()
	Cursor		rl.MouseCursor
}


func NewMouseComp() *MouseComp {
	comp := &MouseComp{}
	return comp
}

func (c *MouseComp) SetCursor(cursor rl.MouseCursor) {
	c.Cursor = cursor
	// rl.SetMouseCursor(c.Cursor)
}

func (c *MouseComp) OnMove(action func(x, y float32)) {
	c.Move = action
}

func (c *MouseComp) OnClick(action func()) {
	c.Click = action
}

func (c *MouseComp) OnRelease(action func()) {
	c.Release = action
}

func (c *MouseComp) GetComponentId() string {
	return "mouse"
}

func (c *MouseComp) OnInit() {
	rl.SetMouseCursor(c.Cursor)
}

func (c *MouseComp) OnUpdate(scene *GameObject) {

	x := rl.GetMouseX()
	y := rl.GetMouseY()

	if c.Move != nil {
		c.Move(float32(x), float32(y))
	}

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		if c.Click != nil {
			c.Click()
		}
	}

}

func (c *MouseComp) OnDraw(scene *GameObject) {}
