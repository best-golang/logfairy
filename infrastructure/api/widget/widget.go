package widget

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/uniplaces/logfairy/dto/dashboard"
	"github.com/uniplaces/logfairy/infrastructure/api"
	"github.com/uniplaces/logfairy/infrastructure/api/dto"
)

const widgetTokenName = "widgets"

var headers = map[string]string{
	"Content-Type": "application/json",
}

// Client is a graylog client specialized in widget actions
type Client struct {
	api.Graylog
}

// New create an instance of Graylog deashboard api client
func New(graylog api.Graylog) Client {
	return Client{Graylog: graylog}
}

// Get try to return a widget given the widget id and the dashboard id
func (client *Client) Get(widgetID string, dashnoardID string) (dashboard.Widget, error) {
	auth, err := client.GetAuth(widgetTokenName)
	if err != nil {
		return dashboard.Widget{}, err
	}

	endpoint := fmt.Sprintf(GetEndpoint.String(), dashnoardID, widgetID)
	response, status, err := client.Client.Get(endpoint, headers, auth, nil)
	if err != nil {
		return dashboard.Widget{}, err
	}

	if err := client.HandleFailure(response, status); err != nil {
		return dashboard.Widget{}, err
	}

	dashboardFound := dashboard.Widget{}
	if err := json.Unmarshal(response, &dashboardFound); err != nil {
		return dashboard.Widget{}, err
	}

	return dashboardFound, nil
}

// Create create a widget for the given dashboard
func (client *Client) Create(widgetToCreate dashboard.Widget, dashnoardID string) (string, error) {
	auth, err := client.GetAuth(widgetTokenName)
	if err != nil {
		return "", err
	}

	body, err := json.Marshal(widgetToCreate)
	if err != nil {
		return "", err
	}

	endpoint := fmt.Sprintf(CreateEndpoint.String(), dashnoardID)
	response, status, err := client.Client.Post(endpoint, headers, auth, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	if err := client.HandleFailure(response, status); err != nil {
		return "", err
	}

	success := dto.WidgetCreation{}
	if err := json.Unmarshal(response, &success); err != nil {
		return "", err
	}

	return success.WidgetID, nil
}
