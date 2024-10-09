package main

import (
	"net/http"
	"time"
)

type LatencyWrapper struct {
	Mux http.Handler
}

func (l *LatencyWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Logger.Debug(r.Method, r.URL.Path, r.UserAgent())
	t := time.Now()
	l.Mux.ServeHTTP(w, r)
	Logger.Debug("Time Taken to Complete the Request", time.Since(t).String())
}

func NewLatencyWrapper(Mux http.Handler) *LatencyWrapper {
	return &LatencyWrapper{
		Mux: Mux,
	}
}
