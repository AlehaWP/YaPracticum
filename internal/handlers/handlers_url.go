package handlers

import (
	"io"
	"net/http"

	"github.com/AlehaWP/YaPracticum.git/internal/repository"
)

// handlerUrlPost Saves url from request body to repository
func HandlerUrlPost(repo repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		textBody, _ := io.ReadAll(r.Body)
		defer r.Body.Close()
		retUrl := "http://" + r.Host + "/" + repo.SaveURL(textBody)
		w.WriteHeader(201)
		io.WriteString(w, retUrl)
	}
}

// handlerUrlGet Returns url from repository to resp.Head - "Location"
func HandlerUrlGet(repo repository.Repository) http.HandlerFunc {
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
