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

// Widget is a graylog client specialized in widget actions
type Widget struct {
	api.Client
	api.GraylogBase
}

// New create an instance of Graylog deashboard api client
func New(client api.Client, graylog api.GraylogBase) Widget {
	return Widget{Client: client, GraylogBase: graylog}
}

// Get try to return a widget given the widget id and the dashboard id
func (widget *Widget) Get(widgetID string, dashnoardID string) (dashboard.Widget, error) {
	auth, err := widget.GetAuth(widgetTokenName)
	if err != nil {
		return dashboard.Widget{}, err
	}

	endpoint := fmt.Sprintf(GetEndpoint.String(), dashnoardID, widgetID)
	response, status, err := widget.Client.Get(endpoint, headers, auth, nil)
	if err != nil {
		return dashboard.Widget{}, err
	}

	if err := widget.HandleFailure(response, status); err != nil {
		return dashboard.Widget{}, err
	}

	widgetFound := dashboard.Widget{}
	if err := json.Unmarshal(response, &widgetFound); err != nil {
		return dashboard.Widget{}, err
	}

	return widgetFound, nil
}

// Create create a widget for the given dashboard
func (widget *Widget) Create(widgetToCreate dashboard.Widget, dashnoardID string) (string, error) {
	auth, err := widget.GetAuth(widgetTokenName)
	if err != nil {
		return "", err
	}

	body, err := json.Marshal(widgetToCreate)
	if err != nil {
		return "", err
	}

	endpoint := fmt.Sprintf(CreateEndpoint.String(), dashnoardID)
	response, status, err := widget.Client.Post(endpoint, headers, auth, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	if err := widget.HandleFailure(response, status); err != nil {
		return "", err
	}

	success := dto.WidgetCreation{}
	if err := json.Unmarshal(response, &success); err != nil {
		return "", err
	}

	return success.WidgetID, nil
}
