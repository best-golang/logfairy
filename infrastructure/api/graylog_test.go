package api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uniplaces/logfairy/infrastructure/api"
	"github.com/uniplaces/logfairy/infrastructure/api/service/mocks"
)

const (
	username = "tester"
	password = "123456"
)

func TestCreation(t *testing.T) {
	t.Parallel()

	client := &mocks.Client{}
	graylog := api.New(client, username, password)

	assert.NotNil(t, graylog)

	client.AssertExpectations(t)
}

func TestGetAuth(t *testing.T) {
	t.Parallel()

	endpoint := "/api/users/tester/tokens/tests"
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	auth := []string{username, password}

	client := &mocks.Client{}
	client.On(
		"Post",
		endpoint,
		headers,
		auth,
		nil,
	).Return(
		[]byte("{\"name\": \"tests\", \"token\": \"abcde123456\", \"last_access\": \"\"}"),
		200,
		nil,
	).Once()
	graylog := api.New(client, username, password)

	tokenAuth, err := graylog.GetAuth("tests")

	assert.Nil(t, err)
	assert.Len(t, tokenAuth, 2)
	client.AssertExpectations(t)
}
