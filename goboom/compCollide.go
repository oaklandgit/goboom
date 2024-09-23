package goboom

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type CollideComp struct {
	GameObject 		*GameObject // my game object
	Shape 			ColliderShape // my shape for collisions
	Tags 			[]string // My identities
	Colliders 		[]Collider // who to collide with and what to do
	CollidingWith 	[]*GameObject // who am I currently colliding with?
}

type ColliderShape interface {
	GetBoundingBox() rl.Rectangle
	IsCollidingWith(other *CollideComp) bool
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
	return rl.NewVector2(0, 0)
}

func (c CollisionCircle) GetRadius() float32 {
	return c.Radius
}

func (c CollisionCircle) IsCollidingWith(other *CollideComp) bool {
	if other.Shape == nil {
		return false
	}
	
    if rect, ok := other.Shape.(*CollisionRect); ok {
        return rl.CheckCollisionCircleRec(c.GetCenter(), c.GetRadius(), rect.GetBoundingBox())
    }

    if circle, ok := other.Shape.(*CollisionCircle); ok {
        return rl.CheckCollisionCircles(c.GetCenter(), c.GetRadius(), circle.GetCenter(), circle.GetRadius())
    }

	return false
}

// RECT COLLISION

type CollisionRect struct {
	Width float32
	Height float32
}

func (c CollisionRect) GetBoundingBox() rl.Rectangle {
	return rl.NewRectangle(0, 0, c.Width, c.Height)
}

func (c CollisionRect) IsCollidingWith(other *CollideComp) bool {
	if other.Shape == nil {
		return false
	}

    if rect, ok := other.Shape.(*CollisionRect); ok {
        return rl.CheckCollisionRecs(c.GetBoundingBox(), rect.GetBoundingBox())
    }

    if circle, ok := other.Shape.(*CollisionCircle); ok {
        return rl.CheckCollisionCircleRec(circle.GetCenter(), circle.GetRadius(), c.GetBoundingBox())
    }

    return false
	
}


type Collider struct {
	Vs     string
	Action func(obj1, obj2 *GameObject)
}

func NewCollideComp(shape ColliderShape, tags ...string) *CollideComp {
	comp := &CollideComp{
		Shape: shape,
		Tags: tags,
	}
	return comp
}

func (c *CollideComp) NewCollider(vs string, action func(obj1, obj2 *GameObject)) Collider {

	collider := Collider{
		Vs: vs,
		Action: action,
	}

	return collider
}

func (c *CollideComp) GetComponentId() string {
	return "collide"
}

func (c *CollideComp) SetGameObject(g *GameObject) {
	c.GameObject = g
}

func (c *CollideComp) GetGameObject() *GameObject {
	return c.GameObject
}


func (c *CollideComp) OnInit() {}

func (c *CollideComp) OnUpdate(scene *GameObject) {

	// fmt.Println("HELLO", c.GetGameObject().GetTags())
	fmt.Println("YO", scene.GetGame().GetTitle())
	// fmt.Printf("Scene: %+v\n", c.GameObject.GetScene())
	// log.Fatal()

	objs := scene.GetAll()
	collideObjs := []*GameObject{}

	for _, o := range objs {
		if o.HasComponent("collide") {
			collideObjs = append(collideObjs, o)
		}
	}

	// reset collidingWith
	for _, o := range collideObjs {
		o.GetComponent("collide").(*CollideComp).CollidingWith =
			o.GetComponent("collide").(*CollideComp).CollidingWith[:0]
	}
	
	for i, o1 := range collideObjs {
		for j, o2 := range collideObjs {
			if i != j {

				// get the collide components
				c1 := o1.GetComponent("collide").(*CollideComp)
				c2 := o2.GetComponent("collide").(*CollideComp)

				
				// check if they are colliding
				if c1.Shape.IsCollidingWith(c2) {

					fmt.Println("COLLISION!", c1.Tags, "and", c2.Tags)
					c1.CollidingWith = append(c1.CollidingWith, o2)
					c2.CollidingWith = append(c2.CollidingWith, o1)

					// check for colliders
					for _, c := range c1.Colliders {
						if c.Vs == o2.Tags[0] {
							c.Action(o1, o2)
						}
					}

					for _, c := range c2.Colliders {
						if c.Vs == o1.Tags[0] {
							c.Action(o2, o1)
						}
					}

				}

			}
		}
	}
}

func (c *CollideComp) OnDraw(scene *GameObject) {

	objs := scene.GetAll()

	for _, o := range objs {
		if o.HasComponent("collide") {
			c := o.GetComponent("collide").(*CollideComp)
			b := c.Shape.GetBoundingBox()
			rl.DrawRectangleLines(int32(o.X+b.X), int32(o.Y+b.Y), int32(b.Width), int32(b.Height), rl.Red)
		}
	}

}
