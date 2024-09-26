package goboom

import (
	"fmt"
)

func (g *Game) NewScene(id string) *GameObject {
	
	scene := &GameObject{
		Scale: Scale{
			ScaleX: 1,
			ScaleY: 1,
		},
		Shape: Shape{
			Visible: true,
			Width: g.Width,
			Height: g.Height,
		},
		OnUpdate: func() {},
		OnDraw: func() {},
	}

	scene.SetGame(g)
	scene.SetId(id)
	g.Scenes = append(g.Scenes, scene)	
	
	return scene
}

func (g *Game) GetCurrentScene() *GameObject {
	if g.Scenes[g.CurrentScene] == nil {
		fmt.Println("Error: CurrentScene is nil")
		return nil
	}
	return g.Scenes[g.CurrentScene]
}

func (g *Game) SetScene(index int) {
    if index >= 0 && index < len(g.Scenes) {
        g.CurrentScene = index
    }
}