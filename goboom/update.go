package goboom

func (g *Game) Update() {

	g.CheckInput()

	scene := g.GetCurrentScene()
	scene.Update()

	for _, obj := range scene.GetChildren() {

		if obj.IsDeleted() {
			scene.Remove(obj)
			continue
		}

		obj.Update()
		obj.CheckInput()
	}

}
