package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// SCENE GRAPH CURRENTLY WORKS! DON'T BREAK IT!

func (g *Game) Draw() {

	scene := g.GetCurrentScene()
	scene.OnDraw()

	for _, obj := range scene.GetChildren() {
		drawWithTransforms(obj)
    }
}


func drawWithTransforms(obj *GameObject) {

	offsetX := obj.GetX() + (obj.GetWidth() * obj.GetOriginX())
	offsetY := obj.GetY() + (obj.GetHeight() * obj.GetOriginY())

	rl.PushMatrix()
	
		rl.Translatef(offsetX, offsetY, 0)
		rl.Rotatef(obj.GetAngle(), 0, 0, 1)
		rl.Scalef(obj.GetScaleX(), obj.GetScaleY(), 1)
		rl.Translatef(-offsetX, -offsetY, 0)

		obj.OnDraw()

		if len(obj.GetChildren()) > 0 {

			rl.Translatef(obj.GetX(), obj.GetY(), 0)
			
			for _, c := range obj.GetChildren() {
				drawWithTransforms(c)
			}
		}

	rl.PopMatrix()


}