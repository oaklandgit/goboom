package goboom

import "time"

func (obj *GameObject) SetLifespan(millisecs int) *GameObject {
	
	go func() {
        time.Sleep(time.Duration(millisecs) * time.Millisecond)
        obj.Delete = true
    }()

	return obj
}