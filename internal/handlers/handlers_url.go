package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/AlehaWP/YaPracticum.git/internal/global"
	"github.com/AlehaWP/YaPracticum.git/internal/repository"
)

// HandlerUrlPost saves url from request body to repository.
func HandlerURLPost(repo global.Repository, baseURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		textBody, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(400)
			return
		}
		retURL := baseURL + "/" + repo.SaveURL(textBody, "")
		w.Header().Add("Content-Type", r.Header.Get("Content-Type"))
		w.WriteHeader(201)
		w.Write([]byte(retURL))

	}
}

//HandlerAPIURLPost saves url from body request.
func HandlerAPIURLPost(repo global.Repository, baseURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tURLJson := &struct {
			URLLong string `json:"url"`
		}{}
		textBody, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(400)
			return
		}
		err = json.Unmarshal(textBody, tURLJson)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		tResJSON := &struct {
			URLShorten string `json:"result"`
		}{
			URLShorten: baseURL + "/" + repo.SaveURL([]byte(tURLJson.URLLong), ""),
		}

		res, err := json.Marshal(tResJSON)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		w.Header().Add("Content-Type", r.Header.Get("Content-Type"))
		w.WriteHeader(201)
		w.Write(res)
	}
}

// HandlerUrlGet returns url from repository to resp.Head - "Location".
func HandlerURLGet(repo global.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id := ctx.Value(repository.Key("id")).(string)
		val, err := repo.GetURL(id)
		if err != nil {
			w.WriteHeader(400)
			io.WriteString(w, err.Error())
			return
		}
		w.Header().Add("Location", val)
		w.Header().Add("Content-Type", r.Header.Get("Content-Type"))
		w.WriteHeader(307)
	}
}
