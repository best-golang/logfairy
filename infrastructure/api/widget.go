package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/uniplaces/logfairy/dto/dashboard"
	"github.com/uniplaces/logfairy/infrastructure/api/dto"
)

// endpoint constants
const (
	// /api/dashboards/<dashboard_id>/widgets/<widget_id>
	singleWidget = "/api/dashboards/%s/widgets/%s"
	// /api/dashboards/<dashboard_id>/widgets
	createWidget    = "/api/dashboards/%s/widgets"
	widgetTokenName = "widgets"
)

type WidgetClient struct {
	Graylog
}

// New create an instance of Graylog deashboard api client
func NewWidgetClient(graylog Graylog) WidgetClient {
	return WidgetClient{Graylog: graylog}
}

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

	failure, err := widgetClient.handleFailure(response, status)
	if err != nil {
		return dashboard.Widget{}, err
	}
	if failure != nil {
		return dashboard.Widget{}, errors.New(failure.Message)
	}

	dashboardFound := dashboard.Widget{}
	if err := json.Unmarshal(response, &dashboardFound); err != nil {
		return dashboard.Widget{}, err
	}

	return dashboardFound, nil
}

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

	failure, err := widgetClient.handleFailure(response, status)
	if err != nil {
		return "", err
	}
	if failure != nil {
		return "", errors.New(failure.Message)
	}

	success := dto.WidgetCreation{}
	if err := json.Unmarshal(response, &success); err != nil {
		return "", err
	}

	return success.WidgetID, nil
}
