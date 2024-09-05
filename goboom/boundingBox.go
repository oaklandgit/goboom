package goboom

import rl "github.com/gen2brain/raylib-go/raylib"


func (r *GameObject) GetBoundingBox() rl.Rectangle {
    rect := rl.Rectangle{
            X: r.GetGlobalX() + r.GetGlobalWidth() / 2,
            Y: r.GetGlobalY() + r.GetGlobalHeight() / 2,
            Width: r.GetGlobalWidth(),
            Height: r.GetGlobalHeight()}

    return rect
}

func (r *GameObject) GetGlobalX() float32 {
	if r.GetParent() == nil {
		return r.GetX()
	}
	return r.GetX() + r.GetParent().GetGlobalX()
}

func (r *GameObject) GetGlobalY() float32 {
	if r.GetParent() == nil {
		return r.GetY()
	}
	return r.GetY() + r.GetParent().GetGlobalY()
}


func (r *GameObject) GetGlobalScaleX() float32 {
	if r.GetParent() == nil {
		return r.GetScaleX()
	}
	return r.GetScaleX() * r.GetParent().GetGlobalScaleX()
}

func (r *GameObject) GetGlobalScaleY() float32 {
	if r.GetParent() == nil {
		return r.GetScaleY()
	}
	return r.GetScaleY() * r.GetParent().GetGlobalScaleY()
}

func (r *GameObject) GetGlobalWidth() float32 {
	
	if r.GetParent() == nil {
		return r.GetWidth()
	}
	
	return r.GetWidth() * r.GetParent().GetGlobalScaleX()
}

func (r *GameObject) GetGlobalHeight() float32 {
	if r.GetParent() == nil {
		return r.GetHeight()
	}
	return r.GetHeight() * r.GetParent().GetGlobalScaleY()
}

func (r *GameObject) GetGlobalAngle() float32 {
	if r.GetParent() == nil {
		return r.GetAngle()
	}
	return r.GetAngle() + r.GetParent().GetGlobalAngle()
}

