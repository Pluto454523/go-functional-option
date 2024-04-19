package main

import (
	"fmt"
)

type Server struct {
	Addr    string
	Port    int
	Timeout int
}

type ServerOption func(*Server)

func WithAddr(addr string) ServerOption {
	return func(s *Server) {
		s.Addr = addr
	}
}

func WithPort(port int) ServerOption {
	return func(s *Server) {
		s.Port = port
	}
}

func WithTimeout(timeout int) ServerOption {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

// ...ServerOption เรียกว่า Variadic Parameters
// คือสามารถส่ง parameter ที่เป็น type ServerOption ได้มากกว่า 1 parameter
func NewServer(options ...ServerOption) *Server {
	server := &Server{
		Addr:    "localhost",
		Port:    8080,
		Timeout: 30,
	}

	for _, option := range options {
		option(server)
	}

	return server
}

func main() {

	server1 := NewServer()
	fmt.Println(server1)

	server2 := NewServer(
		WithAddr("127.0.0.1"),
		WithPort(9090),
		WithTimeout(60),
	)

	fmt.Println(server2)
}
