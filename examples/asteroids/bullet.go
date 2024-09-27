package main

import (
	"fmt"
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func incrementScore(scene *boom.GameObject, points int) {
	scoreboard := scene.GetById("scoreboard")
	score += points
	scoreboard.GetComponent("text").(*boom.TextComp).SetText(fmt.Sprintf("Score: %d", score))
}

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
		boom.Splosion(
			scene,
			asteroid.GetX(), asteroid.GetY(),
			6, 12,
			1, 3,
			func() *boom.GameObject {
				return boom.Rectangle(0, 0, 7, 5, rl.Blank, rl.White, 1)
			})
		incrementScore(scene, 10)
		asteroid.SetLifespan(0)
	})


	bullet.AddComponents(collide)

	return bullet
}