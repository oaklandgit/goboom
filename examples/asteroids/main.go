package main

import (
	boom "goboom/goboom"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"golang.org/x/exp/rand"
)

const (
	ASTEROIDS_COUNT = 12
)

func main() {

	rand.Seed(uint64(time.Now().UnixNano()))
	game := boom.NewGame(800, 600, "Asteroids")
	game.SetBgColor(rl.Black)

	scene := game.GetCurrentScene()

	// ship := createPlayer(scene)

	for i := 0; i < ASTEROIDS_COUNT ; i++ {
		x := rand.Intn(int(game.GetWidth()))
		y := rand.Intn(int(game.GetWidth()))
		angle := float32(rand.Intn(360))

        asteroid := createAsteroid(float32(x), float32(y))
		asteroid.SetVelocityByHeading(angle, 1)
        scene.Add(asteroid)
    }

	// ship.SetXY(game.GetWidth()/2, game.GetHeight()/2)
	// scene.Add(ship)

	game.Run()

}