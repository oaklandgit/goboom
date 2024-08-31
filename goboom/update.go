package goboom

func (g *Game) Update(dt float32) {



	for _, obj := range g.GameObjects {

		if obj.IsDeleted() {
			g.Remove(obj)
			continue
		}

		obj.OnUpdate()
		obj.OnInput()
	}

}