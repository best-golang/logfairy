package stream_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	dto "github.com/uniplaces/logfairy/dto/stream"
	graylogmock "github.com/uniplaces/logfairy/infrastructure/api/mocks"
	servicemock "github.com/uniplaces/logfairy/infrastructure/api/service/mocks"
	"github.com/uniplaces/logfairy/infrastructure/api/stream"
)

const (
	streamResponse = `{
		"title": "foo",
		"description": "description for foo",
		"matching_type": "AND",
		"rules": [
			{
				"field": "_env",
				"description": "",
				"type": 1,
				"inverted": false,
				"value": "prod"
			},
			{
				"field": "_app-id",
				"description": "",
				"type": 1,
				"inverted": false,
				"value": "foo"
			}
		],
		"remove_matches_from_default_stream": false,
		"index_set_id": "5b0bfb3bgg857f3b700b58g5"
	}`
)

func TestCreation(t *testing.T) {
	t.Parallel()

	client := &servicemock.Client{}
	graylog := &graylogmock.GraylogBase{}
	streamClient := stream.New(client, graylog)
	assert.NotNil(t, streamClient)

	client.AssertExpectations(t)
	graylog.AssertExpectations(t)
}

func TestList(t *testing.T) {
	t.Parallel()
	endpoint := "/api/streams"
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	auth := []string{"abcde123456", "tests"}
	client := &servicemock.Client{}
	graylog := &graylogmock.GraylogBase{}

	graylog.On("GetAuth", "streams").Return(auth, nil).Once()
	graylog.On("HandleFailure", []byte(streamResponse), 200).Return(nil).Once()

	client.On("Get", endpoint, headers, auth, nil).Return([]byte(streamResponse), 200, nil).Once()

	streamClient := stream.New(client, graylog)
	assert.NotNil(t, streamClient)

	streamClient.List()

	client.AssertExpectations(t)
	graylog.AssertExpectations(t)
}

func TestGet(t *testing.T) {
	t.Parallel()
	endpoint := "/api/streams/streamID"
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	auth := []string{"abcde123456", "tests"}
	client := &servicemock.Client{}
	graylog := &graylogmock.GraylogBase{}

	graylog.On("GetAuth", "streams").Return(auth, nil).Once()
	graylog.On("HandleFailure", []byte(streamResponse), 200).Return(nil).Once()

	client.On("Get", endpoint, headers, auth, nil).Return([]byte(streamResponse), 200, nil).Once()

	streamClient := stream.New(client, graylog)
	assert.NotNil(t, streamClient)

	streamClient.Get("streamID")

	client.AssertExpectations(t)
	graylog.AssertExpectations(t)
}

func TestCreate(t *testing.T) {
	t.Parallel()
	endpoint := "/api/streams"
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	auth := []string{"abcde123456", "tests"}
	client := &servicemock.Client{}
	graylog := &graylogmock.GraylogBase{}
	streamToCreate := dto.Stream{
		Title:       "bar",
		Description: "description for bar stream",
	}
	body, err := json.Marshal(streamToCreate)
	assert.NotNil(t, body)
	assert.Nil(t, err)

	graylog.On("GetAuth", "streams").Return(auth, nil).Once()
	graylog.On("HandleFailure", []byte(streamResponse), 200).Return(nil).Once()

	client.On("Post", endpoint, headers, auth, bytes.NewBuffer(body)).Return([]byte(streamResponse), 200, nil).Once()

	streamClient := stream.New(client, graylog)
	assert.NotNil(t, streamClient)

	streamClient.Create(streamToCreate)

	client.AssertExpectations(t)
	graylog.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	t.Parallel()
	endpoint := "/api/streams/StreamID"
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	auth := []string{"abcde123456", "tests"}
	client := &servicemock.Client{}
	graylog := &graylogmock.GraylogBase{}
	streamToUpdate := dto.Stream{
		Title:       "bar",
		Description: "description for bar stream",
	}
	body, err := json.Marshal(streamToUpdate)
	assert.NotNil(t, body)
	assert.Nil(t, err)

	graylog.On("GetAuth", "streams").Return(auth, nil).Once()
	graylog.On("HandleFailure", []byte(streamResponse), 200).Return(nil).Once()

	client.On("Put", endpoint, headers, auth, bytes.NewBuffer(body)).Return([]byte(streamResponse), 200, nil).Once()

	streamClient := stream.New(client, graylog)
	assert.NotNil(t, streamClient)

	streamClient.Update("StreamID", streamToUpdate)

	client.AssertExpectations(t)
	graylog.AssertExpectations(t)
}

func TestResume(t *testing.T) {
	t.Parallel()
	endpoint := "/api/streams/StreamID/resume"
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	auth := []string{"abcde123456", "tests"}
	client := &servicemock.Client{}
	graylog := &graylogmock.GraylogBase{}

	graylog.On("GetAuth", "streams").Return(auth, nil).Once()
	graylog.On("HandleFailure", []byte(streamResponse), 200).Return(nil).Once()

	client.On("Post", endpoint, headers, auth, nil).Return([]byte(streamResponse), 200, nil).Once()

	streamClient := stream.New(client, graylog)
	assert.NotNil(t, streamClient)

	streamClient.Resume("StreamID")

	client.AssertExpectations(t)
	graylog.AssertExpectations(t)
}

func TestPause(t *testing.T) {
	t.Parallel()
	endpoint := "/api/streams/StreamID/pause"
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	auth := []string{"abcde123456", "tests"}
	client := &servicemock.Client{}
	graylog := &graylogmock.GraylogBase{}

	graylog.On("GetAuth", "streams").Return(auth, nil).Once()
	graylog.On("HandleFailure", []byte(streamResponse), 200).Return(nil).Once()

	client.On("Post", endpoint, headers, auth, nil).Return([]byte(streamResponse), 200, nil).Once()

	streamClient := stream.New(client, graylog)
	assert.NotNil(t, streamClient)

	streamClient.Pause("StreamID")

	client.AssertExpectations(t)
	graylog.AssertExpectations(t)
}
