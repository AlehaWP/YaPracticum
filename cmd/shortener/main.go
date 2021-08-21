package main

import (
	"io"
	"fmt"
	"net/http"
	"crypto/md5"
)

var Urls map[string]string

func MD5(data []byte) string {
	h := md5.Sum(data)
	return fmt.Sprintf("%x", h)
}

func handlerPost (w http.ResponseWriter, r *http.Request){
	textBody, err := io.ReadAll(r.Body)
	if err != nil {

	}
	defer r.Body.Close()
	mdUrl := MD5(textBody)
	Urls[mdUrl] = string(textBody)
	w.WriteHeader(201)
	io.WriteString (w, mdUrl)
}

func handlerGet(w http.ResponseWriter, r *http.Request){
	id := r.URL.Path[1:]
	if val, ok := Urls[id]; ok {
		w.Header().Add("Location", val)
		w.WriteHeader(307)
	} else {
		w.WriteHeader(400)
	}
	w.Write([]byte(id))
}

func handlerFunction (w http.ResponseWriter, r *http.Request) {
	switch r.Method{
		case  http.MethodPost:
			handlerPost(w, r)
		case http.MethodGet:
			handlerGet(w, r)
	    default:
			w.WriteHeader(400)
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
