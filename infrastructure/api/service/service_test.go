package service_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uniplaces/logfairy/infrastructure/api/service"
)

const (
	baseURL = "http://base-url.lh"
	timeout = 5
)

func TestCreation(t *testing.T) {
	t.Parallel()
	client := service.New(baseURL, timeout)

	assert.NotNil(t, client)
}

func TestGet(t *testing.T) {
	t.Parallel()

	var sendRouteCalled bool
	serverMux := http.NewServeMux()
	serverMux.HandleFunc(
		"/api/ping",
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)

			w.Write([]byte("{\"test\": true}"))
			w.WriteHeader(http.StatusOK)
			sendRouteCalled = true
		},
	)
	server := httptest.NewServer(serverMux)
	defer server.Close()

	client := service.New(server.URL, timeout)
	res, status, err := client.Get("/api/ping", map[string]string{}, []string{}, nil)

	assert.Nil(t, err)
	assert.Equal(t, 200, status)
	assert.Equal(t, []byte("{\"test\": true}"), res)
	assert.True(t, sendRouteCalled)
}

func TestPost(t *testing.T) {
	t.Parallel()

	var sendRouteCalled bool
	serverMux := http.NewServeMux()
	serverMux.HandleFunc(
		"/api/create",
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)

			w.Write([]byte("{\"test\": true}"))
			w.WriteHeader(http.StatusOK)
			sendRouteCalled = true
		},
	)
	server := httptest.NewServer(serverMux)
	defer server.Close()

	client := service.New(server.URL, timeout)
	res, status, err := client.Post("/api/create", map[string]string{}, []string{}, nil)

	assert.Nil(t, err)
	assert.Equal(t, 200, status)
	assert.Equal(t, []byte("{\"test\": true}"), res)
	assert.True(t, sendRouteCalled)
}

func TestPut(t *testing.T) {
	t.Parallel()

	var sendRouteCalled bool
	serverMux := http.NewServeMux()
	serverMux.HandleFunc(
		"/api/update",
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)

			w.Write([]byte("{\"test\": true}"))
			w.WriteHeader(http.StatusOK)
			sendRouteCalled = true
		},
	)
	server := httptest.NewServer(serverMux)
	defer server.Close()

	client := service.New(server.URL, timeout)
	res, status, err := client.Post("/api/update", map[string]string{}, []string{}, nil)

	assert.Nil(t, err)
	assert.Equal(t, 200, status)
	assert.Equal(t, []byte("{\"test\": true}"), res)
	assert.True(t, sendRouteCalled)
}

func TestGet_BadRequest(t *testing.T) {
	t.Parallel()

	var sendRouteCalled bool
	serverMux := http.NewServeMux()
	serverMux.HandleFunc(
		"/api/ping",
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
		},
	)
	server := httptest.NewServer(serverMux)
	defer server.Close()

	client := service.New(server.URL, timeout)
	res, status, err := client.Get("/api/pong", map[string]string{}, []string{}, nil)

	assert.Nil(t, err)
	assert.Equal(t, 404, status)
	assert.Equal(t, []byte("404 page not found\n"), res)
	assert.False(t, sendRouteCalled)
}
