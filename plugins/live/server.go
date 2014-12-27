package live

import "encoding/json"
import "sync/atomic"
import "net"
import "log"
import "fmt"
import "os"

// Server accepts connections to a trace socket for remote debugging.
type Server struct {
	name    string
	log     *log.Logger
	events  chan interface{}
	clients int32
}

// New server for program `name` and
// listen for connections in a new goroutine.
func New(name string) *Server {
	s := NewServer(name)
	go func() {
		s.log.Fatalf("failed to bind server: %s", s.Listen())
	}()
	return s
}

// NewServer for program `name`.
func NewServer(name string) *Server {
	return &Server{
		name:   name,
		log:    log.New(os.Stderr, "[trace] ", log.LstdFlags),
		events: make(chan interface{}, 1000),
	}
}

// Listen and start accepting connections.
func (s *Server) Listen() error {
	os.Remove(s.Path())

	lsock, err := net.Listen("unix", s.Path())
	if err != nil {
		return err
	}
	defer lsock.Close()

	for {
		sock, err := lsock.Accept()
		if err != nil {
			s.log.Printf("failed to accept: %s", err)
			continue
		}

		go s.chat(sock)
	}
}

// Relay events to `sock` as JSON.
func (s *Server) chat(sock net.Conn) {
	s.log.Printf("client connect")
	defer s.log.Printf("client disconnect")

	enc := json.NewEncoder(sock)
	defer sock.Close()

	atomic.AddInt32(&s.clients, 1)
	defer atomic.AddInt32(&s.clients, -1)

	for {
		select {
		case e := <-s.events:
			err := enc.Encode(e)

			if err != nil {
				s.log.Printf("failed to write: %s", err)
				return
			}
		}
	}
}

// Path to socket.
func (s *Server) Path() string {
	return fmt.Sprintf("/tmp/%s", s.name)
}

// HandleEvent implementation.
func (s *Server) HandleEvent(e interface{}) {
	if atomic.LoadInt32(&s.clients) > 0 {
		s.events <- e
	}
}
