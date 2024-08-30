package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Shape struct {
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

func NewShape() Shape {
	return Shape{
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

func (s *Shape) OnInit() {}

func (s *Shape) OnUpdate() {

	s.Position.X += s.VelX
	s.Position.Y += s.VelY

}