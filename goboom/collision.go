package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Collider struct {
	Vs     string
	Action func(obj1, obj2 *GameObject)
}

type Collisions struct {
	CollidingWith []*GameObject
	Colliders     []*Collider
}


func detectCollision(obj1, obj2 *GameObject) bool {
	a := obj1.GetBoundingBox()
	b := obj2.GetBoundingBox()
	return  rl.CheckCollisionRecs(a, b)
}


// this should be a root game object
func (node *GameObject) CheckCollisions() {

	objs := node.GetAllTagged()
	
	// Check for collisions
	for i, o1 := range objs {
		for j, o2 := range objs {
			if i != j {

				//reset collidingWith
				o1.CollidingWith = o1.CollidingWith[:0]
				o2.CollidingWith = o2.CollidingWith[:0]

				// re-detect collisions
				if detectCollision(o1, o2) {
					
					o1.CollidingWith = append(o1.CollidingWith, o2)
					o2.CollidingWith = append(o2.CollidingWith, o1)

					// check for colliders
					for _, c := range o1.Colliders {
						if c.Vs == o2.Tags[0] {
							c.Action(o1, o2)
						}
					}
				}
			}
		}
	}
}

func (o *GameObject) AddCollider(vs string, action func(obj1, obj2 *GameObject)) *GameObject {
	o.Colliders = append(o.Colliders, &Collider{Vs: vs, Action: action})
	return o
}
