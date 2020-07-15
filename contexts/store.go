package contexts

import (
	"fmt"
	"net/http"
	"time"
)

// Server - returns a HTTP Handler func
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}

// Store interface
type Store interface {
	Fetch() string
	Cancel()
}

// SpyStore - Spy
type SpyStore struct {
	response  string
	cancelled bool
}

// Fetch - Wait
func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

// Cancel - cancel
func (s *SpyStore) Cancel() {
	s.cancelled = true
}
