### Notes

- component shapes should always draw at 0, 0. Then let their position be contolled by their GameObjects.
- Scene graph: for draw functions, don't adjust object positions to global. The draw routine will do that using push/pop matrix. BUT do use global x/y when calculating collisions, drawing debug boxes, deciding where an object (e.g. bullet) should spawn, etc.

### Up Next

- move collision shape drawing responsibility to component
- all shapes are not centering the same. E.g. ellipses/circles vs. non-regular polygons

### To Do

- ensure objects with collision components have tags
- pass a customizable "cooldown" period into a collision event
- offscreen die
- utilize delta time
- load sprite atlas
- define animations from sprite atlas
- implement motion scripts / tweening (e.g. for enemy movement)
- platformer example
  - gravity
  - grounded
  - jump component
- compAnchor - a way to keep objects dynamically anchored (centered, etc) to their parent

### Done

- why isn't groupArray in missile example not PutCenter-ing properly?
- centralize collision detection. Current approach is non-performant
- "MousleCommand" - mouse control example
- Debug collision boxes
- bring back debug modes
- rotational velocity component
- get collisions working (start with circle collisions)
- figure out why bricks example isn't rendering the brick group properly
- every object should have a unique id. For groups of objects append an index
- tilemaps
- Implement convenience functions for velocities
