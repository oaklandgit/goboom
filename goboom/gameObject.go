package goboom

import "fmt"

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

	Game *Game
	Parent *GameObject
	Children []*GameObject
	Components []Component
}

func (g *GameObject) String() string {
    return fmt.Sprintf("Game: %s, Id: %s", g.Game.Title, g.Id)
}

func NewGameObject() *GameObject {

	obj := &GameObject{
		Scale: Scale{
			ScaleX: 1,
			ScaleY: 1,
		},
		Shape: Shape{
			Visible: true,
		},
	}

	obj.OnUpdate = func() {
		// obj.X += obj.VelX
		// obj.Y += obj.VelY
	}

	obj.OnDraw = func() {}
	obj.OnInit = func() {}

	return obj
}

func (obj *GameObject) GetScene() *GameObject {
	if obj.Game == nil {
		fmt.Println("Error: Game is nil")
		return nil
	}
	return obj.GetGame().GetCurrentScene()
}