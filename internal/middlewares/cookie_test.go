package middlewares

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/AlehaWP/YaPracticum.git/internal/test"
	"github.com/stretchr/testify/assert"
)

var repoMock *test.RepoMock
var once sync.Once

func nextHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Моя домашняя страница!"))
}

func TestSetCookieUser(t *testing.T) {
	once.Do(InitMocks)

	type args struct {
		userID  string
		finded  bool
		cUserID string
		cErr    error
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Тестй",
			args{
				"1231323434",
				true,
				"",
				nil,
			},
			"1231323434",
		},
	}

	for _, tt := range tests {
		r := httptest.NewRequest("GET", "/", strings.NewReader(""))
		w := httptest.NewRecorder()

		ctx := context.Background()
		r = r.WithContext(ctx)
		repoMock.On("FindUser", ctx, tt.args.userID).Return(tt.args.finded)
		repoMock.On("CreateUser", ctx, tt.args.userID).Return(tt.args.cUserID, tt.args.cErr)

		NewCookie(repoMock)

		c := &http.Cookie{
			Name:  "UserID",
			Value: tt.args.userID,
		}
		r.AddCookie(c)

		handler := SetCookieUser(http.HandlerFunc(nextHandler))
		handler.ServeHTTP(w, r)
		res := w.Result()

		assert.Equal(t, tt.want, res.Cookies()[0].Value, "Не верный ответ POST")
	}
}

func InitMocks() {
	repoMock = new(test.RepoMock)
}
