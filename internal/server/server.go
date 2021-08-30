package server

import (
	"log"
	"net/http"

	"github.com/AlehaWP/YaPracticum.git/internal/handlers"
	"github.com/AlehaWP/YaPracticum.git/internal/middlewares"
	"github.com/AlehaWP/YaPracticum.git/internal/repository"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	http.Server
}

//Start Server with router
func (s *Server) Start(addr string, repo repository.Repository) {
	r := chi.NewRouter()
	r.Post("/", handlers.HandlerURLPost(repo))
	r.Route("/{id}", func(r chi.Router) {
		r.Use(middlewares.URLCtx)
		r.Get("/", handlers.HandlerURLGet(repo))
	})
	s.Addr = addr
	s.Handler = r
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}
