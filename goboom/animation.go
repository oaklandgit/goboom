package goboom

import rl "github.com/gen2brain/raylib-go/raylib"

type Animation struct {
	Frames []rl.Texture2D
	CurrentFrame int
}