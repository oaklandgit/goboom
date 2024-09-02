package main

import (
	boom "goboom/goboom"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const ROTATE_SPEED = 10

func createPlayer(scene boom.Renderable) boom.Renderable {

	ship := boom.NewGroup(0, 0)
	ship.SetOrigin(0.9, 0.5)
	ship.SetId("player")
	ship.SetWrap(true)

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
		scene.Add(bullet)
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
