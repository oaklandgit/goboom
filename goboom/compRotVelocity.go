package goboom

import (
	"math"
)

type RotVelocityComp struct {
	CompSetGet
	Speed 		float32
	Drag 		float32
}


func NewRotVelocityComp(speed float32) *RotVelocityComp {
	comp := &RotVelocityComp{
		Speed: speed,
		Drag: 1,
	}
	return comp
}


func (c *RotVelocityComp) GetComponentId() string {
	return "rotvelocity"
}

func (c *RotVelocityComp) OnInit() {}

func (c *RotVelocityComp) OnUpdate(scene *GameObject) {
	obj := c.GameObject

	obj.Angle += c.Speed

	// Apply drag
	c.Speed *= c.Drag

	// clamp to zero
	if math.Abs(float64(c.Speed)) < 0.01 {
		c.Speed = 0
	}
}

func (c *RotVelocityComp) OnDraw(scene *GameObject) {}

func (c *RotVelocityComp) GetRotVel() float32 {
	return c.Speed
}

func (c *RotVelocityComp) SetSpeed(speed float32) {
	c.Speed = speed
}

func (c *RotVelocityComp) SetDrag(drag float32) {
	c.Drag = drag
}