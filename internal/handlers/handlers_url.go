package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/AlehaWP/YaPracticum.git/internal/global"
)

var Repo global.Repository
var BaseURL string
var Opt global.Options

func HandlerUserPostURLs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := ctx.Value(global.CtxString("UserID")).(string)

	ud := Repo.GetUserURLs(userID)

	if len(ud) == 0 {
		w.WriteHeader(204)
		return
	}

	res, err := json.Marshal(ud)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(res)
}

func HandlerCheckDBConnect(w http.ResponseWriter, r *http.Request) {
	if err := Repo.CheckDBConnection(); err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
}

// HandlerUrlPost saves url from request body to repository.
func HandlerURLPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := ctx.Value(global.CtxString("UserID")).(string)

	textBody, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(400)
		return
	}
	retURL := Repo.SaveURL(textBody, BaseURL+"/", userID)
	w.Header().Add("Content-Type", r.Header.Get("Content-Type"))
	w.WriteHeader(201)
	w.Write([]byte(retURL))
}

//HandlerAPIURLPost saves url from body request.
func HandlerAPIURLPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID := ctx.Value(global.CtxString("UserID")).(string)

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
		URLShorten: Repo.SaveURL([]byte(tURLJson.URLLong), BaseURL+"/", userID),
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

	id := ctx.Value(global.CtxString("url_id")).(string)
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

func NewHandlers(repo global.Repository, opt global.Options) {
	Repo = repo
	BaseURL = opt.RespBaseURL()
	Opt = opt
}
