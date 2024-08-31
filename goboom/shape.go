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
	InputManager
}

func NewGameObject() GameObject {
	return GameObject{
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
}

func (obj *GameObject) OnInit() {}

func (obj *GameObject) OnUpdate() {

	obj.Position.X += obj.VelX
	obj.Position.Y += obj.VelY

}