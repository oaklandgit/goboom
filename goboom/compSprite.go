package goboom

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type SpriteComp struct {
	GameObject *GameObject
	Paths []string
	Frames []rl.Texture2D
	CurrentFrame int
	Animations map[string][]Animation
	CurrentAnim string
}

type Animation struct {
	Frames []int
	Speed int
	Loop bool
}


func NewSpriteComp(paths ...string) *SpriteComp {
	return &SpriteComp{
		Paths: paths,
	}
}

func (c *SpriteComp) GetComponentId() string {
	return "sprite"
}

func (c *SpriteComp) SetGameObject(g *GameObject) {
	c.GameObject = g
}

func (c *SpriteComp) OnInit() {
	fmt.Println("Initializing sprite component")
	for _, path := range c.Paths {
		c.Frames = append(c.Frames, rl.LoadTexture(path))
	}
}

func (c *SpriteComp) OnUpdate() {}

func (c *SpriteComp) OnDraw() {
	obj := c.GameObject
	centerX := obj.X + c.GetWidth() * obj.GetOriginX()
	centerY := obj.Y + c.GetHeight() * obj.GetOriginY()
	rl.PushMatrix()
	rl.Translatef(centerX, centerY, 0)
	rl.Scalef(obj.GetScaleX(), obj.GetScaleY(), 1)
	rl.Rotatef(obj.GetAngle(), 0, 0, 1)
	rl.Translatef(-centerX, -centerY, 0)
	
	rl.DrawTexture(
		c.Frames[c.CurrentFrame],
		int32(obj.X - c.GetWidth() * obj.GetOriginX()),
		int32(obj.Y - c.GetHeight() * obj.GetOriginY()),
		rl.White)
	rl.PopMatrix()
}

func (c *SpriteComp) GetWidth() float32 {
	return float32(c.Frames[c.CurrentFrame].Width)
}

func (c *SpriteComp) GetHeight() float32 {
	return float32(c.Frames[c.CurrentFrame].Height)
}

func (c *SpriteComp) AddFrame(path string) {
	c.Paths = append(c.Paths, path)
}

func (c *SpriteComp) AddAnim(id string, frames []int, speed int, loop bool) {
	c.Animations[id] = append(c.Animations[id],
		Animation{
			Frames: frames,
			Speed: speed,
			Loop: loop,
		})
}
	

func (c *SpriteComp) NextFrame() {
	c.CurrentFrame++
	if c.CurrentFrame >= len(c.Frames) {
		c.CurrentFrame = 0
	}
}

func (c *SpriteComp) PrevFrame() {
	c.CurrentFrame--
	if c.CurrentFrame < 0 {
		c.CurrentFrame = len(c.Frames) - 1
	}
}
