package goboom

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type VelocityComp struct {
	GameObject 		*GameObject
	VelX, VelY 		float32
}


func NewVelocityComp(vx, vy float32) *VelocityComp {
	comp := &VelocityComp{
		VelX: vx,
		VelY: vy,
	}
	return comp
}


func (c *VelocityComp) GetComponentId() string {
	return "velocity"
}

func (c *VelocityComp) SetGameObject(g *GameObject) {
	c.GameObject = g
}

func (c *VelocityComp) OnInit() {}

func (c *VelocityComp) OnUpdate() {
	obj := c.GameObject
	obj.X += c.VelX
	obj.Y += c.VelY
}

func (c *VelocityComp) OnDraw() {}

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