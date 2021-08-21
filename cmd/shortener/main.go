package main

import (
	"io"
	"fmt"
	"net/http"
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
			io.WriteString (w, mdUrl)
		case http.MethodGet:
			id := r.URL.Path[1:]
			io.WriteString (w, Urls[id])
	    default:
			io.WriteString (w, "default")
	}
}


func main() {
	Urls =  make(map[string]string)
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlerFunction)
	server := &http.Server{
        Addr: "localhost:8080",
		Handler: mux,
    }
	server.ListenAndServe()
}
