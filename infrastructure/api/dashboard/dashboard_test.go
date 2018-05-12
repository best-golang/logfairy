package dashboard_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	dto "github.com/uniplaces/logfairy/dto/dashboard"
	"github.com/uniplaces/logfairy/infrastructure/api/dashboard"
	graylogmock "github.com/uniplaces/logfairy/infrastructure/api/mocks"
	servicemock "github.com/uniplaces/logfairy/infrastructure/api/service/mocks"
)

const (
	dashboardResponse = `{
		"title": "bar",
		"description": "description for bar dashboard"
	}`
)

func TestCreation(t *testing.T) {
	t.Parallel()

	client := &servicemock.Client{}
	graylog := &graylogmock.GraylogBase{}
	dashboardClient := dashboard.New(client, graylog)
	assert.NotNil(t, dashboardClient)

	client.AssertExpectations(t)
	graylog.AssertExpectations(t)
}

func TestGet(t *testing.T) {
	t.Parallel()
	endpoint := "/api/dashboards/dashboardID"
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	auth := []string{"abcde123456", "tests"}
	client := &servicemock.Client{}
	graylog := &graylogmock.GraylogBase{}

	graylog.On("GetAuth", "dashboards").Return(auth, nil).Once()
	graylog.On("HandleFailure", []byte(dashboardResponse), 200).Return(nil).Once()

	client.On("Get", endpoint, headers, auth, nil).Return([]byte(dashboardResponse), 200, nil).Once()

	dashboardClient := dashboard.New(client, graylog)
	assert.NotNil(t, dashboardClient)

	dashboardClient.Get("dashboardID")

	client.AssertExpectations(t)
	graylog.AssertExpectations(t)
}

func TestList(t *testing.T) {
	t.Parallel()
	endpoint := "/api/dashboards"
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	auth := []string{"abcde123456", "tests"}
	client := &servicemock.Client{}
	graylog := &graylogmock.GraylogBase{}

	graylog.On("GetAuth", "dashboards").Return(auth, nil).Once()
	graylog.On("HandleFailure", []byte(dashboardResponse), 200).Return(nil).Once()

	client.On("Get", endpoint, headers, auth, nil).Return([]byte(dashboardResponse), 200, nil).Once()

	dashboardClient := dashboard.New(client, graylog)
	assert.NotNil(t, dashboardClient)

	dashboardClient.List()

	client.AssertExpectations(t)
	graylog.AssertExpectations(t)
}

func TestCreate(t *testing.T) {
	t.Parallel()
	endpoint := "/api/dashboards"
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	auth := []string{"abcde123456", "tests"}
	client := &servicemock.Client{}
	graylog := &graylogmock.GraylogBase{}
	dashboardToCreate := dto.Dashboard{
		Title:       "bar",
		Description: "description for bar dashboard",
	}
	body, err := json.Marshal(dashboardToCreate)
	assert.NotNil(t, body)
	assert.Nil(t, err)

	graylog.On("GetAuth", "dashboards").Return(auth, nil).Once()
	graylog.On("HandleFailure", []byte(dashboardResponse), 200).Return(nil).Once()

	client.On("Post", endpoint, headers, auth, bytes.NewBuffer(body)).Return([]byte(dashboardResponse), 200, nil).Once()

	dashboardClient := dashboard.New(client, graylog)
	assert.NotNil(t, dashboardClient)

	dashboardClient.Create(dashboardToCreate)

	client.AssertExpectations(t)
	graylog.AssertExpectations(t)
}
