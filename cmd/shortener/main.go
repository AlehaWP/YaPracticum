package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"context"
	"log"
	"github.com/go-chi/chi/v5"
)

//Repository Interface bd urls
type Repository interface {
	GetURL(string) (string, bool)
	SaveURL([]byte) string
}
//UrlsData Repository of urls. Realize Repository interface
type UrlsData map[string][]byte

func (u *UrlsData) SaveURL(url []byte) string {
	r := MD5(url)
	(*u)[r] = url
	return r
}

func (u *UrlsData) GetURL(id string) (string, bool) {
	if r, ok := (*u)[id]; ok {
		return string(r), true
	}
	return "", false
}

//MD5 func for hash.
func MD5(b []byte) string {
	h := md5.Sum(b)
	return fmt.Sprintf("%x", h)
}

type Server struct {
	http.Server
}

//Start Server with router
func (s *Server) start(addr string, repo Repository) {
	r := chi.NewRouter()
	r.Post("/", handlerUrlPost(repo))
	r.Route("/{id}", func(r chi.Router){
		r.Use(UrlCtx)
		r.Get("/", handlerUrlGet(repo))
	})
	s.Addr = addr
	s.Handler = r
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}


// handlerUrlPost Saves url from request body to repository
func handlerUrlPost(repo Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		textBody, _ := io.ReadAll(r.Body)
		defer r.Body.Close()
		retUrl := "http://" + r.Host + "/" + repo.SaveURL(textBody)
		w.WriteHeader(201)
		io.WriteString(w, retUrl)
	}
}

// UrlCtx for parameter transfer without direct access to router
func UrlCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	  ctx := context.WithValue(r.Context(), "id", chi.URLParam(r, "id"))
	  next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// handlerUrlGet Returns url from repository to resp.Head - "Location"
func handlerUrlGet(repo Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := ctx.Value("id").(string)
		if val, ok := repo.GetURL(id); ok {
			w.Header().Add("Location", val)
			w.WriteHeader(307)
		} else {
			w.WriteHeader(400)
		}
	}
}

func main() {
	s := new(Server)
	urlData := make(UrlsData)
	s.start("localhost:8080", &urlData)
}
