
package server

import (
	"log"
	"sync"
	"time"
)

type Server struct {
	Mu         sync.Mutex
	Data       map[string]string
	Requests   int
	ShutdownCh chan struct{}
}

func NewServer() *Server {
	return &Server{
		Data:       make(map[string]string),
		ShutdownCh: make(chan struct{}),
	}
}

func (s *Server) StartBackgroundWorker() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			s.Mu.Lock()
			log.Printf("Status: %d requests, %d items in database\n", s.Requests, len(s.Data))
			s.Mu.Unlock()
		case <-s.ShutdownCh:
			log.Println("Background worker shutting down.")
			return
		}
	}
}

func (s *Server) Shutdown() {
	close(s.ShutdownCh)
}
