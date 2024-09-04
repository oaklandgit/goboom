package main

import (
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

    points := []float32{}
    sides := rl.GetRandomValue(MIN_SIDES, MAX_SIDES)
    size := rl.GetRandomValue(MIN_SIZE, MAX_SIZE)
    angleIncrement := 2 * math.Pi / float64(sides)

    for i := 0; i < int(sides); i++ {
        r := rl.GetRandomValue(int32(size/2), int32(size))
        angle := float64(i) * angleIncrement
        points = append(points, float32(r)*float32(math.Cos(angle)))
        points = append(points, float32(r)*float32(math.Sin(angle)))
    }

    // Ensure the polygon is closed by adding the first point at the end
    points = append(points, points[0], points[1])

    asteroid := boom.NewPolygon(x, y, rl.White, true, points...)
    asteroid.SetWrap(true)
    asteroid.SetPadding(asteroid.GetWidth(), asteroid.GetHeight())
    return asteroid
}