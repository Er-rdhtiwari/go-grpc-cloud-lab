package clock

import "time"

// Clock is a tiny interface to make time deterministic in tests.
// WHY: time.Now() is nondeterministic; Clock lets tests control time.
type Clock interface {
	Now() time.Time
}

type RealClock struct{}

func (RealClock) Now() time.Time { return time.Now().UTC() }
