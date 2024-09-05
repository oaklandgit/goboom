package goboom

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Velocity struct {
	VelX float32
	VelY float32
}

func (v *Velocity) SetVelocity(x, y float32) {
	v.VelX = x
	v.VelY = y
}

func (v *Velocity) SetVelocityVect(vec rl.Vector2) {
	v.VelX = vec.X
	v.VelY = vec.Y
}

func (v *Velocity) GetVelocityXY() (float32, float32) {
	return v.VelX, v.VelY
}

func (v *Velocity) GetVelocityVect() rl.Vector2 {
	return rl.NewVector2(v.VelX, v.VelY)
}

func (v *Velocity) GetVelX() float32 {
	return v.VelX
}

func (v *Velocity) GetVelY() float32 {
	return v.VelY
}

func (v *Velocity) SetVelocityByHeading(heading, speed float32) {
	// -90 to account for the fact that 0 degrees is to the right
	// I may need to revisit this later
	heading = (heading - 90) * rl.Deg2rad
	
	v.VelX = speed * float32(math.Cos(float64(heading)))
	v.VelY = speed * float32(math.Sin(float64(heading)))
}