package main

import (
	"net/http"
)

func (S *Server) SetRoutes() {
	Mux := S.Handler.(*http.ServeMux)
	Mux.Handle("/check", pingHandler())
}

func pingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Check Completed!!\n"))
	}
}
