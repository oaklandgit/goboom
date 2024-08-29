package main

func (g *Game) Update(dt float32) {

	for _, input := range g.Inputs {
		CheckInput(input)
	}

	for _, obj := range g.GameObjects {
		obj.OnUpdate()
	}

}