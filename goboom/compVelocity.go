package goboom

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type VelocityComp struct {
	CompSetGet
	VelX, VelY 	float32
	Drag 		float32
}


func NewVelocityComp(vx, vy float32) *VelocityComp {
	comp := &VelocityComp{
		VelX: vx,
		VelY: vy,
		Drag: 1,
	}
	return comp
}


func (c *VelocityComp) GetComponentId() string {
	return "velocity"
}

func (c *VelocityComp) OnInit() {}

func (c *VelocityComp) OnUpdate(scene *GameObject) {
	obj := c.GameObject

	// Apply velocity
	obj.X += c.VelX
	obj.Y += c.VelY

	// Apply drag
	c.VelX *= c.Drag
	c.VelY *= c.Drag

	// clamp to zero
	if math.Abs(float64(c.VelX)) < 0.01 {
		c.VelX = 0
	}
	if math.Abs(float64(c.VelY)) < 0.01 {
		c.VelY = 0
	}
}

func (c *VelocityComp) OnDraw(scene *GameObject) {}

func (c *VelocityComp) GetVelX() float32 {
	return c.VelX
}

func (c *VelocityComp) GetVelY() float32 {
	return c.VelY
}

func (c *VelocityComp) GetVelXY() (float32, float32) {
	return c.VelX, c.VelY
}

// Alias
func (c *VelocityComp) Set(args ...interface{}) {
	c.SetVelocity(args...)
}

func (c *VelocityComp) SetVelocityByHeading(heading, speed float32) {
	// -90 to account for the fact that 0 degrees is to the right
	// I may need to revisit this later
	heading = (heading - 90) * rl.Deg2rad
	c.VelX = speed * float32(math.Cos(float64(heading)))
	c.VelY = speed * float32(math.Sin(float64(heading)))
}

func (c *VelocityComp) AddVelocityByHeading(heading, speed float32) {
	// -90 to account for the fact that 0 degrees is to the right
	// I may need to revisit this later
	heading = (heading - 90) * rl.Deg2rad
	c.VelX += speed * float32(math.Cos(float64(heading)))
	c.VelY += speed * float32(math.Sin(float64(heading)))
}

func (c *VelocityComp) ApplyDrag(drag float32) {
	c.Drag = drag
}

func (c *VelocityComp) SetVelocity(args ...interface{}) {
    if len(args) == 1 {
        if v, ok := args[0].(rl.Vector2); ok {
            c.VelX = v.X
            c.VelY = v.Y
            return
        }
    }

    if len(args) == 2 {
		// int case
		if vx, ok := args[0].(int); ok {
			if vy, ok := args[1].(int); ok {
				c.VelX = float32(vx)
				c.VelY = float32(vy)
				return
			}
		}
        // float32 case
        if vx, ok := args[0].(float32); ok {
            if vy, ok := args[1].(float32); ok {
                c.VelX = vx
                c.VelY = vy
                return
            }
        }
        // float64 case
        if vx, ok := args[0].(float64); ok {
            if vy, ok := args[1].(float64); ok {
                c.VelX = float32(vx)
                c.VelY = float32(vy)
                return
            }
        }
    }

    fmt.Println("Invalid arguments to SetVelocity")
}

func (c *VelocityComp) GetVelVector() rl.Vector2 {
	return rl.NewVector2(c.VelX, c.VelY)
}

func (c *VelocityComp) GetWidth() float32 {
	return 0
}

func (c *VelocityComp) GetHeight() float32 {
	return 0
}