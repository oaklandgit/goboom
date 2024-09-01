package main

import (
	boom "goboom/goboom"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func createPlayer(game *boom.Game) boom.Renderable {

	ship := boom.NewPolygon(
		0, 0, rl.White, false, 0, 0, 20, 10, 0, 20)
	ship.SetStroke(rl.White, 2)
	ship.SetOrigin(0.5, 0.5)

	ship.AddInput(rl.KeyRight, boom.KeyDown, func() {
		ship.SetAngle(ship.GetAngle() + ROTATE_SPEED)
	})

	ship.AddInput(rl.KeyLeft, boom.KeyDown, func() {
		ship.SetAngle(ship.GetAngle() - ROTATE_SPEED)
	})

	ship.AddInput(rl.KeyUp, boom.KeyDown, func() {
		ship.SetVelocityByHeading(ship.GetAngle(), 2)
	})

	// always spawn the ship invincible for 3 seconds
	cancelInvincibility := boom.PowerUp(ship, Invincible, 10 * time.Second)

	// cancel the invincibility if user fires
	ship.AddInput(rl.KeyX, boom.KeyPressed, func() {
		bullet := createBullet(ship.GetX(), ship.GetY(), ship.GetAngle())
		
		if cancelInvincibility != nil {
			cancelInvincibility()
			cancelInvincibility = nil
		}
		game.Add(bullet)
	})


	return ship

}

func Invincible(r boom.Renderable) func(boom.Renderable) {
	r.SetStroke(rl.DarkGray, 2)
	r.AddTags("invincible")
	return func(r boom.Renderable) {
		r.SetStroke(rl.White, 2)
		r.RemoveTag("invincible")
	}
}
