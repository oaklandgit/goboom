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
	
	if c.GameObject.Width == 0 {
		c.GameObject.Width = float32(c.Frames[0].Width)
	}

	if c.GameObject.Height == 0 {
		c.GameObject.Height = float32(c.Frames[0].Height)
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

	rl.DrawTexturePro(
		c.Frames[c.CurrentFrame],
		rl.NewRectangle(0, 0, float32(c.Frames[c.CurrentFrame].Width), float32(c.Frames[c.CurrentFrame].Height)), // Source rectangle
		rl.NewRectangle(
			obj.X - obj.GetWidth()*obj.GetOriginX(), // Destination X
			obj.Y - obj.GetHeight()*obj.GetOriginY(), // Destination Y
			obj.GetWidth(), // Destination width
			obj.GetHeight(), // Destination height
		),
		rl.NewVector2(obj.GetWidth()*obj.GetOriginX(), obj.GetHeight()*obj.GetOriginY()), // Origin
		0, // Rotation
		rl.White, // Tint
	)
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
