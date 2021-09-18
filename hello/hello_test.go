package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello There!"
	got := Hello()

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
