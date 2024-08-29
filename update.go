package main

func (g *Game) Update(dt float32) {

	for _, obj := range g.GameObjects {
		obj.OnUpdate()
	}

}