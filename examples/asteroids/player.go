package main

import (
	boom "goboom/goboom"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const ROTATE_SPEED = 10
const THRUST_FORCE = 2
const BULLET_FORCE = 10

func createPlayer(scene *boom.GameObject) *boom.GameObject {

	ship := boom.NewGroup(0, 0)
	ship.SetId("player")
	ship.AddTags("player")
	ship.SetWrap(true)

	// WINDOW
	window := boom.NewCircle(0, 0, 3, rl.Blank)
	window.SetFill(rl.White)
	window.SetOrigin(0.5, 0.5)
	boom.PutLeft(ship, window, -3, -4)
	ship.Add(window)

	// BODY
	body := boom.NewEllipse(0, 0, 7, 12, rl.Blue)
	body.SetStrokeWeight(2)
	body.SetOrigin(0.5, 0.2)

	ship.Add(body)

	ship.AddInput(rl.KeyRight, boom.KeyDown, func() {
		ship.AddAngle(ROTATE_SPEED)
	})

	ship.AddInput(rl.KeyLeft, boom.KeyDown, func() {
		ship.AddAngle(- ROTATE_SPEED)
	})

	ship.AddInput(rl.KeyUp, boom.KeyDown, func() {
		ship.SetVelocityByHeading(ship.GetAngle(), THRUST_FORCE)
	})

	// always spawn the ship invincible for 3 seconds
	cancelInvincibility := boom.PowerUp(ship, Invincible, 10 * time.Second)

	// cancel the invincibility if user fires
	ship.AddInput(rl.KeyX, boom.KeyPressed, func() {
		bullet := createBullet(
			window.GetGlobalX(),
			window.GetGlobalY(),
			window.GetGlobalAngle(),
			BULLET_FORCE)
		
		if cancelInvincibility != nil {
			cancelInvincibility()
			cancelInvincibility = nil
		}
		scene.Add(bullet)
	})


	return ship

}

func Invincible(r *boom.GameObject) func(*boom.GameObject) {
	r.SetStroke(rl.DarkGray, 2)
	r.AddTags("invincible")
	return func(r *boom.GameObject) {
		r.SetStroke(rl.White, 2)
		r.RemoveTag("invincible")
	}
}
