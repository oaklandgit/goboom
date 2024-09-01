package goboom

func (g *Game) Update() {

	for _, obj := range g.GameObjects {

		if obj.IsDeleted() {
			g.Remove(obj)
			continue
		}

		obj.Update()
		obj.CheckInput()
	}

}
