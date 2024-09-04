package goboom

import rl "github.com/gen2brain/raylib-go/raylib"


func GetBoundingBox(r *GameObject) rl.Rectangle {
    rect := rl.Rectangle{
            X: GetGlobalX(r),
            Y: GetGlobalY(r),
            Width: GetGlobalWidth(r),
            Height: GetGlobalHeight(r)}

    return rect
}

func GetGlobalX(r *GameObject) float32 {
	if r.GetParent() == nil {
		return r.GetX()
	}
	return r.GetX() + GetGlobalX(r.GetParent())
}

func GetGlobalY(r *GameObject) float32 {
	if r.GetParent() == nil {
		return r.GetY()
	}
	return r.GetY() + GetGlobalY(r.GetParent())
}


func GetGlobalScaleX(r *GameObject) float32 {
	if r.GetParent() == nil {
		return r.GetScaleX()
	}
	return r.GetScaleX() * GetGlobalScaleX(r.GetParent())
}

func GetGlobalScaleY(r *GameObject) float32 {
	if r.GetParent() == nil {
		return r.GetScaleY()
	}
	return r.GetScaleY() * GetGlobalScaleY(r.GetParent())
}

func GetGlobalWidth(r *GameObject) float32 {
	
	if r.GetParent() == nil {
		return r.GetWidth()
	}
	
	return r.GetWidth() * GetGlobalScaleX(r.GetParent())
}

func GetGlobalHeight(r *GameObject) float32 {
	if r.GetParent() == nil {
		return r.GetHeight()
	}
	return r.GetHeight() * GetGlobalScaleY(r.GetParent())
}

func GetGlobalAngle(r *GameObject) float32 {
	if r.GetParent() == nil {
		return r.GetAngle()
	}
	return r.GetAngle() + GetGlobalAngle(r.GetParent())
}

