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
	BALL_VELX = -1
	BALL_VELY = 2
)

func main() {

	rand.Seed(uint64(time.Now().UnixNano()))
	game := boom.NewGame(800, 600, "Brickout")
	game.SetBgColor(rl.DarkBlue)

	scene := game.GetCurrentScene()

	// Paddle
	paddle := createPaddle()
	boom.PutBottom(scene, paddle, 0, -60)

	// Ball
	ball := createBall(BALL_VELX, BALL_VELY)
	boom.PutCenter(scene, ball, 0, 0)

	// One Brick
	redBrick := func() *boom.GameObject {
		return createBrick(rl.Red)
	}

	// Bricks
	bricks := boom.GridArray(3, 14, 6, redBrick)
	bricks.SetId("the bricks")

	fmt.Println("Bricks", bricks.GetBoundingBox().Width, bricks.GetBoundingBox().Height)
	boom.PutCenter(scene, bricks, 0, 100)
	
	
	// Walls
	leftWall := createWall(0, 0, WALL_WIDTH, scene.GetHeight())
	rightWall := createWall(scene.GetWidth() - WALL_WIDTH, 0, WALL_WIDTH, scene.GetHeight())
	ceiling := createWall(0, 0, scene.GetWidth(), WALL_WIDTH)
	
	scene.Add(paddle, ball, bricks, ceiling, leftWall, rightWall)
	
	// scene.Print()
	// fmt.Println(scene.GetWidth(), scene.GetHeight())
	
	game.Run()

}
