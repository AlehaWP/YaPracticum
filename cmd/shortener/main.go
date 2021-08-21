package main

import (
	"io"
	"fmt"
	"net/http"
	"time"
	_"os"
	"context"
	"strings"
	"crypto/md5"
)

var Urls map[string]string // =  make(map[int]string)

func MD5(data []byte) string {
	h := md5.Sum(data)
	return fmt.Sprintf("%x", h)
}

func handlerFunction (w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	switch r.Method{
		case  http.MethodPost:
			textBody, err := io.ReadAll(r.Body)
			if err != nil {

			}
			defer r.Body.Close()

			url = string(textBody)
			mdUrl := MD5(textBody)
			Urls[mdUrl] = url
			w.WriteHeader(201)
			io.WriteString (w, mdUrl)
		case http.MethodGet:
			id := r.URL.Path[1:]
			if val, ok := Urls[id]; ok {
				w.Header().Add("Location", val)
				w.WriteHeader(307)
			} else {
				w.WriteHeader(500)
			}
			w.Write([]byte(id))
	    default:
			w.WriteHeader(404)
	}
}

func scanQuit(pSigChan chan bool){
	var inputText string
	for strings.ToLower(inputText) < "quit" {
		fmt.Println("Для остановки введите Quit")
		fmt.Scanf("%s\n", &inputText)
	}
	pSigChan <- true
}

func main() {
	Urls =  make(map[string]string)
	mux := http.NewServeMux()
	// handler := http.HandlerFunc(handlerFunction)
	// mux.Handle("/", handler)
	mux.HandleFunc("/", handlerFunction)
	server := &http.Server{
        Addr: "localhost:8080",
		Handler: mux,
    }
	go func(){
		server.ListenAndServe()
	}()

   	// sigChan := make(chan os.Signal, 1)
	// signal.Notify(sigChan, os.Interrapt)
	sigChan := make(chan bool)
	go scanQuit(sigChan)
	<- sigChan

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	server.Shutdown(ctx)

}
