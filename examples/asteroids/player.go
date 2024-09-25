package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const ROTATE_SPEED = 10
const THRUST_FORCE = 0.1
const BULLET_FORCE = 10

func createPlayer() *boom.GameObject {

	ship := boom.NewGroup(0, 0)
	ship.SetId("player")
	ship.AddTags("player")
	ship.SetWrap(true)

	// WINDOW
	window := boom.Ellipse(0, 0, 3, 3, rl.Blank, rl.White, 2)
	window.SetFill(rl.White)
	window.SetOrigin(0.5, 0.5)
	boom.PutLeft(ship, window, -3, -4)
	ship.Add(window)

	// BODY
	body := boom.Ellipse(0, 0, 7, 12, rl.Blue, rl.White, 2)
	body.SetOrigin(0.5, 0.2)
	ship.Add(body)

	// MOVEMENT
	velocity := boom.NewVelocityComp(0, 0)
	control := boom.NewInputComp()

	control.NewInput(rl.KeyRight, boom.KeyDown, func() {
		ship.AddAngle(ROTATE_SPEED)
	})

	control.NewInput(rl.KeyLeft, boom.KeyDown, func() {
		ship.AddAngle(-ROTATE_SPEED)
	})

	control.NewInput(rl.KeyUp, boom.KeyDown, func() {
		velocity.AddVelocityByHeading(ship.GetAngle(), THRUST_FORCE)
		
	})

	ship.AddComponents(velocity, control)

	// always spawn the ship invincible for 3 seconds
	// cancelInvincibility := boom.PowerUp(ship, Invincible, 10 * time.Second)

	// cancel the invincibility if user fires
	// ship.AddInput(rl.KeyX, boom.KeyPressed, func() {
	// 	bullet := createBullet(
	// 		window.GetGlobalX(),
	// 		window.GetGlobalY(),
	// 		window.GetGlobalAngle(),
	// 		BULLET_FORCE)
		
	// 	if cancelInvincibility != nil {
	// 		cancelInvincibility()
	// 		cancelInvincibility = nil
	// 	}
	// 	scene.Add(bullet)
	// })


	return ship

}

// func Invincible(r *boom.GameObject) func(*boom.GameObject) {
// 	r.SetStroke(rl.DarkGray, 2)
// 	r.AddTags("invincible")
// 	return func(r *boom.GameObject) {
// 		r.SetStroke(rl.White, 2)
// 		r.RemoveTag("invincible")
// 	}
// }
