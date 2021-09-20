package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/AlehaWP/YaPracticum.git/internal/global"
)

var Repo global.Repository

func SetCookieUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("user_id")
		cv := ""
		if err == nil {
			cv = c.Value
		}
		if ok := Repo.FindUser(cv); !ok {
			cv, err = Repo.CreateUser()
			if err != nil {
				fmt.Println("Can't create cookie", err)
				next.ServeHTTP(w, r)
				return
			}
		}

		c = &http.Cookie{
			Name:  "user_id",
			Value: cv,
		}
		http.SetCookie(w, c)

		ctx := context.WithValue(r.Context(), global.CtxString("userId"), cv)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func NewCookie(repo global.Repository) {
	Repo = repo
}
