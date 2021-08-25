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

type Urls map[string]Url

func (u Urls) SaveURL(url Url) string {
	r := u.MD5(url)
	u[r] = url
	return r
}

func (u Urls) GetURL(id string) string {
	return u[id].String()
}

func (u Urls) MD5(data Url) string {
	h := md5.Sum(data)
	return fmt.Sprintf("%x", h)
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
	if val, ok := urlsData[id]; ok {
		w.Header().Add("Location", GetURL(id))
		w.WriteHeader(307)
	} else {
		w.WriteHeader(400)
	}
}

func handlerFunction(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handlerPost(w, r)
	case http.MethodGet:
		handlerGet(w, r)
	default:
		w.WriteHeader(400)
	}
}

func startServer() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlerFunction)
	s := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	go func() {
		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()
	return s
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
	urlsData = make(map[string]Url)
	server := startServer()

	sigQuitChan := make(chan bool)
	go scanQuit(sigQuitChan)
	<-sigQuitChan

	if err := server.Shutdown(context.Background()); err != nil {
		log.Printf("HTTP server Shutdown: %v", err)
	}
}
