package goboom

type Hierarchy struct {
	Game *Game
	Parent Renderable
	Children []Renderable
}

func GetGlobalX(r Renderable) float32 {
	if r.GetParent() == nil {
		return r.GetX()
	}
	return r.GetX() + GetGlobalX(r.GetParent())
}

func GetGlobalY(r Renderable) float32 {
	if r.GetParent() == nil {
		return r.GetY()
	}
	return r.GetY() + GetGlobalY(r.GetParent())
}

func GetGlobalAngle(r Renderable) float32 {
	if r.GetParent() == nil {
		return r.GetAngle()
	}
	return r.GetAngle() + GetGlobalAngle(r.GetParent())
}

func GetGlobalScaleX(r Renderable) float32 {
	if r.GetParent() == nil {
		return r.GetScaleX()
	}
	return r.GetScaleX() * GetGlobalScaleX(r.GetParent())
}

func (h *Hierarchy) SetGame(game *Game) {
	h.Game = game
}

func (h *Hierarchy) GetGame() *Game {
	return h.Game
}

func (h *Hierarchy) SetParent(parent Renderable) {
	h.Parent = parent
}

func (h *Hierarchy) GetParent() Renderable {
	return h.Parent
}

func (h *Hierarchy) RemoveParent() {
	h.Parent = nil
}

func (h *Hierarchy) Add(children ...Renderable) {
	for _, c := range children {
		h.Children = append(h.Children, c)

		if c.GetParent() != nil {
			c.RemoveParent()
		}

		// c.SetParent(h) TO DO
	}
}

func (h *Hierarchy) GetChildren() []Renderable {
	return h.Children
}


func (h *Hierarchy) Remove(child Renderable) {
	for i, c := range h.Children {
		if c == child {
			h.Children = append(h.Children[:i], h.Children[i+1:]...)
			break
		}
	}
}