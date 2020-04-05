package hello

import "testing"

func TestReturnGreeting(t *testing.T) {
	got := ReturnGreeting("Hi there")
	want := "Hi there yourself!"
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
