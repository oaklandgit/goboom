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

func createAsteroid(x, y float32) *boom.GameObject {

    points := ""
    sides := rl.GetRandomValue(MIN_SIDES, MAX_SIDES)
    size := rl.GetRandomValue(MIN_SIZE, MAX_SIZE)
    angleIncrement := 2 * math.Pi / float64(sides)

    for i := 0; i < int(sides); i++ {
        r := rl.GetRandomValue(int32(size/2), int32(size))
        angle := float64(i) * angleIncrement

        points += fmt.Sprintf("%.1f %.1f ", float32(r)*float32(math.Cos(angle)), float32(r)*float32(math.Sin(angle)))
    }

    // Ensure the polygon is closed by adding the first point at the end
    // points = append(points, points[0], points[1])

    fmt.Println(points)

    asteroid := boom.NewPolygon(x, y, rl.White, true, points)
    asteroid.SetWrap(true)
    asteroid.AddTags("asteroid")
    asteroid.SetPadding(asteroid.GetWidth(), asteroid.GetHeight())

    asteroid.AddCollider("player", func(a, b *boom.GameObject) {
        fmt.Println("Asteroid collided with player!")
    })

    return asteroid
}