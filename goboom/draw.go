package goboom

import rl "github.com/gen2brain/raylib-go/raylib"

type Renderable interface {
	OnInit()
	OnUpdate()
	OnDraw()

	GetX() float32
	GetY() float32
	GetWidth() float32
	GetHeight() float32
	GetScaleX() float32
	GetScaleY() float32
	GetAngle() float32
	GetOriginX() float32
	GetOriginY() float32

	IsVisible() bool
	IsDeleted() bool
}

func (g *Game) Draw() {

	for _, obj := range g.GameObjects {
		
		if obj.IsVisible() {

			offsetX := obj.GetX() + obj.GetWidth() * obj.GetOriginX()
			offsetY := obj.GetY() + obj.GetHeight() * obj.GetOriginY()

			rl.PushMatrix()
			rl.Translatef(offsetX, offsetY, 0)

			rl.Rotatef(obj.GetAngle(), 0, 0, 1)
			rl.Scalef(obj.GetScaleX(), obj.GetScaleY(), 1)
			
			rl.Translatef(-offsetX, -offsetY, 0)

			obj.OnDraw()

			rl.PopMatrix()
	
		}

	}

}