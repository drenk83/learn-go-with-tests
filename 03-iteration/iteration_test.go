package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("test_1", func(t *testing.T){
		got := Repeat("a", 5)
		want := "aaaaa"

		if got != want {
			t.Errorf("want: %q, got: %q", want, got)
		}
	})

	t.Run("test_2", func(t *testing.T){
		got := Repeat("a", 10)
		want := "aaaaaaaaaa"

		if got != want {
			t.Errorf("want: %q, got: %q", want, got)
		}
	})

	t.Run("test_3", func(t *testing.T){
		got := Repeat("b", 100)
		want := "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"

		if got != want {
			t.Errorf("want: %q, got: %q", want, got)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 100)
	}
}