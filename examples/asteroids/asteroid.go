package main

import (
	"fmt"
	boom "goboom/goboom"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)
const (
    MIN_SIDES = 5
    MAX_SIDES = 9
    MIN_SIZE = 20
    MAX_SIZE = 80
)

func createAsteroid(scene *boom.GameObject, x, y float32) *boom.GameObject {

    points := ""
    sides := rl.GetRandomValue(MIN_SIDES, MAX_SIDES)
    size := rl.GetRandomValue(MIN_SIZE, MAX_SIZE)
    angleIncrement := 2 * math.Pi / float64(sides)

    for i := 0; i < int(sides); i++ {
        r := rl.GetRandomValue(int32(size/2), int32(size))
        angle := float64(i) * angleIncrement

        points += fmt.Sprintf("%.1f %.1f ", float32(r)*float32(math.Cos(angle)), float32(r)*float32(math.Sin(angle)))
    }

    fmt.Println(points)

    asteroid := boom.Polygon(x, y, points, true, rl.Blank, rl.White, 2)
    asteroid.SetSize(float32(size * 2), float32(size * 2))

    wrap := boom.NewWrapComp(true, true)
    wrap.SetPadding(asteroid.GetWidth(), asteroid.GetHeight())
    collide := boom.NewCollideComp(boom.CollisionCircle{Radius: float32(size)})
    collide.NewCollider("player", func(a, p *boom.GameObject) {
        fmt.Println("Asteroid collided with player!")
        boom.Splosion(
            scene,
            p.GetGlobalX(), p.GetGlobalY(),
            5, 10,
            2, 4,
            func() *boom.GameObject {
                return boom.Rectangle(0, 0, 3, 3, rl.Blank, rl.White, 2)
            })
    })
    spin := boom.NewRotVelocityComp(float32(rl.GetRandomValue(-3, 3)) * 0.1)

    asteroid.AddComponents(collide, wrap, spin)

    asteroid.AddTags("asteroid")

    // asteroid.AddCollider("player", func(a, b *boom.GameObject) {
    //     fmt.Println("Asteroid collided with player!")
    // })

    return asteroid
}