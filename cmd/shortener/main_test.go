package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
	r := httptest.NewRequest("POST", "http://localhost:8082/", strings.NewReader("www.example.com"))
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, r)
	res := w.Result()
	b, _ := io.ReadAll(res.Body)
	defer res.Body.Close()
	assert.Equal(t, 201, w.Result().StatusCode, "Не верный код ответа POST")
	assert.Equal(t, "http://localhost:8082/123123asdasd", string(b), "Не верный ответ POST")

	r = httptest.NewRequest("GET", "/123123asdasd", strings.NewReader(""))
	w = httptest.NewRecorder()
	handler.ServeHTTP(w, r)
	assert.Equal(t, 307, w.Result().StatusCode, "Не верный код ответа GET")
	assert.Equal(t, w.Header().Get("Location"), "www.example.com", "Не верный ответ GET")

	r = httptest.NewRequest("GET", "/123123", strings.NewReader(""))
	w = httptest.NewRecorder()
	handler.ServeHTTP(w, r)
	assert.Equal(t, 400, w.Result().StatusCode, "Не верный код ответа GET")
}
