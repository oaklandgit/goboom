package main

import (
	"fmt"
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)


var score int = 0

func main() {

	game := boom.NewGame(800, 600, "Mousle Command")
	game.SetBgColor(rl.Black)
	scene := game.GetCurrentScene()

	// Earth
	floor := scene.GetGame().Height - 100
	width := scene.GetGame().Width
	earth := boom.Rectangle(0, floor, width, 100, rl.Blue, rl.Black, 1)
	collide := boom.NewCollideComp(boom.CollisionRect{Width: width, Height: 100})
	earth.AddComponent(collide)
	earth.AddTags("earth")

	// Bases
	bases := boom.GridArray(1, 4, 120, NewBase)
	boom.PutCenter(scene, bases, -30, 180)

	// Scoreboard
	scoreboard := boom.Text(0, 0, fmt.Sprintf("Score: %d", score), 30, rl.White)
	scoreboard.SetId("scoreboard")

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
		defend.AddComponent(collide)
		collide.NewCollider("missile", func(d, m *boom.GameObject) {
			// score += 10
		})
	})

	crosshairs.AddComponent(mouse)
	scene.Add(scoreboard, crosshairs, earth, bases)
	boom.PutCenter(scene, crosshairs, 0, 0)
	boom.PutTopCenter(scene, scoreboard, 0, 32)

	go spawnMissiles(scene)

	game.Run()

}



