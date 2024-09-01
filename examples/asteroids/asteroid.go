package main

import (
	boom "goboom/goboom"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func createAsteroid(x, y float32) boom.Renderable {

    points := []float32{}
    sides := rl.GetRandomValue(5, 9)
    size := rl.GetRandomValue(20, 50)
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
    return asteroid
}