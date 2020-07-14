package syncs

import "sync"

// Counter - Counter object
type Counter struct {
	mu    sync.Mutex
	count int
}

// Inc - Increments the Counter
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

// Value - Returns the value
func (c *Counter) Value() int {
	return c.count
}

// NewCounter - creates new counter
func NewCounter() *Counter {
	return &Counter{}
}
