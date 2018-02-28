package api

import "io"

// Client defines the contracts to operate http requests
type Client interface {
	Get(endpoint string, headers map[string]string, auth []string, body io.Reader) ([]byte, int, error)
	Post(endpoint string, headers map[string]string, auth []string, body io.Reader) ([]byte, int, error)
	Put(endpoint string, headers map[string]string, auth []string, body io.Reader) ([]byte, int, error)
}
