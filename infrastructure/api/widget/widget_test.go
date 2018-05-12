package widget_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	dto "github.com/uniplaces/logfairy/dto/dashboard"
	graylogmock "github.com/uniplaces/logfairy/infrastructure/api/mocks"
	servicemock "github.com/uniplaces/logfairy/infrastructure/api/service/mocks"
	"github.com/uniplaces/logfairy/infrastructure/api/widget"
)

const (
	widgetResponse = `{
		"cache_time": 1800,
		"description": "foo foo bar bar",
		"type": "STREAM_SEARCH_RESULT_COUNT",
		"config": {
			"timerange": {
				"type": "relative",
				"range": 86400
			},
			"lower_is_better": false,
			"stream_id": d87d7afd-89c5-4233-a20b-4476785f11cb,
			"trend": true,
			"query": "foo bar"
		}
	}`
)

func TestCreation(t *testing.T) {
	t.Parallel()

	client := &servicemock.Client{}
	graylog := &graylogmock.GraylogBase{}
	widgetClient := widget.New(client, graylog)
	assert.NotNil(t, widgetClient)

	client.AssertExpectations(t)
	graylog.AssertExpectations(t)
}

func TestGet(t *testing.T) {
	t.Parallel()
	endpoint := "/api/dashboards/dashnoardID/widgets/widgetID"
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	auth := []string{"abcde123456", "tests"}
	client := &servicemock.Client{}
	graylog := &graylogmock.GraylogBase{}

	graylog.On("GetAuth", "widgets").Return(auth, nil).Once()
	graylog.On("HandleFailure", []byte(widgetResponse), 200).Return(nil).Once()

	client.On("Get", endpoint, headers, auth, nil).Return([]byte(widgetResponse), 200, nil).Once()

	widgetClient := widget.New(client, graylog)
	assert.NotNil(t, widgetClient)

	widgetClient.Get("widgetID", "dashnoardID")

	client.AssertExpectations(t)
	graylog.AssertExpectations(t)
}

func TestCreate(t *testing.T) {
	t.Parallel()
	endpoint := "/api/dashboards/dashnoardID/widgets"
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	auth := []string{"abcde123456", "tests"}
	client := &servicemock.Client{}
	graylog := &graylogmock.GraylogBase{}
	widgetToCreate := dto.Widget{
		Description: "foo foo bar bar",
		Type:        "STREAM_SEARCH_RESULT_COUNT",
	}
	body, err := json.Marshal(widgetToCreate)
	assert.NotNil(t, body)
	assert.Nil(t, err)

	graylog.On("GetAuth", "widgets").Return(auth, nil).Once()
	graylog.On("HandleFailure", []byte(widgetResponse), 200).Return(nil).Once()

	client.On("Post", endpoint, headers, auth, bytes.NewBuffer(body)).Return([]byte(widgetResponse), 200, nil).Once()

	widgetClient := widget.New(client, graylog)
	assert.NotNil(t, widgetClient)

	widgetClient.Create(widgetToCreate, "dashnoardID")

	client.AssertExpectations(t)
	graylog.AssertExpectations(t)
}
