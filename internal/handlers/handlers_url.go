package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/AlehaWP/YaPracticum.git/internal/global"
)

var Repo global.Repository
var BaseURL string

// HandlerUrlPost saves url from request body to repository.
func HandlerURLPost(w http.ResponseWriter, r *http.Request) {
	textBody, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(400)
		return
	}
	retURL := BaseURL + "/" + Repo.SaveURL(textBody, "")
	w.Header().Add("Content-Type", r.Header.Get("Content-Type"))
	w.WriteHeader(201)
	w.Write([]byte(retURL))

}

//HandlerAPIURLPost saves url from body request.
func HandlerAPIURLPost(w http.ResponseWriter, r *http.Request) {
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
		URLShorten: BaseURL + "/" + Repo.SaveURL([]byte(tURLJson.URLLong), ""),
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

// HandlerUrlGet returns url from repository to resp.Head - "Location".
func HandlerURLGet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := ctx.Value("url_id").(string)
	val, err := Repo.GetURL(id)
	if err != nil {
		w.WriteHeader(400)
		io.WriteString(w, err.Error())
		return
	}
	w.Header().Add("Location", val)
	w.Header().Add("Content-Type", r.Header.Get("Content-Type"))
	w.WriteHeader(307)
}

func NewHandlers(repo global.Repository, baseURL string) {
	Repo = repo
	BaseURL = baseURL
}
