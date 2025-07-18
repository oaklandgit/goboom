package goboom

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// CollideComp is a component that allows an object to collide with other objects
type CollideComp struct {
	CompSetGet
	Shape 				ColliderShape // my shape for collisions
	Colliders 			[]Collider // who to collide with and what to do
	CollidingWith 		[]*GameObject // who am I currently colliding with?
	LastCollisionTime 	map[*GameObject]time.Time 
}


// ColliderShape is an interface for collision shapes
type ColliderShape interface {
	GetBoundingBox() rl.Rectangle
}

// CIRCLE COLLISION

type CollisionCircle struct {
	Radius float32
}

func (c CollisionCircle) GetBoundingBox() rl.Rectangle {
	r := c.Radius
	return rl.NewRectangle(-r, -r, r*2, r*2)
}

func (c CollisionCircle) GetCenter() rl.Vector2 {
	return rl.NewVector2(-c.Radius, -c.Radius)
}

func (c CollisionCircle) GetRadius() float32 {
	return c.Radius
}

// RECT COLLISION

type CollisionRect struct {
	Width float32
	Height float32
}

func (c CollisionRect) GetBoundingBox() rl.Rectangle {
	return rl.NewRectangle(c.Width/2, c.Height/2, c.Width, c.Height)
}

type Collider struct {
	Vs     string
	Action func(obj1, obj2 *GameObject)
}

func NewCollideComp(shape ColliderShape, tags ...string) *CollideComp {
	comp := &CollideComp{
		Shape: shape,
		LastCollisionTime: make(map[*GameObject]time.Time), 
	}
	return comp
}

func (c *CollideComp) NewCollider(vs string, action func(obj1, obj2 *GameObject)) *CollideComp {

	collider := Collider{
		Vs: vs,
		Action: action,
	}

	c.Colliders = append(c.Colliders, collider)

	return c
}

func (c *CollideComp) GetComponentId() string {
	return "collide"
}

func (c *CollideComp) OnInit() {
	if rect, ok := c.Shape.(CollisionRect); ok &&
		rect.GetBoundingBox().Width == 0 &&
		rect.GetBoundingBox().Height == 0 {
        	obj := c.GameObject
        	c.Shape = CollisionRect{
				Width:  obj.GetBoundingBox().Width,
				Height: obj.GetBoundingBox().Height,
        	}
    }

	if circle, ok := c.Shape.(CollisionCircle); ok &&
		circle.GetBoundingBox().Width == 0 &&
		circle.GetBoundingBox().Height == 0 {
        	obj := c.GameObject
        	c.Shape = CollisionCircle{
				Radius:  obj.GetBoundingBox().Width/2,
        	}
    }

}

func (c *CollideComp) OnUpdate(scene *GameObject) {}

func (c *CollideComp) OnDraw(scene *GameObject) {}
