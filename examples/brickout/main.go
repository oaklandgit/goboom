package main

import (
	"fmt"
	boom "goboom/goboom"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"golang.org/x/exp/rand"
)

const (
	PADDLE_SPEED = 10
	WALL_WIDTH = 22
	BALL_VELX = -2
	BALL_VELY = 3
)

var score int

func main() {

	rand.Seed(uint64(time.Now().UnixNano()))
	game := boom.NewGame(800, 600, "Brickout")
	game.SetBgColor(rl.DarkBlue)

	scene := game.GetCurrentScene()

	// Paddle
	paddle := createPaddle()
	boom.PutBottom(scene, paddle, 0, -60)

	// Ball
	ball := createBall(scene, BALL_VELX, BALL_VELY)
	boom.PutCenter(scene, ball, 0, 0)

	// One Brick
	makeRedBrick := func() *boom.GameObject {
		return createBrick(rl.Red)
	}

	// Bricks
	bricks := boom.GridArray(3, 14, 4, makeRedBrick)
	// bricks.SetXY(20, 20)
	boom.PutCenter(scene, bricks, 0, -140)
	bricks.SetId("the bricks")

	// Scoreboard
	scoreboard := boom.Text(0, 0, fmt.Sprintf("Score: %d", score), 30, rl.White)
	scoreboard.SetId("scoreboard")
	
	// Walls
	leftWall := createWall(0, 0, WALL_WIDTH, scene.GetHeight())
	leftWall.AddTags("wall")
	rightWall := createWall(scene.GetWidth() - WALL_WIDTH, 0, WALL_WIDTH, scene.GetHeight())
	rightWall.AddTags("wall")
	ceiling := createWall(0, 0, scene.GetWidth(), WALL_WIDTH)
	ceiling.AddTags("ceiling")
	
	scene.Add(paddle, ball, bricks, ceiling, leftWall, rightWall, scoreboard)

	boom.PutTopCenter(scene, scoreboard, 0, 32)
	
	// scene.Print()
	// fmt.Println(scene.GetWidth(), scene.GetHeight())
	
	game.Run()

}
