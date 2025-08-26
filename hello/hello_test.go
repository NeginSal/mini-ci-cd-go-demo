package hello

import "testing"

func TestHelloDefault(t *testing.T) {
	got := Hello("")
	want := "Hello, DevOps World!"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestHelloWithName(t *testing.T) {
	got := Hello("Negin")
	want := "Hello, Negin World!"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
