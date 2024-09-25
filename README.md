### Notes

- Scene graph: for draw functions, don't adjust object positions to global. The draw routine will do that using push/pop matrix. BUT do use global x/y when calculating collisions, drawing debug boxes, etc.

### To Do

- get collisions working (start with circle collisions)
- figure out why bricks example isn't rendering the brick group properly
- every object should have a unique id. For groups of objects append an index
- utilize delta time

### Backlog

- load sprite atlas
- tilemaps
- define animations from sprite atlas
- implement motion scripts / tweening (e.g. for enemy movement)
- Implement convenience functions for velocities

### Done
