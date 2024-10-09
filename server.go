package main

import (
	"net/http"
)

type Server struct {
	Addr    string
	Handler http.Handler
}

func NewServer(Addr string) *Server {
	return &Server{
		Addr:    Addr,
		Handler: http.NewServeMux(),
	}
}

func (S *Server) Serve() error {
	Logger.Info("Server Started at Port", S.Addr)

	// Before Starting the Server, use the wrapper
	S.Handler = NewLatencyWrapper(S.Handler)

	err := http.ListenAndServe(S.Addr, S.Handler)
	if err != nil {
		return err
	}
	return nil
}
