package goboom

func (h *GameObject) SetGame(game *Game) {
	h.Game = game
}

func (h *GameObject) GetGame() *Game {
	return h.Game
}

func (h *GameObject) SetParent(parent *GameObject) {
	h.Parent = parent
}

func (h *GameObject) GetParent() *GameObject {
	return h.Parent
}

func (h *GameObject) RemoveParent() {
	h.Parent = nil
}

func (h *GameObject) Add(children ...*GameObject) {
	for _, c := range children {
		h.Children = append(h.Children, c)

		if c.GetParent() != nil {
			c.RemoveParent()
		}

		c.SetParent(h)
		c.SetGame(h.GetGame())
	}
}

func (h *GameObject) GetChildren() []*GameObject {
	return h.Children
}


func (h *GameObject) Remove(child *GameObject) {
	for i, c := range h.Children {
		if c == child {
			h.Children = append(h.Children[:i], h.Children[i+1:]...)
			break
		}
	}
}