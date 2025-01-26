package main

import "testing"

func TestHello(t *testing.T) {
	wants := "Hello Go"

	got := hello()

	if wants != got {
		t.Fatalf("want %s, got %s\n", wants, got)
	}
}
