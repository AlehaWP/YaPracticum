package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Url []byte

func (u Url) String() string {
	return string(u)
}

func (u Url) MD5() string {
	h := md5.Sum(u)
	return fmt.Sprintf("%x", h)
}

type Repository interface {
	GetURL(string) (Url, bool)
	SaveURL(Url) string
}

type UrlsData map[string]Url

func (u UrlsData) SaveURL(url Url) string {
	r := url.MD5()
	u[r] = url
	return r
}

func (u UrlsData) GetURL(id string) (Url, bool) {
	if r, ok := u[id]; ok {
		return r, true
	}
	return nil, false
}

type Server struct {
	http.Server
}

func (s Server) start(addr string, router func(http.ResponseWriter, *http.Request)) {
	s.Addr = addr
	s.Handler = http.HandlerFunc(router)
	go func() {
		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()
}

// func (s Server) stop() {
// 	if err := s.Shutdown(context.Background()); err != nil {
// 		log.Printf("HTTP server Shutdown: %v", err)
// 	}
// }

// func (s Server) listenChanToQuit(f func(chan bool)) chan bool {
// 	sigQuitChan := make(chan bool)
// 	go f(sigQuitChan)
// 	return sigQuitChan
// }

func Router(repo Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlerPost(w, r, repo)
		case http.MethodGet:
			handlerGet(w, r, repo)
		default:
			w.WriteHeader(400)
		}
	}
}

func handlerPost(w http.ResponseWriter, r *http.Request, repo Repository) {
	textBody, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	retUrl := "http://" + r.Host + "/" + repo.SaveURL(Url(textBody))
	w.WriteHeader(201)
	io.WriteString(w, retUrl)
}

func handlerGet(w http.ResponseWriter, r *http.Request, repo Repository) {
	id := r.URL.Path[1:]
	if val, ok := repo.GetURL(id); ok {
		w.Header().Add("Location", val.String())
		w.WriteHeader(307)
	} else {
		w.WriteHeader(400)
	}
}

// func scanQuit(ch chan bool) {
// 	var inputText string
// 	for strings.ToLower(inputText) != "quit" {
// 		fmt.Println("For server stop please input: quit")
// 		fmt.Scanf("%s\n", &inputText)
// 	}
// 	ch <- true
// }

func main() {
	s := Server{http.Server{}}
	s.start("localhost:8082", Router(make(UrlsData)))
	// <-s.listenChanToQuit(scanQuit)
	// s.stop()
}
