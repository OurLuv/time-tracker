package handler

import (
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func (s *Server) Start(port string, h http.Handler) error {
	s.server = &http.Server{
		Addr:           port,
		Handler:        h,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
