package greetings

import (
	"testing"
)

func TestHello(t *testing.T) {
	name := "vector"
	want := "Hi, vector. Welcome!"
	if got, _ := Hello(name); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
