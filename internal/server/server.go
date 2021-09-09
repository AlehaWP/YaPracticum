package server

import (
	"net/http"

	"github.com/AlehaWP/YaPracticum.git/internal/handlers"
	"github.com/AlehaWP/YaPracticum.git/internal/middlewares"
	"github.com/AlehaWP/YaPracticum.git/internal/projectenv"
	"github.com/AlehaWP/YaPracticum.git/internal/repository"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	http.Server
}

//Start server with router.
func (s *Server) Start(repo repository.Repository) {
	r := chi.NewRouter()
	r.Post("/", handlers.HandlerURLPost(repo))
	//тут косяк
	r.Route("/{id}", func(r chi.Router) {
		r.Use(middlewares.URLCtx)
		r.Get("/", handlers.HandlerURLGet(repo))
	})
	r.Post("/api/shorten", handlers.HandlerAPIURLPost(repo))
	s.Addr = projectenv.Envs.ServAddr
	s.Handler = r
	s.ListenAndServe()
}
