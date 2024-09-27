package goboom


type TweenComp struct {
	GameObject 		*GameObject
	ChangeFunc 		func()
	EndCondition 	func() bool
	AfterFunc 		func()
}


func NewTweenComp() *TweenComp {
	comp := &TweenComp{}
	return comp
}

func (c *TweenComp) GetComponentId() string {
	return "tween"
}

func (c *TweenComp) SetGameObject(g *GameObject) {
	c.GameObject = g
}

func (c *TweenComp) GetGameObject() *GameObject {
	return c.GameObject
}

func (c *TweenComp) OnInit() {}

func (c *TweenComp) OnUpdate(scene *GameObject) {

	if c.EndCondition != nil {
		if c.EndCondition() {
			if c.AfterFunc != nil {
				c.AfterFunc()
			}
			return
		}
	}
	c.ChangeFunc()

}

func (c *TweenComp) OnDraw(scene *GameObject) {}
