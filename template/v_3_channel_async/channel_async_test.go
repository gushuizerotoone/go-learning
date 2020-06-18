package v_3_channel_async

import (
	"testing"
	"time"
)

func TestChannelAsync(t *testing.T) {
	server := NewServer()
	server.Process()

	go server.ListenToProcess()
	time.Sleep(time.Minute)
}
