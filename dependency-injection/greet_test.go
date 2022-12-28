package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	var buf bytes.Buffer
	Greet(&buf, "Chris")

	got := buf.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
