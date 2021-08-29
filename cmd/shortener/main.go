package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"strings"
	"context"
	"log"
	"github.com/go-chi/chi/v5"
)


type Repository interface {
	GetURL(string) (string, bool)
	SaveURL([]byte) string
}

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


func MD5(b []byte) string {
	h := md5.Sum(b)
	return fmt.Sprintf("%x", h)
}

type Server struct {
	http.Server
}

func (s *Server) start(addr string, repo Repository) {
	r := chi.NewRouter()
	r.Post("/", handlerUrlPost(repo))
	r.Route("/{id}", func(r chi.Router){
		r.Use(UrlCtx)
		r.Get("/", handlerUrlGet(repo))
	})
	// r.Get("/{id}", handlerUrlGet(repo))
	s.Addr = addr
	s.Handler = r
	go func() {
		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()
	// s.ListenAndServe()
}


func (s *Server) listenChanToQuit(f func(chan bool)) {
	sigQuitChan := make(chan bool)
	go f(sigQuitChan)
	<-sigQuitChan
	s.stop()
}

func (s *Server) stop() {
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("HTTP server Shutdown: %v", err)
	}
}



func handlerUrlPost(repo Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		textBody, _ := io.ReadAll(r.Body)
		defer r.Body.Close()
		retUrl := "http://" + r.Host + "/" + repo.SaveURL(textBody)
		w.WriteHeader(201)
		io.WriteString(w, retUrl)
	}
}

func UrlCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	  ctx := context.WithValue(r.Context(), "id", chi.URLParam(r, "id"))
	  next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func handlerUrlGet(repo Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := ctx.Value("id").(string)
		if val, ok := repo.GetURL(id); ok {
			log.Println(val)
			w.Header().Add("Location", val)
			w.WriteHeader(307)
		} else {
			w.WriteHeader(400)
		}
	}
}

func scanQuit(ch chan bool) {
	var inputText string
	for strings.ToLower(inputText) != "quit" {
		fmt.Println("For server stop please input: quit")
		fmt.Scanf("%s\n", &inputText)
	}
	ch <- true
}

func main() {
	s := new(Server)
	urlData := make(UrlsData)
	s.start("localhost:8080", &urlData)
	s.listenChanToQuit(scanQuit)
}
