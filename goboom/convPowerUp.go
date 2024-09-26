package goboom

import (
	"time"
)

func PowerUp(r *GameObject,  f func(*GameObject) func(*GameObject), d time.Duration) func() {

	revert := f(r) // do the function and set revert to the function returned by f
	cancelChannel := make(chan struct{})

	go func() {
		select {
		case <-time.After(d):
			if revert != nil {
				revert(r)
			}
		case <-cancelChannel:
			if revert != nil {
				revert(r)
			}
		}
	}()

	return func() {
		close(cancelChannel)
	}
	
}