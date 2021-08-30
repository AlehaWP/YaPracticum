package server

import (
	"github.com/go-chi/chi/v5"
	"handlers"
)

type Server struct {
	http.Server
}

//Start Server with router
func (s *Server) start(addr string, repo Repository) {
	r := chi.NewRouter()
	r.Post("/", handlerUrlPost(repo))
	r.Route("/{id}", func(r chi.Router) {
		r.Use(UrlCtx)
		r.Get("/", handlerUrlGet(repo))
	})
	s.Addr = addr
	s.Handler = r
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}
