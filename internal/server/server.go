package server

import (
	"context"
	"fmt"
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/AlehaWP/YaPracticum.git/internal/handlers"
	"github.com/AlehaWP/YaPracticum.git/internal/middlewares"
	"github.com/AlehaWP/YaPracticum.git/internal/models"
	"github.com/AlehaWP/YaPracticum.git/internal/repository"
	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/acme/autocert"
)

type Server struct {
	http.Server
	opt  models.Options
	repo models.Repository
}

func (s *Server) newChiRouter() *chi.Mux {
	r := chi.NewRouter()
	handlers.NewHandlers(s.repo, s.opt)
	middlewares.NewCookie(s.repo)
	r.Use(middlewares.SetCookieUser, middlewares.ZipHandlerRead, middlewares.ZipHandlerWrite)

	r.HandleFunc("/debug/pprof/*", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	r.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	r.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	r.Handle("/debug/pprof/mutex", pprof.Handler("mutex"))
	r.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	r.Handle("/debug/pprof/block", pprof.Handler("block"))
	r.Handle("/debug/pprof/allocs", pprof.Handler("allocs"))
	// r.Use(middlewares.ZipHandlerRead, middlewares.ZipHandlerWrite)

	r.Get("/api/user/urls", handlers.HandlerUserPostURLs)
	r.Get("/ping", handlers.HandlerCheckDBConnect)
	r.Get("/api/internal/stats", handlers.HandlerReturnStats)
	r.Route("/{id}", func(r chi.Router) {
		r.Use(middlewares.URLCtx)
		r.Get("/", handlers.HandlerURLGet)
	})
	r.Post("/", handlers.HandlerURLPost)
	r.Post("/api/shorten", handlers.HandlerAPIURLPost)
	r.Post("/api/shorten/batch", handlers.HandlerAPIURLsPost)
	r.Delete("/api/user/urls", handlers.HandlerReturnStats)
	return r
}

//Start server with router.
func (s *Server) Start(ctx context.Context, opt models.Options) {
	sr, err := repository.NewServerRepo(ctx, opt.DBConnString())
	if err != nil {
		fmt.Println("Ошибка при подключении к БД: ", err)
		return
	}
	defer sr.Close()
	s.opt = opt
	s.repo = sr
	s.Handler = s.newChiRouter()
	s.Addr = opt.ServAddr()

	if s.opt.HTTPS() {
		manager := &autocert.Manager{
			Cache:  autocert.DirCache("cache-dir"),
			Prompt: autocert.AcceptTOS,
			// HostPolicy: autocert.HostWhitelist(s.opt.ServAddr()),
		}

		s.TLSConfig = manager.TLSConfig()
		go s.ListenAndServeTLS("", "")
	}

	if !s.opt.HTTPS() {
		go s.ListenAndServe()
	}

	fmt.Println("сервер HTTP начал работу")

	<-ctx.Done()
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	s.Shutdown(ctx)
}
