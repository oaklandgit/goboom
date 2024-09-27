package goboom

import "golang.org/x/exp/rand"

func RandIntBetween(min, max int) int {
    return rand.Intn(max-min+1) + min // Generate a random number between min and max
}

func RandFloatBetween(min, max float32) float32 {
    return min + rand.Float32()*(max-min) // Generate a random float32 between min and max
}