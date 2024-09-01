package goboom

type LifeCycle struct {
	OnInit   func()
	OnUpdate func()
	OnDraw   func()
}

func (lc *LifeCycle) Init() {
	if lc.OnInit != nil {
		lc.OnInit()
	}
}

func (lc *LifeCycle) Update() {
	if lc.OnUpdate != nil {
		lc.OnUpdate()
	}
}

func (lc *LifeCycle) Draw() {
	if lc.OnDraw != nil {
		lc.OnDraw()
	}
}
