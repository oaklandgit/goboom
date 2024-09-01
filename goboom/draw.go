package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) Draw() {

	scene := g.GetCurrentScene()
	scene.Draw()

	for _, obj := range scene.GetChildren(){
        if obj.IsVisible() {
            drawWithTransforms(obj)
        }

    }
}



func drawWithTransforms(obj Renderable) {


	offsetX := obj.GetX() + obj.GetWidth() / 2
	offsetY := obj.GetY() + obj.GetHeight() / 2

	rl.PushMatrix()
	rl.Translatef(offsetX, offsetY, 0)
	rl.Rotatef(obj.GetAngle(), 0, 0, 1)
	rl.Scalef(obj.GetScaleX(), obj.GetScaleY(), 1)
	rl.Translatef(-offsetX, -offsetY, 0)
	
	obj.Draw()

	if len(obj.GetChildren()) > 0 {

		rl.Translatef(-obj.GetWidth()/2, -obj.GetHeight()/2, 0)
	
		for _, c := range obj.GetChildren() {
			drawWithTransforms(c)
		}
	}

	rl.PopMatrix()


}