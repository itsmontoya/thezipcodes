package thezipcodes

import (
	"time"
)

const rateLimitDuration = time.Second

func makeRateLimiter() (r rateLimiter) {
	r.c = make(chan semaphore)
	go r.scan()
	return
}

type rateLimiter struct {
	c chan semaphore
}

func (r *rateLimiter) scan() {
	var nextRequest time.Time
	for sem := range r.c {
		now := time.Now()
		if now.Before(nextRequest) {
			delta := nextRequest.Sub(now)
			time.Sleep(delta)
		}

		sem.push()
		nextRequest = time.Now().Add(rateLimitDuration)
	}
}

func (r *rateLimiter) Request() {
	sem := make(semaphore)
	r.c <- sem
	<-sem
}
