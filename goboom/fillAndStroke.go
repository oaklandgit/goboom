package goboom

import rl "github.com/gen2brain/raylib-go/raylib"

type Stroke struct {
	StrokeWeight float32
	StrokeColor  rl.Color
}

type Fill struct {
	FillColor rl.Color
}

func (s *Stroke) SetStroke(color rl.Color, weight float32) *Stroke {
	s.StrokeWeight = weight
	s.StrokeColor = color
	return s
}

func (f *Fill) SetFill(color rl.Color) *Fill {
	f.FillColor = color
	return f
}
