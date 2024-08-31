package goboom

import "time"

type Lifespan struct {
	Delete bool
}

func (l *Lifespan) IsDeleted() bool {
	return l.Delete
}

func (obj *GameObject) SetLifespan(millisecs int) *GameObject {
	
	go func() {
        time.Sleep(time.Duration(millisecs) * time.Millisecond)
        obj.Delete = true
    }()

	return obj
}