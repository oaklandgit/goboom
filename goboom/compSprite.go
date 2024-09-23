package goboom

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type SpriteComp struct {
	GameObject *GameObject
	ImagePaths []string
	Frames []rl.Texture2D
	CurrentFrame int
	Animations map[string]Animation
	CurrentAnim string
	ElapsedTime float32
}

type Animation struct {
	Frames []int
	Speed int
	Loop bool
}

func Sprite(x, y float32, imagePaths ...string) *GameObject {
	obj := NewGameObject()
	obj.X = x
	obj.Y = y

	comp := NewSpriteComp(imagePaths...)
	obj.AddComponent(comp)
	
	return obj
}


func NewSpriteComp(imagePaths ...string) *SpriteComp {
	return &SpriteComp{
		ImagePaths: imagePaths,
		Animations: make(map[string]Animation),
	}
}

func (c *SpriteComp) GetComponentId() string {
	return "sprite"
}

func (c *SpriteComp) SetGameObject(g *GameObject) {
	c.GameObject = g
}

func (c *SpriteComp) GetGameObject() *GameObject {
	return c.GameObject
}

func (c *SpriteComp) OnInit() {
	fmt.Println("Initializing sprite component")
	for _, path := range c.ImagePaths {
		c.Frames = append(c.Frames, rl.LoadTexture(path))
	}
}

func (c *SpriteComp) OnUpdate(scene *GameObject) {
	if c.CurrentAnim == "" {
        return
    }

	// fmt.Println(c.ElapsedTime)
	anim := c.Animations[c.CurrentAnim]
	c.ElapsedTime += rl.GetFrameTime()

	if c.ElapsedTime >= 1.0/float32(anim.Speed) {
        c.ElapsedTime = 0
        c.CurrentFrame++
        if c.CurrentFrame >= len(anim.Frames) {
            if anim.Loop {
                c.CurrentFrame = 0
            } else {
                c.CurrentFrame = len(anim.Frames) - 1
            }
        }
    }
}

func (c *SpriteComp) OnDraw(scene *GameObject) {
	obj := c.GameObject
	centerX := obj.X + obj.GetWidth() * obj.GetOriginX()
	centerY := obj.Y + obj.GetHeight() * obj.GetOriginY()
	rl.PushMatrix()
	rl.Translatef(centerX, centerY, 0)
	rl.Scalef(obj.GetScaleX(), obj.GetScaleY(), 1)
	rl.Rotatef(obj.GetAngle(), 0, 0, 1)
	rl.Translatef(-centerX, -centerY, 0)
	
	rl.DrawTexture(
		c.Frames[c.CurrentFrame],
		int32(obj.X - obj.GetWidth() * obj.GetOriginX()),
		int32(obj.Y - obj.GetHeight() * obj.GetOriginY()),
		rl.White)
	rl.PopMatrix()
}

func (c *SpriteComp) ModifyWidth(w float32) float32 {
	localW := float32(c.Frames[c.CurrentFrame].Width)
	if localW > w {
		return localW
	}
	return w
}

func (c *SpriteComp) ModifyHeight(h float32) float32 {
	localH := float32(c.Frames[c.CurrentFrame].Height)
	if localH > h {
		return localH
	}
	return h
}


func (c *SpriteComp) AddFrame(path string) {
	c.ImagePaths = append(c.ImagePaths, path)
}

func (c *SpriteComp) AddAnim(id string, frames []int, speed int, loop bool) {
	c.Animations[id] =
		Animation{
			Frames: frames,
			Speed: speed,
			Loop: loop,
		}
}

func (c *SpriteComp) Play(id string) {
	c.CurrentAnim = id
	c.CurrentFrame = 0
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
