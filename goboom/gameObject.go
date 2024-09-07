package goboom

type GameObject struct {
	OnInit   func()
	OnUpdate func()
	OnDraw   func()

	Identity
	
	Shape
	Stroke
	Fill
	
	Position
	Rotation
	Scale
	Pivot
	
	Velocity
	Wrapping
	
	Lifespan
	InputHandler
	Collisions

	Game *Game
	Parent *GameObject
	Children []*GameObject
	Components []Component
}

func NewGameObject() *GameObject {

	obj := GameObject{
		Scale: Scale{
			ScaleX: 1,
			ScaleY: 1,
		},
		// Stroke: Stroke{
		// 	StrokeWeight: 2,
		// 	StrokeColor: rl.Magenta, // default stroke color
		// },
		// Fill: Fill{
		// 	FillColor: rl.Blank,
		// },
		Shape: Shape{
			Visible: true,
		},
	}

	obj.OnUpdate = func() {
		obj.X += obj.VelX
		obj.Y += obj.VelY
	}

	obj.OnDraw = func() {}

	obj.OnInit = func() {}

	return &obj
}

func (obj *GameObject) GetScene() *GameObject {
	return obj.GetGame().GetCurrentScene()
}