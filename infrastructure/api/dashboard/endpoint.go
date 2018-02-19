package dashboard

// Endpoint define the allowed endpoint
type Endpoint string

const (
	ListEndpoint   Endpoint = "/api/dashboards"
	GetEndpoint    Endpoint = "/api/dashboards/%s"
	CreateEndpoint Endpoint = "/api/dashboards"
)

// String return the string representation of the Endpoint
func (endpoint Endpoint) String() string {
	return string(endpoint)
}
