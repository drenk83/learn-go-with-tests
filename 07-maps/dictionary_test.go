package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "it is test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "it is test"

		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		err := dict.Add("test", "test value")

		assertError(t, err, nil)
		assertDefinition(t, dict, "test", "test value")
	})

	t.Run("existing word", func(t *testing.T) {
		dict := Dictionary{"test": "test value"}
		err := dict.Add("test", "value")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dict, "test", "test value")
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		dict := Dictionary{"test": "value"}
		err := dict.Update("test", "newValue")

		assertError(t, err, nil)
		assertDefinition(t, dict, "test", "newValue")
	})

	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		err := dict.Update("test", "newValue")

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		dict := Dictionary{"test": "value"}
		err := dict.Delete("test")

		assertError(t, err, nil)

		_, err = dict.Search("test")
		assertError(t, err, ErrNotFound)
	})

	t.Run("non-existing word", func(t *testing.T) {
		dict := Dictionary{}
		err := dict.Delete("test")

		assertError(t, err, ErrWordDoesNotExist)
	})
}

// Helpers
func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("want: %q, got: %q", want, got)
	}
}

func assertDefinition(t testing.TB, dict Dictionary, key, value string) {
	t.Helper()

	got, err := dict.Search(key)

	if err != nil {
		t.Fatal("expected to get an error")
	}

	assertString(t, got, value)
}
