package widget

// Endpoint define the allowed endpoint
type Endpoint string

const (
	GetEndpoint    Endpoint = "/api/dashboards/%s/widgets/%s"
	CreateEndpoint Endpoint = "/api/dashboards/%s/widgets"
)

// String return the string representation of the Endpoint
func (endpoint Endpoint) String() string {
	return string(endpoint)
}
