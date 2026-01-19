package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls []string
}

func (s *SpySleeper) Sleep() {
	s.Calls = append(s.Calls, "sleep")
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleepr Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleepr.Sleep()
	}
	fmt.Fprintf(out, "Go!")
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
