package goboom

// func GetGlobalX(r Renderable) float32 {
// 	if r.GetParent() == nil {
// 		return r.GetX()
// 	}
// 	return r.GetX() + GetGlobalX(r.GetParent())
// }

// func GetGlobalY(r Renderable) float32 {
// 	if r.GetParent() == nil {
// 		return r.GetY()
// 	}
// 	return r.GetY() + GetGlobalY(r.GetParent())
// }

// func GetGlobalAngle(r Renderable) float32 {
// 	if r.GetParent() == nil {
// 		return r.GetAngle()
// 	}
// 	return r.GetAngle() + GetGlobalAngle(r.GetParent())
// }

// func GetGlobalScaleX(r Renderable) float32 {
// 	if r.GetParent() == nil {
// 		return r.GetScaleX()
// 	}
// 	return r.GetScaleX() * GetGlobalScaleX(r.GetParent())
// }

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

// func Nest(parent Renderable, c ...Renderable) {
// 	for _, child := range c {

// 		parent.SetChildren(append(parent.GetChildren(), child))

// 		if child.GetParent() != nil {
// 			child.RemoveParent()
// 		}

// 		child.SetParent(parent)
// 	}
// }



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