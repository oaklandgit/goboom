package main

import (
	boom "goboom/goboom"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"golang.org/x/exp/rand"
)

const ROTATE_SPEED = 10


func main() {

	rand.Seed(uint64(time.Now().UnixNano()))
	game := boom.NewGame(800, 600, "Asteroids")
	game.SetBgColor(rl.Black)

	ship := createPlayer(game)

	for i := 0; i < 5; i++ {
		x := rand.Intn(int(game.GetWidth()))
		y := rand.Intn(int(game.GetWidth()))
		angle := float32(rand.Intn(360))

        asteroid := createAsteroid(float32(x), float32(y))
		asteroid.SetVelocityByHeading(angle, 1)
        game.Add(asteroid)
    }


	game.PutCenter(ship, 0, 0)
	game.Add(ship)



	game.Run()

}