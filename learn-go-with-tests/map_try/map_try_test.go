package map_try

import (
	"testing"
)

func TestXxx(t *testing.T) {
	dictionary := Dictionary{"test": "test value"}

	t.Run("existed word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "test value"

		assertString(t, got, want)
	})

	t.Run("no existed word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrMsg

		if err == nil {
			t.Fatal("expect to get an error")
		}

		assertString(t, err.Error(), want)
	})
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
