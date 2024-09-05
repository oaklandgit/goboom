package goboom

func (g *Game) Update() {

	g.CheckInput()

	scene := g.GetCurrentScene()
	scene.OnUpdate()

	for _, obj := range scene.GetChildren() {

		if obj.IsDeleted() {
			scene.Remove(obj)
			continue
		}

		obj.OnUpdate()
		obj.OnWrap()
		obj.CheckInput()
	}

}
