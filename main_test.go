package main

import "testing"

func TestHelloWorld(t *testing.T) {
	t.Logf("testing hello world")

	got := greeting("Joel")

	wanted := "Hi Joel"

	if got != wanted {
		t.Fatalf("got: %s | wanted : %s", got, wanted)
	}
}
