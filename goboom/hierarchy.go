package goboom

func (obj *GameObject) SetGame(game *Game) {
	obj.Game = game
}

func (obj *GameObject) GetGame() *Game {
	return obj.Game
}

func (obj *GameObject) SetParent(parent *GameObject) {
	obj.Parent = parent
}

func (obj *GameObject) GetParent() *GameObject {
	return obj.Parent
}

func (obj *GameObject) RemoveParent() {
	obj.Parent = nil
}

func (obj *GameObject) Add(children ...*GameObject) {
	for _, c := range children {
		obj.Children = append(obj.Children, c)

		if c.GetParent() != nil {
			c.RemoveParent()
		}

		c.SetParent(obj)
		// c.SetGame(obj.GetGame())
	}
}

func (obj *GameObject) GetChildren() []*GameObject {
	return obj.Children
}


func (obj *GameObject) Remove(child *GameObject) {
	for i, c := range obj.Children {
		if c == child {
			obj.Children = append(obj.Children[:i], obj.Children[i+1:]...)
			break
		}
	}
}