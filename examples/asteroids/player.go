package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const ROTATE_SPEED = 10
const THRUST_FORCE = 0.1
const BULLET_FORCE = 10

func createPlayer(scene *boom.GameObject) *boom.GameObject {

	ship := boom.Polygon(0, 0, "-10 10 0 -10 10 10", false, rl.Blue, rl.White, 2)
	ship.SetSize(18, 18)
	
	bulletOrigin := boom.Ellipse(-5, -12, 10, 10, rl.Blank, rl.White, 2)
	ship.Add(bulletOrigin)
	
	// MOVEMENT
	velocity := boom.NewVelocityComp(0, 0)
	velocity.ApplyDrag(0.99)
	control := boom.NewInputComp()
	wrap := boom.NewWrapComp(true)
	wrap.SetPadding(ship.GetWidth(), ship.GetHeight())

	control.NewInput(rl.KeyRight, boom.KeyDown, func() {
		ship.AddAngle(ROTATE_SPEED)
	})

	control.NewInput(rl.KeyLeft, boom.KeyDown, func() {
		ship.AddAngle(-ROTATE_SPEED)
	})

	control.NewInput(rl.KeyUp, boom.KeyDown, func() {
		velocity.AddVelocityByHeading(ship.GetAngle(), THRUST_FORCE)
	})

	control.NewInput(rl.KeySpace, boom.KeyPressed, func() {
		bullet := createBullet(
			bulletOrigin.GetGlobalX(),
			bulletOrigin.GetGlobalY(),
			bulletOrigin.GetGlobalAngle(),
			BULLET_FORCE)
		scene.Add(bullet)
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
