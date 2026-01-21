package sync_test

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		if counter.Value() != 3 {
			t.Errorf("got %d, want %d", counter.Value(), 3)
		}
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantCount)

		for range wantCount {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		wg.Wait()

		if counter.Value() != wantCount {
			t.Errorf("got %d, want %d", counter.Value(), wantCount)
		}
	})
}
