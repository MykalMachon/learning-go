package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("Hello function", func(t *testing.T) {
		got := SayHello()
		want := "Hello, world!"
		if got != want {
			t.Errorf("got %q and wanted %q", got, want)
		}
	})
}
