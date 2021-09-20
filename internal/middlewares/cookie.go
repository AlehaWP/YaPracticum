package middlewares

import (
	"fmt"
	"net/http"

	"github.com/AlehaWP/YaPracticum.git/internal/repository"
)

func SetCookieUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("UserID")
		cv := ""
		if err == nil {
			cv = c.Value
		}
		if _, ok := repository.FindUser(cv); !ok {
			cv, err = repository.CreateUser()
		}
		if err != nil {
			fmt.Println("Can't create cookie", err)
			next.ServeHTTP(w, r)
			return
		}

		fmt.Println(cv)
		c.Name = "UserID"
		c.Value = cv
		http.SetCookie(w, c)
		next.ServeHTTP(w, r)
	}
}
