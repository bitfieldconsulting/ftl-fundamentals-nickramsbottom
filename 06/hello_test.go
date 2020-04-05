package hello

import "testing"

func TestHello(t *testing.T) {
	got := Greeting()
	want := "Howdy, Gopherinos!"
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
