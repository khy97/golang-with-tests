package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

// CountdownOperationsSpy - Spy for CountDownOperations
type CountdownOperationsSpy struct {
	Calls []string
}

// Sleep - For ops test
func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// Write - for ops test
func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

// Sleeper for mocking
type Sleeper interface {
	Sleep()
}

// SpySleeper - For tests
type SpySleeper struct {
	Calls int
}

// Sleep for test
func (s *SpySleeper) Sleep() {
	s.Calls++
}

// DefaultSleeper for default run
type DefaultSleeper struct{}

// Sleep for default run
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const finalWord = "Go!"
const count = 3

// Countdown - Counts down
func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := count; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(writer, strconv.Itoa(i))
	}
	sleeper.Sleep()
	fmt.Fprint(writer, finalWord)
}

func main() {
	Countdown(os.Stdin, &DefaultSleeper{})
}
