package main

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
}

func (g *Game) Draw() {

	for _, obj := range g.GameObjects {
		
		if obj.IsVisible() {

			rl.PushMatrix()

			rl.Translatef(
				obj.GetX() + obj.GetWidth() * obj.GetOriginX(),
				obj.GetY() + obj.GetHeight() * obj.GetOriginY(), 0)

			rl.Rotatef(obj.GetAngle(), 0, 0, 1)
			rl.Scalef(obj.GetScaleX(), obj.GetScaleY(), 1)

			rl.Translatef(
				-obj.GetX() + obj.GetWidth() * obj.GetOriginX(),
				-obj.GetY() + obj.GetHeight() * obj.GetOriginY(), 0)


			obj.OnDraw()

			rl.PopMatrix()
	
		}

	}

}