package goboom

type Scene struct {
	Game *Game
	Children []Entity
}

func (g *Game) NewScene(id string) Entity {
	
	scene := &Scene{
		Game: g,
	}

	// scene.SetId(id)
	g.Scenes = append(g.Scenes, scene)
	
	return scene
}

func (s *Scene) Init() {}
func (s *Scene) Update() {}
func (s *Scene) Draw() {}
func (s *Scene) GetParent() Entity { return nil }
func (s *Scene) GetChildren() []Entity { return s.Children }
func (s *Scene) Add(entities ...Entity) {
	for _, e := range entities {
		s.Children = append(s.Children, e)
	}
}
func (s *Scene) Remove(e Entity) {}
func (s *Scene) IsDeleted() bool { return false }

func (g *Game) GetCurrentScene() *GameObject {
	return g.Scenes[g.CurrentScene]
}

func (g *Game) SetScene(index int) {
    if index >= 0 && index < len(g.Scenes) {
        g.CurrentScene = index
    }
}