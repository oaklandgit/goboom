package goboom

import rl "github.com/gen2brain/raylib-go/raylib"

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

			obj.Draw()

			rl.PopMatrix()
	
		}

	}

}