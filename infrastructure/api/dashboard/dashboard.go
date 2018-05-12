package dashboard

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/uniplaces/logfairy/dto/dashboard"
	"github.com/uniplaces/logfairy/infrastructure/api"
	"github.com/uniplaces/logfairy/infrastructure/api/dto"
)

const dashboardTokenName = "dashboards"

var headers = map[string]string{
	"Content-Type": "application/json",
}

// Dashboard is a graylog client specialized in dashboard actions
type Dashboard struct {
	api.Client
	api.GraylogBase
}

// New create an instance of Graylog deashboard api client
func New(client api.Client, graylog api.GraylogBase) Dashboard {
	return Dashboard{Client: client, GraylogBase: graylog}
}

// List return all the dashboards it can reach
func (client *Dashboard) List() (dashboard.Dashboards, error) {
	auth, err := client.GetAuth(dashboardTokenName)
	if err != nil {
		return dashboard.Dashboards{}, err
	}

	response, status, err := client.Client.Get(ListEndpoint.String(), headers, auth, nil)
	if err != nil {
		return dashboard.Dashboards{}, err
	}

	if err := client.HandleFailure(response, status); err != nil {
		return dashboard.Dashboards{}, err
	}

	list := dashboard.Dashboards{}
	if err := json.Unmarshal(response, &list); err != nil {
		return dashboard.Dashboards{}, err
	}

	return list, nil
}

// Get try to return a dashboard given the dashboard id
func (client *Dashboard) Get(dashboardID string) (dashboard.Dashboard, error) {
	auth, err := client.GetAuth(dashboardTokenName)
	if err != nil {
		return dashboard.Dashboard{}, err
	}

	endpoint := fmt.Sprintf(GetEndpoint.String(), dashboardID)
	response, status, err := client.Client.Get(endpoint, headers, auth, nil)
	if err != nil {
		return dashboard.Dashboard{}, err
	}

	if err := client.HandleFailure(response, status); err != nil {
		return dashboard.Dashboard{}, err
	}

	dashboardFound := dashboard.Dashboard{}
	if err := json.Unmarshal(response, &dashboardFound); err != nil {
		return dashboard.Dashboard{}, err
	}

	return dashboardFound, nil
}

// Create create a dashboard
func (client *Dashboard) Create(dashboardToCreate dashboard.Dashboard) (string, error) {
	auth, err := client.GetAuth(dashboardTokenName)
	if err != nil {
		return "", err
	}

	body, err := json.Marshal(dashboardToCreate)
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

	success := dto.DashboardCreation{}
	if err := json.Unmarshal(response, &success); err != nil {
		return "", err
	}

	return success.DashboardID, nil
}
