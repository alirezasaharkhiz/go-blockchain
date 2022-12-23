package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	port uint16
}

func NewServer(p uint16) *Server {
	return &Server{p}
}

func (s *Server) Port() uint16 {
	return s.port
}

func (s *Server) Run() {
	http.HandleFunc("/", Ping)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(s.port)), nil))
}

func Ping(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Pong!")
}
