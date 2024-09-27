package main

import (
	"fmt"
	boom "goboom/goboom"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"golang.org/x/exp/rand"
)

const (
	ASTEROIDS_COUNT = 12
)

var score = 0

func main() {

	rand.Seed(uint64(time.Now().UnixNano()))
	game := boom.NewGame(800, 600, "Asteroids")
	game.SetBgColor(rl.Black)

	scene := game.GetCurrentScene()
	scoreboard := boom.Text(0, 0, fmt.Sprintf("Score: %d", score), 30, rl.White)
	scoreboard.SetId("scoreboard")
	
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
	boom.PutTopCenter(scene, scoreboard, 0, 12)

	game.Run()

}