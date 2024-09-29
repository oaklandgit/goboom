package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func createCrosshairs(scene *boom.GameObject, incrementScore func(int)) *boom.GameObject {

	// A blank game object to hold the crosshairs
	crosshairs := boom.NewGameObject()
	mouse := boom.NewMouseComp()
	mouse.SetCursor(rl.MouseCursorCrosshair)
	mouse.OnMove(func(x, y float32) {
		crosshairs.SetXY(x, y)
	})
	mouse.OnClick(func() {
		defend := createBurst(scene, crosshairs.GetX(), crosshairs.GetY(), rl.White)
		defend.AddTags("defend")
		collide := boom.NewCollideComp(boom.CollisionCircle{Radius: 20})
		collide.NewCollider("missile", func(d, m *boom.GameObject) {
			incrementScore(10)
			m.SetLifespan(0)
		})
		defend.AddComponent(collide)
	})

	crosshairs.AddComponent(mouse)

	return crosshairs

}