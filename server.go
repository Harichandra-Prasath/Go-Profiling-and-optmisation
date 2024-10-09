package main

import (
	"net/http"
)

type Server struct {
	Addr string
	Mux  *http.ServeMux
}

func NewServer(Addr string) *Server {
	return &Server{
		Addr: Addr,
		Mux:  http.NewServeMux(),
	}
}

func (S *Server) Serve() error {

	err := http.ListenAndServe(S.Addr, S.Mux)
	if err != nil {
		return err
	}
	return nil
}
