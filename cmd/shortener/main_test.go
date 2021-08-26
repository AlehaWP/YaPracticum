package main

import (
	"githubcom/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type UrlsMock struct {
	mock.Mock
}

func (m UrlsMock) SaveURL(url Url) string {
	args := m.Called(url)
	return args.String(0)
}

func (m UrlsMock) GetURL(id string) (Url, bool) {
	args := m.Called(id)
	return args.Get(0).(Url), args.Bool(1)
}

func TestRouter(t *testing.T) {
	repoMock := UrlsMock{}
	repoMock.On("GetURL", "123123asdasd").Return(Url("www.example.com"), true)
	repoMock.On("GetURL", "123123").Return(Url(""), false)
	repoMock.On("SaveURL", Url("www.example.com")).Return("123123asdasd")

	handler := http.HandlerFunc(Router(repoMock))
	r := httptest.NewRequest("POST", "/", strings.NewReader("www.example.com"))
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, r)

	t.Errorf("%v", w.Result().StatusCode)
	r = httptest.NewRequest("GET", "/123123asdasd", strings.NewReader(""))
	handler.ServeHTTP(w, r)
	t.Errorf("%v", w.Header().Get("Location"))
}
