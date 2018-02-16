package api

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/uniplaces/logfairy/dto/stream"
	"github.com/uniplaces/logfairy/infrastructure/api/dto"
)

// endpoint used from stream client
const (
	streams = "/api/streams"
	// /api/streams/<stream_id>
	singleStream    = "/api/streams/%s"
	streamTokenName = "streams"
)

// StreamClient is a graylog client specialized in stream actions
type StreamClient struct {
	Graylog
}

// New create an instance of Graylog stream api client
func NewStreamClient(graylog Graylog) StreamClient {
	return StreamClient{Graylog: graylog}
}

// List return all the streams it can reach
func (streamClient *StreamClient) List() (stream.Streams, error) {
	auth, err := streamClient.getAuth(streamTokenName)
	if err != nil {
		return stream.Streams{}, err
	}

	response, status, err := streamClient.client.Get(streams, headers, auth, nil)
	if err != nil {
		return stream.Streams{}, err
	}

	if err := streamClient.handleFailure(response, status); err != nil {
		return stream.Streams{}, err
	}

	list := stream.Streams{}
	if err := json.Unmarshal(response, &list); err != nil {
		return stream.Streams{}, err
	}

	return list, nil
}

// Get try to return a stream given the stream id
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

	if err := streamClient.handleFailure(response, status); err != nil {
		return stream.Stream{}, err
	}

	streamFound := stream.Stream{}
	if err := json.Unmarshal(response, &streamFound); err != nil {
		return stream.Stream{}, err
	}

	return streamFound, nil
}

// Create create a stream
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

	if err := streamClient.handleFailure(response, status); err != nil {
		return "", err
	}

	success := dto.StreamCreation{}
	if err := json.Unmarshal(response, &success); err != nil {
		return "", err
	}

	return success.StreamID, nil
}
