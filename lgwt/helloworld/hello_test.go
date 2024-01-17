package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris!"
		assetCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"
		assetCorrectMessage(t, got, want)
	})
}

func assetCorrectMessage(t testing.TB, got, want string) {
	// is needed to tell the test suite that this method is a helper.
	// By doing this when it fails the line number reported will be in our function call rather than inside our test helper.
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
