package main

import (
	boom "goboom/goboom"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func createPlayer(scene boom.Renderable) boom.Renderable {

	ship := boom.NewGroup(0, 0)
	ship.SetOrigin(0, 0.5)

	shape := boom.NewPolygon(0, 0, rl.White, true,
		29, 7.65, 27.45, 22.4, 23.76, 22.01, 24.16, 18.21, 21.76, 18.21, 19.87, 21.26, 9.4, 21.26, 7.52, 18.21, 4.84, 18.21, 5.24, 22.01, 1.55, 22.4, 0, 7.65, 3.69, 7.26, 4.55, 15.43, 7.54, 15.43, 10.95, 0, 18.05, 0, 21.46, 15.43, 24.46, 15.43, 25.32, 7.26, 29, 7.65)
	shape.SetAngle(90)

	ship.Add(shape)
	
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
