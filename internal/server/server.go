package server

import (
	"net/http"

	"github.com/AlehaWP/YaPracticum.git/internal/global"
	"github.com/AlehaWP/YaPracticum.git/internal/handlers"
	"github.com/AlehaWP/YaPracticum.git/internal/middlewares"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	http.Server
}

//Start server with router.
func (s *Server) Start(repo global.Repository, opt global.Options) {
	r := chi.NewRouter()
	handlers.NewHandlers(repo, opt.RespBaseURL())
	middlewares.NewCookie(repo)
	r.Use(middlewares.SetCookieUser, middlewares.ZipHandlerRead, middlewares.ZipHandlerWrite)
	//r.Use(middlewares.ZipHandlerRead, middlewares.ZipHandlerWrite)
	r.Post("/", handlers.HandlerURLPost)
	r.Get("/user/urls", handlers.HandlerUserPostURLs)
	r.Route("/{id}", func(r chi.Router) {
		r.Use(middlewares.URLCtx)
		r.Get("/", handlers.HandlerURLGet)
	})
	r.Post("/api/shorten", handlers.HandlerAPIURLPost)

	s.Addr = opt.ServAddr()
	s.Handler = r
	s.ListenAndServe()
}
