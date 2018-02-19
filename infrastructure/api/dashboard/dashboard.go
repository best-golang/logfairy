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

// Client is a graylog client specialized in dashboard actions
type Client struct {
	api.Graylog
}

// New create an instance of Graylog deashboard api client
func New(graylog api.Graylog) Client {
	return Client{Graylog: graylog}
}

// List return all the dashboards it can reach
func (client *Client) List() (dashboard.Dashboards, error) {
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
func (client *Client) Get(dashboardID string) (dashboard.Dashboard, error) {
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
func (client *Client) Create(dashboardToCreate dashboard.Dashboard) (string, error) {
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
