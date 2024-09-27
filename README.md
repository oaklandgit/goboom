### Notes

- component shapes should always draw at 0, 0. Then let their position be contolled by their GameObjects.
- Scene graph: for draw functions, don't adjust object positions to global. The draw routine will do that using push/pop matrix. BUT do use global x/y when calculating collisions, drawing debug boxes, deciding where an object (e.g. bullet) should spawn, etc.

### In Progress

- "MousleCommand" - mouse control example

### To Do

- compAnchor - a way to keep objects dynamically anchored (centered, etc) to their parent
- utilize delta time
- load sprite atlas
- define animations from sprite atlas
- implement motion scripts / tweening (e.g. for enemy movement)
- platformer example
  - gravity
  - grounded
  - jump component

### Done

- rotational velocity component
- get collisions working (start with circle collisions)
- figure out why bricks example isn't rendering the brick group properly
- every object should have a unique id. For groups of objects append an index
- tilemaps
- Implement convenience functions for velocities
