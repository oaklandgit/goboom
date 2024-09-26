package goboom

type Shape struct {
	Visible bool
	Opacity float32
	Width  float32
	Height float32
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

func (s *GameObject) SetWidth(w float32) {
	s.Width = w
}

func (s *GameObject) SetHeight(h float32) {
	s.Height = h
}

func (s *GameObject) SetSize(w, h float32) {
	s.SetWidth(w)
	s.SetHeight(h)
}

func (s *GameObject) GetWidth() float32 {
	return s.Width
}

func (s *GameObject) GetHeight() float32 {
	return s.Height
}


