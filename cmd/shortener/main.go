package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Url []byte

func (u Url) String() string {
	return string(u)
}

func (u Url) MD5() string {
	h := md5.Sum(u)
	return fmt.Sprintf("%x", h)
}

type Urls map[string]Url

func (u Urls) SaveURL(url Url) string {
	r := url.MD5()
	u[r] = url
	return r
}

func (u Urls) GetURL(id string) (Url, bool){
	if r, ok := u[id]; ok {
		return r, true
	}
	return nil, false
}



var urlsData Urls

func handlerPost(w http.ResponseWriter, r *http.Request) {
	textBody, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	retUrl := "http://" + r.Host + "/" + urlsData.SaveURL(textBody)
	w.WriteHeader(201)
	io.WriteString(w, retUrl)
}

func handlerGet(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[1:]
	if val, ok := urlsData.GetURL(id); ok {
		w.Header().Add("Location", val.String())
		w.WriteHeader(307)
	} else {
		w.WriteHeader(400)
	}
}

func Router(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handlerPost(w, r)
	case http.MethodGet:
		handlerGet(w, r)
	default:
		w.WriteHeader(400)
	}
}

type Server struct {
	http.Server
}

func (s *Server) start(addr string, router func(http.ResponseWriter, *http.Request)){
	mux := http.NewServeMux()
	mux.HandleFunc("/", router)
	s.Addr = addr
	s.Handler = mux
	go func() {
		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()
}

func (s *Server) stop() {
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("HTTP server Shutdown: %v", err)
	}
}

func (s *Server) listenChanToQuit(f func(chan bool)) {
	sigQuitChan := make(chan bool)
	go f(sigQuitChan)
	<-sigQuitChan
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
	urlsData = make(Urls)
	s := Server{http.Server{}}
	s.start("localhost:8080", Router)
	s.listenChanToQuit(scanQuit)
    s.stop()
}
