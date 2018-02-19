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

// Client is a graylog client specialized in stream actions
type Client struct {
	api.Graylog
}

// New create an instance of Graylog stream api client
func New(graylog api.Graylog) Client {
	return Client{Graylog: graylog}
}

// List return all the streams it can reach
func (client *Client) List() (stream.Streams, error) {
	auth, err := client.GetAuth(streamTokenName)
	if err != nil {
		return stream.Streams{}, err
	}

	response, status, err := client.Client.Get(ListEndpoint.String(), headers, auth, nil)
	if err != nil {
		return stream.Streams{}, err
	}

	if err := client.HandleFailure(response, status); err != nil {
		return stream.Streams{}, err
	}

	list := stream.Streams{}
	if err := json.Unmarshal(response, &list); err != nil {
		return stream.Streams{}, err
	}

	return list, nil
}

// Get try to return a stream given the stream id
func (client *Client) Get(streamID string) (stream.Stream, error) {
	auth, err := client.GetAuth(streamTokenName)
	if err != nil {
		return stream.Stream{}, err
	}

	endpoint := fmt.Sprintf(GetEndpoint.String(), streamID)
	response, status, err := client.Client.Get(endpoint, headers, auth, nil)
	if err != nil {
		return stream.Stream{}, err
	}

	if err := client.HandleFailure(response, status); err != nil {
		return stream.Stream{}, err
	}

	streamFound := stream.Stream{}
	if err := json.Unmarshal(response, &streamFound); err != nil {
		return stream.Stream{}, err
	}

	return streamFound, nil
}

// Create create a stream
func (client *Client) Create(streamToCreate stream.Stream) (string, error) {
	auth, err := client.GetAuth(streamTokenName)
	if err != nil {
		return "", err
	}

	body, err := json.Marshal(streamToCreate)
	if err != nil {
		return "", err
	}

	response, status, err := client.Client.Post(CreateEndpoint.String(), headers, auth, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	if err := client.HandleFailure(response, status); err != nil {
		return "", err
	}

	success := dto.StreamCreation{}
	if err := json.Unmarshal(response, &success); err != nil {
		return "", err
	}

	return success.StreamID, nil
}
