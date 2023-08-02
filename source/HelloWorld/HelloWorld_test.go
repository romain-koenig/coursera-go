// Testing Hello World

package main

import "testing"

func TestGreeting(t *testing.T) {
	want := "Hello, World\n"
	got := Greeting()
	if got != want {
		t.Errorf("Greeting() = %q; want %q", got, want)
	}
}
