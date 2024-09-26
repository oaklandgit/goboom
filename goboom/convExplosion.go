package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewExplosion(
	scene *GameObject,
	x, y float32,
	minCount, maxCount int32,
	minForce, maxForce int32,
	particle func() *GameObject) {

	count := rl.GetRandomValue(minCount, maxCount)

	for i := 0; i < int(count); i++ {
		a := float32(rl.GetRandomValue(0, 360))
		d := rl.GetRandomValue(100, 1_000)
		strength := float32(rl.GetRandomValue(minForce, maxForce))
		p := particle()
		p.SetXY(x, y)
		p.SetLifespan(int(d))
		p.SetAngle(a)

		vel := NewVelocityComp(0, 0)
		vel.SetVelocityByHeading(a, strength)

		p.AddComponent(vel)
		scene.Add(p)
	}

	

}
