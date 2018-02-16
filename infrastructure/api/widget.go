package api

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/uniplaces/logfairy/dto/dashboard"
	"github.com/uniplaces/logfairy/infrastructure/api/dto"
)

// endpoint used from widget client
const (
	// /api/dashboards/<dashboard_id>/widgets/<widget_id>
	singleWidget = "/api/dashboards/%s/widgets/%s"
	// /api/dashboards/<dashboard_id>/widgets
	createWidget    = "/api/dashboards/%s/widgets"
	widgetTokenName = "widgets"
)

// WidgetClient is a graylog client specialized in widget actions
type WidgetClient struct {
	Graylog
}

// New create an instance of Graylog deashboard api client
func NewWidgetClient(graylog Graylog) WidgetClient {
	return WidgetClient{Graylog: graylog}
}

// Get try to return a widget given the widget id and the dashboard id
func (widgetClient *WidgetClient) Get(widgetID string, dashnoardID string) (dashboard.Widget, error) {
	auth, err := widgetClient.getAuth(widgetTokenName)
	if err != nil {
		return dashboard.Widget{}, err
	}

	endpoint := fmt.Sprintf(singleWidget, dashnoardID, widgetID)
	response, status, err := widgetClient.client.Get(endpoint, headers, auth, nil)
	if err != nil {
		return dashboard.Widget{}, err
	}

	if err := widgetClient.handleFailure(response, status); err != nil {
		return dashboard.Widget{}, err
	}

	dashboardFound := dashboard.Widget{}
	if err := json.Unmarshal(response, &dashboardFound); err != nil {
		return dashboard.Widget{}, err
	}

	return dashboardFound, nil
}

// Create create a widget for the given dashboard
func (widgetClient *WidgetClient) Create(widgetToCreate dashboard.Widget, dashnoardID string) (string, error) {
	auth, err := widgetClient.getAuth(widgetTokenName)
	if err != nil {
		return "", err
	}

	body, err := json.Marshal(widgetToCreate)
	if err != nil {
		return "", err
	}

	endpoint := fmt.Sprintf(createWidget, dashnoardID)
	response, status, err := widgetClient.client.Post(endpoint, headers, auth, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	if err := widgetClient.handleFailure(response, status); err != nil {
		return "", err
	}

	success := dto.WidgetCreation{}
	if err := json.Unmarshal(response, &success); err != nil {
		return "", err
	}

	return success.WidgetID, nil
}
