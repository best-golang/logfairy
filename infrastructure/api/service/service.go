package service

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

// Client is a http client
type Service struct {
	BaseURL string
	client  http.Client
}

type serviceRequest struct {
	method  string
	url     string
	headers map[string]string
	auth    []string
	body    io.Reader
}

// New create an instance of http service
func New(BaseURL string, timeout int) Service {
	connectionTimeout := time.Duration(time.Duration(timeout) * time.Second)

	return Service{
		BaseURL: BaseURL,
		client: http.Client{
			Transport: &http.Transport{
				Dial: TimeoutDialer(connectionTimeout),
			},
		},
	}
}

// Get execute a http request with GET method
func (service Service) Get(
	endpoint string,
	headers map[string]string,
	auth []string,
	body io.Reader,
) (
	[]byte,
	int,
	error,
) {
	request := serviceRequest{
		method:  http.MethodGet,
		url:     service.buildURL(endpoint),
		headers: headers,
		auth:    auth,
		body:    nil,
	}

	return service.do(request)
}

// Post execute a http request with POST method
func (service Service) Post(
	endpoint string,
	headers map[string]string,
	auth []string,
	body io.Reader,
) (
	[]byte,
	int,
	error,
) {
	request := serviceRequest{
		method:  http.MethodPost,
		url:     service.buildURL(endpoint),
		headers: headers,
		auth:    auth,
		body:    body,
	}

	return service.do(request)
}

// Post execute a http request with POST method
func (service Service) Put(
	endpoint string,
	headers map[string]string,
	auth []string,
	body io.Reader,
) (
	[]byte,
	int,
	error,
) {
	request := serviceRequest{
		method:  http.MethodPut,
		url:     service.buildURL(endpoint),
		headers: headers,
		auth:    auth,
		body:    body,
	}

	return service.do(request)
}

// Get execute a http request with GET method
func (service Service) do(request serviceRequest) ([]byte, int, error) {
	req, err := http.NewRequest(request.method, request.url, request.body)
	if err != nil {
		return nil, 0, err
	}

	if len(request.auth) == 2 {
		req.SetBasicAuth(request.auth[0], request.auth[1])
	}

	for header, value := range request.headers {
		req.Header.Add(header, value)
	}

	resp, err := service.client.Do(req)
	if err != nil {
		return nil, 0, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return body, resp.StatusCode, nil
}

func (service Service) buildURL(endpoint string) string {
	return service.BaseURL + endpoint
}

func TimeoutDialer(cTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, cTimeout)
		if err != nil {
			return nil, fmt.Errorf("was not possible to reach the endpoint: %s", err)
		}

		return conn, nil
	}
}
