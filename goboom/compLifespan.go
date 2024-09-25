package goboom

import "time"

type LifespanComp struct {
	GameObject *GameObject
	Delete bool
}

func NewLifespanComp() *LifespanComp {
	comp := &LifespanComp{}
	return comp
}

func (l *LifespanComp) SetLifespan(millisecs int) {
	
	go func() {
        time.Sleep(time.Duration(millisecs) * time.Millisecond)
        l.Delete = true
    }()

}


func (l *LifespanComp) GetComponentId() string {
	return "lifespan"
}

func (l *LifespanComp) SetGameObject(g *GameObject) {
	l.GameObject = g
}

func (l *LifespanComp) GetGameObject() *GameObject {
	return l.GameObject
}


func (l *LifespanComp) OnUpdate(scene *GameObject) {
	if l.Delete {
		l.GameObject.GetParent().Remove(l.GameObject)
	}
}

func (l *LifespanComp) OnInit() {}
func (l *LifespanComp) OnDraw(scene *GameObject) {}