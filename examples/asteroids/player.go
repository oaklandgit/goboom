package main

import (
	"fmt"
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func createPlayer(game *boom.Game) boom.Renderable {

	ship := boom.NewPolygon(
		0, 0, rl.White, false, 0, 0, 20, 10, 0, 20)
	ship.SetStroke(rl.White, 2)
	ship.SetOrigin(0.5, 0.5)

	ship.NewInput(rl.KeyRight, boom.KeyDown, func() {
		ship.SetAngle(ship.GetAngle() + ROTATE_SPEED)
	})

	ship.NewInput(rl.KeyLeft, boom.KeyDown, func() {
		ship.SetAngle(ship.GetAngle() - ROTATE_SPEED)
	})

	ship.NewInput(rl.KeyUp, boom.KeyDown, func() {
		ship.SetVelocityByHeading(ship.GetAngle(), 2)
	})

	ship.NewInput(rl.KeyX, boom.KeyPressed, func() {
		fmt.Println("FIRE!")
		bullet := boom.NewCircle(ship.GetX(), ship.GetY(), 3, rl.White)
		bullet.SetFill(rl.White)
		bullet.SetOrigin(0.5, 0.5)
		bullet.SetVelocityByHeading(ship.GetAngle(), 5)
		bullet.SetLifespan(1_000) // 1 second
		game.Add(bullet)
	})

	return ship

}