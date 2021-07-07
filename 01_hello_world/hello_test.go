package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Osvaldo")
	want := "Hello, Osvaldo"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
