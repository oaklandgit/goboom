package main

import (
	"fmt"
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var score int = 0

func createScoreboard() (*boom.GameObject, func(int)) {

	txt := boom.NewTextComp(fmt.Sprintf("Score: %d", score), 30, rl.White)
	scoreboard := boom.NewGameObject()
	scoreboard.AddComponent(txt)
	scoreboard.SetId("scoreboard")

	incrementScore := func(s int) {
		score += s
		txt.SetText(fmt.Sprintf("Score: %d", score))
	}

	return scoreboard, incrementScore

}