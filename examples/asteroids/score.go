package main

import "strconv"

func IncreaseScore(points int) {
	score += points
	scoreString = strconv.Itoa(score)
}