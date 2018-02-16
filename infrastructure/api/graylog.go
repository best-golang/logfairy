package api

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/uniplaces/logfairy/infrastructure/api/dto"
	"github.com/uniplaces/logfairy/infrastructure/api/service"
)

// endpoint constants
const (
	// users/<username>/tokens/<token_name>
	tokensStream = "/api/users/%s/tokens/%s"
)

var headers = map[string]string{
	"Content-Type": "application/json",
}

type Graylog struct {
	client   service.Service
	username string
	password string
	token    string
}

// New create an instance of Graylog api client
func New(client service.Service, username string, password string) Graylog {
	return Graylog{
		client:   client,
		username: username,
		password: password,
		token:    "",
	}
}

func (graylog *Graylog) getAuth(tokenName string) ([]string, error) {
	token, err := graylog.getToken(tokenName)
	if err != nil {
		return []string{}, err
	}

	return []string{token, "token"}, nil
}

func (graylog Graylog) getToken(tokenName string) (string, error) {
	if graylog.token != "" {
		return graylog.token, nil
	}

	if err := graylog.setToken(tokenName); err != nil {
		return "", err
	}

	return graylog.token, nil
}

func (graylog *Graylog) setToken(tokenName string) error {
	auth := []string{graylog.username, graylog.password}

	endpoint := fmt.Sprintf(tokensStream, graylog.username, tokenName)
	response, status, err := graylog.client.Post(endpoint, headers, auth, nil)
	if err != nil {
		return err
	}

	failure, err := graylog.handleFailure(response, status)
	if err != nil {
		return err
	}

	if failure != nil {
		return errors.New(failure.Message)
	}

	tokenResponse := &dto.TokenResponse{}
	if err := json.Unmarshal(response, tokenResponse); err != nil {
		return err
	}

	graylog.token = tokenResponse.Token

	return nil
}

func (graylog *Graylog) handleFailure(response []byte, status int) (*dto.ErrorResponse, error) {
	if status >= 200 && status < 299 {
		return nil, nil
	}

	failure := dto.ErrorResponse{}
	if err := json.Unmarshal(response, &failure); err != nil {
		return nil, err
	}

	return &failure, nil
}
