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

// var score = 0
// var scoreString = "0"

func main() {

	rand.Seed(uint64(time.Now().UnixNano()))
	game := boom.NewGame(800, 600, "Asteroids")
	game.SetBgColor(rl.Black)

	scene := game.GetCurrentScene()
	scoreboard := boom.Text(0, 0, "Score: 0", 30, rl.White)
	
	ship := createPlayer(scene)
	
	for i := 0; i < ASTEROIDS_COUNT ; i++ {
		x := rand.Intn(int(game.GetWidth()))
		y := rand.Intn(int(game.GetWidth()))
		angle := float32(rand.Intn(360))
		
        asteroid := createAsteroid(float32(x), float32(y))
		
		movement := boom.NewVelocityComp(0, 0)
		movement.SetVelocityByHeading(angle, 1)
		
		asteroid.AddComponents(movement)
        scene.Add(asteroid)
    }
	
	scene.Add(ship, scoreboard)
	boom.PutCenter(scene, ship, 0, 0)
	boom.PutTop(scene, scoreboard, 0, 0)

	game.Run()

}