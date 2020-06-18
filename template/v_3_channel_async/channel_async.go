package v_3_channel_async

import (
	"fmt"
	"time"
)

// this template not finish yet
type Server struct {
	listenChannel chan string
}

func NewServer() *Server {
	listenChannel := make(chan string, 2)
	return &Server{listenChannel: listenChannel}
}

func (s *Server) Process() {
	str := "first"
	s.listenChannel <- str

	time.Sleep(time.Second * 10)
	str = "second"
	s.listenChannel <- str
}

func (s *Server) ListenToProcess() {
	for str := range s.listenChannel {
		fmt.Printf("%v: %s", time.Now(), str)
	}
}
