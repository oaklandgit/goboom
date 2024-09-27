package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func spawnMissiles(scene *boom.GameObject) {
	createMissile(scene, rl.NewVector2(100, 0), rl.NewVector2(1, 2))
}

func createMissile(scene *boom.GameObject, start rl.Vector2, v rl.Vector2) {

	startX := start.X
	startY := start.Y

	missile := boom.Circle(start.X, start.Y, 5, rl.Red, rl.Blank, 2)
	missile.SetOrigin(-0.5, -0.5)
	vel := boom.NewVelocityComp(v.X, v.Y)
	missile.AddComponent(vel)

	trail := boom.Line(startX, startY, missile.GetGlobalX(), missile.GetGlobalY(), rl.Red, 2)

	trailComp := trail.GetComponent("line").(*boom.LineComp)
	

	trail.OnUpdate = func() {
		trailComp.End.X = missile.GetGlobalX()
		trailComp.End.Y = missile.GetGlobalY()
	}

	scene.Add(missile, trail)
}