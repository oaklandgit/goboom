package main

import (
	boom "goboom/goboom"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"golang.org/x/exp/rand"
)

const (
	LEFTMOST_STARTX = 40
	LEFTMOST_ANGLE = 160
	RIGHTMOST_ANGLE = 200
	MIN_SPEED = 0.5
	MAX_SPEED = 1.5
)

func spawnMissiles(scene *boom.GameObject) {

	totalMissiles := 10

	colors := []rl.Color{
		rl.Red,
		rl.Orange,
		rl.Pink,
	}

	for i := 0; i < totalMissiles; i++ {
		time.Sleep(2 * time.Second)
		index := rand.Intn(len(colors))
		speed := boom.RandFloatBetween(MIN_SPEED, MAX_SPEED)
		angle := boom.RandIntBetween(LEFTMOST_ANGLE, RIGHTMOST_ANGLE) // towards the base. (180 is straight down)
		posX := boom.RandIntBetween(LEFTMOST_STARTX, int(scene.GetGame().GetWidth() - LEFTMOST_STARTX))
		createMissile(scene, float32(posX), angle, float32(speed), colors[index])
	}
}

func createMissile(scene *boom.GameObject, posX float32, angle int, speed float32, color rl.Color) {

	missile := boom.Circle(posX, 0, 5, color, rl.Blank, 2)
	missile.SetOrigin(-0.5, -0.5)
	vel := boom.NewVelocityComp(0, 0)
	vel.SetVelocityByHeading(float32(angle), speed)
	missile.AddComponent(vel)

	trail := boom.Line(posX, 0, missile.GetGlobalX(), missile.GetGlobalY(), color, 2)
	trail.SetSize(0, 0)

	trailComp := trail.GetComponent("line").(*boom.LineComp)
	
	trail.OnUpdate = func() {
		trailComp.End.X = missile.GetGlobalX()
		trailComp.End.Y = missile.GetGlobalY()
	}

	scene.Add(missile, trail)
}