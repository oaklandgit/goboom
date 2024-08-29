package main

type Visibility struct {
	Visible bool
}

func (v *Visibility) SetVisible(visible bool) {
	v.Visible = visible
}

func (v *Visibility) IsVisible() bool {
	return v.Visible
}

