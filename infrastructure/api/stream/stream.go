package stream

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/uniplaces/logfairy/dto/stream"
	"github.com/uniplaces/logfairy/infrastructure/api"
	"github.com/uniplaces/logfairy/infrastructure/api/dto"
)

const streamTokenName = "streams"

var headers = map[string]string{
	"Content-Type": "application/json",
}

// Stream is a graylog client specialized in stream actions
type Stream struct {
	api.Client
	api.GraylogBase
}

// New create an instance of Graylog stream api client
func New(client api.Client, graylog api.GraylogBase) Stream {
	return Stream{Client: client, GraylogBase: graylog}
}

// List return all the streams it can reach
func (streamClient *Stream) List() (stream.Streams, error) {
	auth, err := streamClient.GetAuth(streamTokenName)
	if err != nil {
		return stream.Streams{}, err
	}

	response, status, err := streamClient.Client.Get(ListEndpoint.String(), headers, auth, nil)
	if err != nil {
		return stream.Streams{}, err
	}

	if err := streamClient.HandleFailure(response, status); err != nil {
		return stream.Streams{}, err
	}

	list := stream.Streams{}
	if err := json.Unmarshal(response, &list); err != nil {
		return stream.Streams{}, err
	}

	return list, nil
}

// Get try to return a stream given the stream id
func (streamClient *Stream) Get(streamID string) (stream.Stream, error) {
	auth, err := streamClient.GetAuth(streamTokenName)
	if err != nil {
		return stream.Stream{}, err
	}

	endpoint := fmt.Sprintf(GetEndpoint.String(), streamID)
	response, status, err := streamClient.Client.Get(endpoint, headers, auth, nil)
	if err != nil {
		return stream.Stream{}, err
	}

	if err := streamClient.HandleFailure(response, status); err != nil {
		return stream.Stream{}, err
	}

	streamFound := stream.Stream{}
	if err := json.Unmarshal(response, &streamFound); err != nil {
		return stream.Stream{}, err
	}

	return streamFound, nil
}

// Create creates a stream
func (streamClient *Stream) Create(streamToCreate stream.Stream) (string, error) {
	auth, err := streamClient.GetAuth(streamTokenName)
	if err != nil {
		return "", err
	}

	body, err := json.Marshal(streamToCreate)
	if err != nil {
		return "", err
	}

	response, status, err := streamClient.Client.Post(CreateEndpoint.String(), headers, auth, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	if err := streamClient.HandleFailure(response, status); err != nil {
		return "", err
	}

	success := dto.StreamCreation{}
	if err := json.Unmarshal(response, &success); err != nil {
		return "", err
	}

	return success.StreamID, nil
}

// Update updates a stream
func (streamClient *Stream) Update(streamID string, streamToUpdate stream.Stream) (string, error) {
	auth, err := streamClient.GetAuth(streamTokenName)
	if err != nil {
		return "", err
	}

	body, err := json.Marshal(streamToUpdate)
	if err != nil {
		return "", err
	}

	endpoint := fmt.Sprintf(GetEndpoint.String(), streamID)
	response, status, err := streamClient.Client.Put(endpoint, headers, auth, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	if err := streamClient.HandleFailure(response, status); err != nil {
		return "", err
	}

	success := dto.StreamCreation{}
	if err := json.Unmarshal(response, &success); err != nil {
		return "", err
	}

	return success.StreamID, nil
}

// Resume resumes a stream
func (streamClient *Stream) Resume(streamID string) (stream.Stream, error) {
	auth, err := streamClient.GetAuth(streamTokenName)
	if err != nil {
		return stream.Stream{}, err
	}

	endpoint := fmt.Sprintf(ResumeEndpoint.String(), streamID)
	response, status, err := streamClient.Client.Post(endpoint, headers, auth, nil)
	if err != nil {
		return stream.Stream{}, err
	}

	if err := streamClient.HandleFailure(response, status); err != nil {
		return stream.Stream{}, err
	}

	streamFound := stream.Stream{}
	if err := json.Unmarshal(response, &streamFound); err != nil {
		return stream.Stream{}, err
	}

	return streamFound, nil
}

// Pause pauses a stream
func (streamClient *Stream) Pause(streamID string) (stream.Stream, error) {
	auth, err := streamClient.GetAuth(streamTokenName)
	if err != nil {
		return stream.Stream{}, err
	}

	endpoint := fmt.Sprintf(PauseEndpoint.String(), streamID)
	response, status, err := streamClient.Client.Post(endpoint, headers, auth, nil)
	if err != nil {
		return stream.Stream{}, err
	}

	if err := streamClient.HandleFailure(response, status); err != nil {
		return stream.Stream{}, err
	}

	streamFound := stream.Stream{}
	if err := json.Unmarshal(response, &streamFound); err != nil {
		return stream.Stream{}, err
	}

	return streamFound, nil
}
