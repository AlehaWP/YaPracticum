package server

import (
	"net/http"

	"github.com/AlehaWP/YaPracticum.git/internal/global"
	"github.com/AlehaWP/YaPracticum.git/internal/handlers"
	"github.com/AlehaWP/YaPracticum.git/internal/middlewares"
	"github.com/AlehaWP/YaPracticum.git/internal/zip"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	http.Server
}

//Start server with router.
func (s *Server) Start(repo global.Repository, opt global.Options) {
	r := chi.NewRouter()
	baseURL := opt.RespBaseURL()
	r.Post("/", zip.ZipHandlerRead(zip.ZipHandlerWrite(handlers.HandlerURLPost(repo, baseURL))))
	r.Route("/{id}", func(r chi.Router) {
		r.Use(middlewares.URLCtx)
		r.Get("/", handlers.HandlerURLGet(repo))
	})
	r.Post("/api/shorten", zip.ZipHandlerRead(zip.ZipHandlerWrite(handlers.HandlerAPIURLPost(repo, baseURL))))

	s.Addr = opt.ServAddr()
	s.Handler = r
	s.ListenAndServe()
}
