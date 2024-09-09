package goboom

type Shape struct {
	Visible bool
	Opacity float32

	Width  float32
	Height float32
	// GetWidth func() float32
	// GetHeight func() float32
}

func (s *Shape) SetVisible(visible bool) {
	s.Visible = visible
}

func (s *Shape) IsVisible() bool {
	return s.Visible
}

func (s *GameObject) CenterX() float32 {
	return s.GetX() + s.GetWidth() / 2

}

func (s *GameObject) CenterY() float32 {
	return s.GetY() + s.GetHeight() / 2
}

func (s *GameObject) GetWidth() float32 {
	// w := s.Width
	// for _, comp := range s.GetComponents() {
	// 	w = comp.ModifyWidth(w)
	// }
	// return w
	return s.Width
}

func (s *GameObject) GetHeight() float32 {
	// h := s.Height
	// for _, comp := range s.GetComponents() {
	// 	h = comp.ModifyWidth(h)
	// }
	// return h
	return s.Height
}


