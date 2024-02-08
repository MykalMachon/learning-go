package greet

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	testBuffer := bytes.Buffer{}
	Greet(&testBuffer, "Mykal")

	got := testBuffer.String()
	want := "Hello, Mykal"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
