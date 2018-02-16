package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/uniplaces/logfairy/dto/stream"
	"github.com/uniplaces/logfairy/infrastructure/api/dto"
)

// endpoint constants
const (
	streams = "/api/streams"
	// /api/streams/<stream_id>
	singleStream    = "/api/streams/%s"
	streamTokenName = "streams"
)

type StreamClient struct {
	Graylog
}

// New create an instance of Graylog stream api client
func NewStreamClient(graylog Graylog) StreamClient {
	return StreamClient{Graylog: graylog}
}

func (streamClient *StreamClient) List() (stream.Streams, error) {
	auth, err := streamClient.getAuth(streamTokenName)
	if err != nil {
		return stream.Streams{}, err
	}

	response, status, err := streamClient.client.Get(streams, headers, auth, nil)
	if err != nil {
		return stream.Streams{}, err
	}

	failure, err := streamClient.handleFailure(response, status)
	if err != nil {
		return stream.Streams{}, err
	}
	if failure != nil {
		return stream.Streams{}, errors.New(failure.Message)
	}

	list := stream.Streams{}
	if err := json.Unmarshal(response, &list); err != nil {
		return stream.Streams{}, err
	}

	return list, nil
}

func (streamClient *StreamClient) Get(streamID string) (stream.Stream, error) {
	auth, err := streamClient.getAuth(streamTokenName)
	if err != nil {
		return stream.Stream{}, err
	}

	endpoint := fmt.Sprintf(singleStream, streamID)
	response, status, err := streamClient.client.Get(endpoint, headers, auth, nil)
	if err != nil {
		return stream.Stream{}, err
	}

	failure, err := streamClient.handleFailure(response, status)
	if err != nil {
		return stream.Stream{}, err
	}
	if failure != nil {
		return stream.Stream{}, errors.New(failure.Message)
	}

	streamFound := stream.Stream{}
	if err := json.Unmarshal(response, &streamFound); err != nil {
		return stream.Stream{}, err
	}

	return streamFound, nil
}

func (streamClient *StreamClient) Create(streamToCreate stream.Stream) (string, error) {
	auth, err := streamClient.getAuth(streamTokenName)
	if err != nil {
		return "", err
	}

	body, err := json.Marshal(streamToCreate)
	if err != nil {
		return "", err
	}

	response, status, err := streamClient.client.Post(streams, headers, auth, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	failure, err := streamClient.handleFailure(response, status)
	if err != nil {
		return "", err
	}
	if failure != nil {
		return "", errors.New(failure.Message)
	}

	success := dto.StreamCreation{}
	if err := json.Unmarshal(response, &success); err != nil {
		return "", err
	}

	return success.StreamID, nil
}
