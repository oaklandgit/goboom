package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameObject struct {
	Identity
	Position
	Rotation
	Lifespan
	Scale
	Velocity
	Visibility
	Stroke
	Fill
	Size
	BoundingBox
	Alpha
	Pivot
	LifeCycle
	InputHandler
}

func NewGameObject() GameObject {

	obj := GameObject{
		Scale: Scale{
			ScaleX: 1,
			ScaleY: 1,
		},
		Stroke: Stroke{
			StrokeWeight: 1,
			StrokeColor: rl.Magenta, // default stroke color
		},
		Fill: Fill{
			FillColor: rl.Blank,
		},
		Visibility: Visibility{
			Visible: true,
		},
	}


	return obj
		
}

func (g *GameObject) Update() {
	g.X += g.VelX
	g.Y += g.VelY
}