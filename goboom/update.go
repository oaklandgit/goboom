package goboom

func (g *Game) Update() {

	g.CheckInput()

	for _, obj := range g.GameObjects {

		if obj.IsDeleted() {
			g.Remove(obj)
			continue
		}

		obj.Update()
		obj.CheckInput()
	}

}
