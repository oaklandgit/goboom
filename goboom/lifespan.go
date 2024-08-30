package goboom

import "time"

type Lifespan struct {
	Delete bool
}

func (l *Lifespan) IsDeleted() bool {
	return l.Delete
}

func (s *Shape) SetLifespan(millisecs int) *Shape {
	
	go func() {
        time.Sleep(time.Duration(millisecs) * time.Millisecond)
        s.Delete = true
    }()

	return s
}