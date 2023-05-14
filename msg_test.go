package msg_test

import (
	"testing"

	"github.com/FollowTheProcess/msg"
)

func TestHello(t *testing.T) {
	got := msg.Hello()
	want := "Hello msg"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
