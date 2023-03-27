package thezipcodes

import (
	"fmt"
	"testing"
	"time"
)

func Test_rateLimiter_Request(t *testing.T) {
	r := makeRateLimiter()
	start := time.Now()
	r.Request()
	r.Request()
	r.Request()
	end := time.Now()
	delta := int(end.Sub(start) / time.Second)
	if delta != 2 {
		t.Fatalf("expected delta of %d seconds and received %d seconds", 2, delta)
	}
	fmt.Println("Duration", delta)
}
