package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

type Server struct {
	Clients map[string]net.Conn
}

func NewServer() *Server {
	return &Server{
		Clients: make(map[string]net.Conn),
	}
}

// TODO: handle closing connections
func (s *Server) Listen() {
	ln, err := net.Listen("tcp", ":7896")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go s.Handle(conn)
	}
}

func (s *Server) Handle(c net.Conn) {
	s.Add(c)

	reader := bufio.NewReader(c)
	for {
		line, _ := reader.ReadString('\n')
		s.Broadcast(line, c.RemoteAddr().String())
	}
}

func (s *Server) Broadcast(msg, addr string) {
	for caddr, conn := range s.Clients {
		if caddr == addr {
			continue
		}

		_, err := io.WriteString(conn, msg)
		if err != nil {
			log.Fatal("Unable to write")
		}
	}
}

// TODO: error if exists
func (s *Server) Add(c net.Conn) {
	addr := c.RemoteAddr().String()
	if _, exists := s.Clients[addr]; !exists {
		fmt.Printf("Connected: %s\n", addr)
		s.Clients[addr] = c
	} else {
		log.Fatal("conn already exists")
	}
}

func main() {
	s := NewServer()
	s.Listen()
}
