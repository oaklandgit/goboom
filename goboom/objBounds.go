package goboom

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (obj *GameObject) GetBoundingBox() rl.Rectangle {
    minX := obj.X
    minY := obj.Y
    maxX := minX + obj.Width // GameObject's own width
    maxY := minY + obj.Height // GameObject's own height

    // Check children GameObjects
    for _, child := range obj.Children {
        childBox := child.GetBoundingBox()
        minX = float32(math.Min(float64(minX), float64(childBox.X)))
        minY = float32(math.Min(float64(minY), float64(childBox.Y)))
        maxX = float32(math.Max(float64(maxX), float64(childBox.Width + childBox.X)))
        maxY = float32(math.Max(float64(maxY), float64(childBox.Height + childBox.Y)))
    }

	return rl.Rectangle{
		X: minX,
		Y: minY,
		Width: maxX - minX,
		Height: maxY - minY,
	}
}

func RotatePoint(x, y, angleDegrees float32) (float32, float32) {
	rads := angleDegrees * (math.Pi / 180)
	sin, cos := float32(math.Sin(float64(rads))), float32(math.Cos(float64(rads)))
    return x*cos - y*sin, x*sin + y*cos
}

func (r *GameObject) GetGlobalX() float32 {
	x, _ := r.GetGlobalXY()
	return x
}


func (r *GameObject) GetGlobalY() float32 {
	_, y := r.GetGlobalXY()
	return y
}


func (r *GameObject) GetGlobalXY() (float32, float32) {
	if r.GetParent() == nil {
		return r.GetX(), r.GetY()
	}
	parent := r.GetParent()
	localX, localY := r.GetX(), r.GetY()
	rotatedX, rotatedY := RotatePoint(localX, localY, parent.GetAngle())
	return rotatedX + parent.GetGlobalX(), rotatedY + parent.GetGlobalY()
}

// func (r *GameObject) GetGlobalY() float32 {
// 	if r.GetParent() == nil {
// 		return r.GetY()
// 	}
// 	return r.GetY() + r.GetParent().GetGlobalY()
// }


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

