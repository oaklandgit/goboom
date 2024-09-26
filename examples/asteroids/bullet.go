package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func createBullet(scene *boom.GameObject, x, y, angle, thrust float32) *boom.GameObject {
	bullet := boom.Rectangle(x, y, 1, 20, rl.Red, rl.Blank, 0)
	bullet.AddTags("bullet")
	bullet.SetOrigin(0, 0)
	vel := boom.NewVelocityComp(0, 0)
	vel.SetVelocityByHeading(angle, thrust)

	bullet.AddComponents(vel)
	bullet.SetAngle(angle)
	bullet.SetLifespan(1_000) // 1 second

	collide := boom.NewCollideComp(boom.CollisionCircle{Radius: 5})
	collide.NewCollider("asteroid", func(self, asteroid *boom.GameObject) {
		boom.NewExplosion(
			scene,
			asteroid.GetX(), asteroid.GetY(),
			4, 6,
			5, 12,
			func() *boom.GameObject {
				return boom.Rectangle(0, 0, 10, 10, rl.White, rl.Blank, 0)
			})
		asteroid.SetLifespan(0)
	})


	bullet.AddComponents(collide)


	// bullet.AddCollider("asteroid", func(self, asteroid *boom.GameObject) {
    //     fmt.Println("Bullet collided with asteroid!")
	// 	bullet.GetScene().NewExplosion(
	// 	 	self.CenterX(), self.CenterY(),
	// 		4, 6,
	// 		5, 12,
	// 		func() *boom.GameObject {
	// 			return boom.NewRectangle(0, 0, 10, 10, rl.White)
	// 		})
	// 	self.SetLifespan(0)
	// 	asteroid.SetLifespan(0)
	// 	IncreaseScore(10)

    // })

	return bullet
}