package tools

import "time"

// Timer defines a handler to wrap all date/time management
type Timer interface {
	// Now returns the current timestamp
	Now() time.Time
}

type timer struct{}

// NewTimer returns an instance of Timer
func NewTimer() Timer {
	return timer{}
}

// Now returns the current timestamp
func (t timer) Now() time.Time {
	return time.Now()
}
