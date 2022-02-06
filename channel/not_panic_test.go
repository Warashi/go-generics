package channel_test

import (
	"testing"

	"github.com/Warashi/go-generics/channel"
)

func TestNotPanic(t *testing.T) {
	ch := channel.NoPanic(make(chan struct{}))
	ch.Close()
	ch.Close()
	ch.Send(struct{}{})
}
