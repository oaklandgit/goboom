package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)


type Entity interface {
	Init()
	Update()
	Draw()
	GetParent() Entity
	GetChildren() []Entity
	Add(...Entity)
	Remove(Entity)
	IsDeleted() bool
}

type Game struct {
	Title string
	Width float32
	Height float32
	FPS int
	BgColor rl.Color
	Scenes []Entity
	CurrentScene int
}


func NewGame(width, height float32, title string) *Game {

	g := &Game{
		Title: title,
		Width: width,
		Height: height,
		BgColor: rl.White,
		FPS: 60,
		Scenes: []Entity{},
		CurrentScene: 0,
	}

	g.NewScene("default")
	return g
}

func (g *Game) SetBgColor(color rl.Color) {
	g.BgColor = color
}

func (g *Game) GetTitle() string {
	return g.Title
}

func (g *Game) GetWidth() float32 {
	return g.Width
}

func (g *Game) GetHeight() float32 {
	return g.Height
}

// put into the scene's own Update method

// func (scene *GameObject) Cleanup() {
// 	for _, c := range scene.GetAll() {
// 		if c.IsDeleted() {
// 			c.GetParent().Remove(c)
// 		}
// 	}
// }