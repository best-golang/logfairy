package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/uniplaces/logfairy/infrastructure/api/dto"
	"github.com/uniplaces/logfairy/infrastructure/api/service"
)

// endpoint constants
const (
	// /api/users/<username>/tokens/<token_name>
	tokensStream = "/api/users/%s/tokens/%s"
)

const (
	authType = "token"
)

var headers = map[string]string{
	"Content-Type": "application/json",
}

type Graylog struct {
	Client   service.Service
	username string
	password string
	token    string
}

// New create an instance of Graylog api client
func New(client service.Service, username string, password string) Graylog {
	return Graylog{
		Client:   client,
		username: username,
		password: password,
		token:    "",
	}
}

// GetAuth get the token to use for basic authentincation
func (graylog *Graylog) GetAuth(tokenName string) ([]string, error) {
	token, err := graylog.getToken(tokenName)
	if err != nil {
		return []string{}, err
	}

	return []string{token, authType}, nil
}

// HandleFailure handle the response in order to handle failures
func (graylog *Graylog) HandleFailure(response []byte, status int) error {
	if status >= http.StatusOK && status < http.StatusMultipleChoices {
		return nil
	}

	failure := dto.ErrorResponse{}
	if err := json.Unmarshal(response, &failure); err != nil {
		return err
	}

	if failure.Message != "" {
		return errors.New(failure.Message)
	}

	return nil
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
	response, status, err := graylog.Client.Post(endpoint, headers, auth, nil)
	if err != nil {
		return err
	}

	if err := graylog.HandleFailure(response, status); err != nil {
		return err
	}

	tokenResponse := &dto.TokenResponse{}
	if err := json.Unmarshal(response, tokenResponse); err != nil {
		return err
	}

	graylog.token = tokenResponse.Token

	return nil
}
