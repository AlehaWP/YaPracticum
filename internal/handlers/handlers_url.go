package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/AlehaWP/YaPracticum.git/internal/repository"
)

// HandlerUrlPost saves url from request body to repository.
func HandlerURLPost(repo repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		textBody, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(400)
			return
		}
		retURL := "http://" + r.Host + "/" + repo.SaveURL(textBody)
		w.WriteHeader(201)
		io.WriteString(w, retURL)

	}
}

func HandlerApiURLPost(repo repository.Repository) http.HandlerFunc {
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
		tResJson := &struct {
			URLShorten string `json:"result"`
		}{
			URLShorten: "http://" + r.Host + "/" + repo.SaveURL([]byte(tURLJson.URLLong)),
		}
		res, err := json.Marshal(tResJson)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		w.WriteHeader(201)
		w.Write(res)
	}
}

// HandlerUrlGet returns url from repository to resp.Head - "Location".
func HandlerURLGet(repo repository.Repository) http.HandlerFunc {
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
		w.WriteHeader(307)
	}
}
