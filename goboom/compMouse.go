package goboom

import rl "github.com/gen2brain/raylib-go/raylib"

type MouseComp struct {
	GameObject 	*GameObject
	Move 		func(x, y float32)
	Click 		func()
	Release 	func()
}


func NewMouseComp() *MouseComp {
	comp := &MouseComp{}
	return comp
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

func (c *MouseComp) SetGameObject(g *GameObject) {
	c.GameObject = g
}

func (c *MouseComp) GetGameObject() *GameObject {
	return c.GameObject
}

func (c *MouseComp) OnInit() {}

func (c *MouseComp) OnUpdate(scene *GameObject) {

	x := rl.GetMouseX()
	y := rl.GetMouseY()

	if c.Move != nil {
		c.Move(float32(x), float32(y))
	}

}

func (c *MouseComp) OnDraw(scene *GameObject) {}
