package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const ROTATE_SPEED = 10
const THRUST_FORCE = 0.1
const BULLET_FORCE = 10

func createPlayer() *boom.GameObject {

	ship := boom.Polygon(0, 0, "0 20 10 0 20 20", false, rl.Blue, rl.White, 2)
	ship.SetOrigin(0.5, 0.5)
	

	// MOVEMENT
	velocity := boom.NewVelocityComp(0, 0)
	control := boom.NewInputComp()
	wrap := boom.NewWrapComp(true)

	control.NewInput(rl.KeyRight, boom.KeyDown, func() {
		ship.AddAngle(ROTATE_SPEED)
	})

	control.NewInput(rl.KeyLeft, boom.KeyDown, func() {
		ship.AddAngle(-ROTATE_SPEED)
	})

	control.NewInput(rl.KeyUp, boom.KeyDown, func() {
		velocity.AddVelocityByHeading(ship.GetAngle(), THRUST_FORCE)
		
	})

	ship.AddComponents(velocity, control, wrap)

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
