package main

import (
	"io"
	"fmt"
	"net/http"
	"time"
	_"os"
	"context"
	"strings"
)

func handlerFunction (w http.ResponseWriter, r *http.Request) {
	switch r.Method{
		case  http.MethodPost:
			io.WriteString (w, "Post")
		case http.MethodGet:
			io.WriteString (w, "Get " + r.URL.Path)
	    default:
			io.WriteString (w, "default")
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
	 mux := http.NewServeMux()
	// handler := http.HandlerFunc(handlerFunction)
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
