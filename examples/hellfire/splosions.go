package main

import (
	boom "goboom/goboom"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const GROWTH_RATE = 3
const FADE_RATE = 5


func createBurst(scene *boom.GameObject, x, y float32, color rl.Color) *boom.GameObject {

	var alpha float32 = 255

	burst := boom.Circle(x, y, 2, color, rl.Blank, 2)
	burst.SetOrigin(-0.5, -0.5)
	shape := burst.GetComponent("ellipse").(*boom.EllipseComp)

	expand := boom.NewTweenComp()
	expand.ChangeFunc = func() {
		burst.SetWidth(burst.GetWidth() + GROWTH_RATE)
		burst.SetHeight(burst.GetHeight() + GROWTH_RATE)
		
		alpha = float32(math.Max(0, float64(alpha - FADE_RATE)))
		shape.FillColor.A = uint8(alpha)

	}
	expand.EndCondition = func() bool {
		return alpha <= 0
	}

	expand.AfterFunc = func() {
		burst.SetLifespan(0)
	}

	burst.AddComponent(expand)
	scene.Add(burst)

	return burst

}