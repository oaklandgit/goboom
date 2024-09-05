package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameObject struct {
	OnInit   func()
	OnUpdate func()
	OnDraw   func()
	
	Identity
	Wrapping
	Position
	Rotation
	Lifespan
	Scale
	Velocity
	Visibility
	Shape
	Stroke
	Fill
	Alpha
	Pivot
	InputHandler

	Game *Game
	Parent *GameObject
	Children []*GameObject
}

func NewGameObject() *GameObject {

	obj := GameObject{
		Scale: Scale{
			ScaleX: 1,
			ScaleY: 1,
		},
		Stroke: Stroke{
			StrokeWeight: 2,
			StrokeColor: rl.Magenta, // default stroke color
		},
		Fill: Fill{
			FillColor: rl.Blank,
		},
		Visibility: Visibility{
			Visible: true,
		},
	}

	obj.OnUpdate = func() {
		obj.X += obj.VelX
		obj.Y += obj.VelY
	}

	return &obj
}

// func (g *GameObject) Update() {
// 	g.X += g.VelX
// 	g.Y += g.VelY
// }

// func (lc *LifeCycle) Init() {
// 	if lc.OnInit != nil {
// 		lc.OnInit()
// 	}
// }

// func (lc *LifeCycle) Update() {
// 	if lc.OnUpdate != nil {
// 		lc.OnUpdate()
// 	}
// }

// func (s *Shape) Draw() {
// 	if s.OnDraw != nil {
// 		s.OnDraw()
// 	}
// }