package stream

// Endpoint define the allowed endpoint
type Endpoint string

const (
	ListEndpoint   Endpoint = "/api/streams"
	GetEndpoint    Endpoint = "/api/streams/%s"
	CreateEndpoint Endpoint = "/api/streams"
)

// String return the string representation of the Endpoint
func (endpoint Endpoint) String() string {
	return string(endpoint)
}
