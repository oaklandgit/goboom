package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameObject struct {
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
	LifeCycle
	InputHandler

	Children []*GameObject
	Parent *GameObject
	Game *Game
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


	return &obj
		
}

func (g *GameObject) Update() {
	g.X += g.VelX
	g.Y += g.VelY
}