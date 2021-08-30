package handlers

import (
	"context"
	"io"
	"log"
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

func (m *UrlsMock) SaveURL(url []byte) string {
	args := m.Called(url)
	return args.String(0)
}

func (m *UrlsMock) GetURL(id string) (string, bool) {
	args := m.Called(id)
	return args.String(0), args.Bool(1)
}

func TestHandlerUrlGet(t *testing.T) {
	dataTests := map[string]map[string]interface{}{
		"test1": {
			"reqID":       "123123asdasd",
			"result":      "www.example.com",
			"resStatus":   307,
			"mockReturn1": "www.example.com",
			"mockReturn2": true,
		},
		"test2": {
			"reqID":       "123123",
			"result":      "",
			"resStatus":   400,
			"mockReturn1": "",
			"mockReturn2": false,
		},
	}

	repoMock := new(UrlsMock)
	handler := http.HandlerFunc(HandlerURLGet(repoMock))

	for key, value := range dataTests {
		log.Println("start test", key)
		repoMock.On("GetURL", value["reqID"].(string)).Return(value["mockReturn1"].(string), value["mockReturn2"].(bool))
		r := httptest.NewRequest("GET", "/"+value["reqID"].(string), strings.NewReader(""))
		w := httptest.NewRecorder()
		ctx := context.WithValue(context.Background(), Key("id"), value["reqID"].(string))
		handler.ServeHTTP(w, r.WithContext(ctx))
		res := w.Result()
		assert.Equal(t, value["resStatus"].(int), res.StatusCode, "Не верный код ответа GET")
		assert.Equal(t, w.Header().Get("Location"), value["result"].(string), "Не верный ответ GET")
		defer res.Body.Close()
	}
}

func TestHandlerUrlPost(t *testing.T) {
	repoMock := new(UrlsMock)
	repoMock.On("SaveURL", []byte("www.example.com")).Return("123123asdasd")
	handler := http.HandlerFunc(HandlerURLPost(repoMock))
	r := httptest.NewRequest("POST", "http://localhost:8082/", strings.NewReader("www.example.com"))
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, r)
	res := w.Result()
	b, _ := io.ReadAll(res.Body)
	assert.Equal(t, 201, res.StatusCode, "Не верный код ответа POST")
	assert.Equal(t, "http://localhost:8082/123123asdasd", string(b), "Не верный ответ POST")

	defer res.Body.Close()
}
