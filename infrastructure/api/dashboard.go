package api

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/uniplaces/logfairy/dto/dashboard"
	"github.com/uniplaces/logfairy/infrastructure/api/dto"
)

// endpoint used from dashboard client
const (
	dashboards = "/api/dashboards"
	// /api/dashboards/<deashboard_id>
	singleDashboard    = "/api/dashboards/%s"
	dashboardTokenName = "dashboards"
)

// DashboardClient is a graylog client specialized in dashboard actions
type DashboardClient struct {
	Graylog
}

// New create an instance of Graylog deashboard api client
func NewDashboardClient(graylog Graylog) DashboardClient {
	return DashboardClient{Graylog: graylog}
}

// List return all the dashboards it can reach
func (dashboardClient *DashboardClient) List() (dashboard.Dashboards, error) {
	auth, err := dashboardClient.getAuth(dashboardTokenName)
	if err != nil {
		return dashboard.Dashboards{}, err
	}

	response, status, err := dashboardClient.client.Get(dashboards, headers, auth, nil)
	if err != nil {
		return dashboard.Dashboards{}, err
	}

	if err := dashboardClient.handleFailure(response, status); err != nil {
		return dashboard.Dashboards{}, err
	}

	list := dashboard.Dashboards{}
	if err := json.Unmarshal(response, &list); err != nil {
		return dashboard.Dashboards{}, err
	}

	return list, nil
}

// Get try to return a dashboard given the dashboard id
func (dashboardClient *DashboardClient) Get(dashboardID string) (dashboard.Dashboard, error) {
	auth, err := dashboardClient.getAuth(dashboardTokenName)
	if err != nil {
		return dashboard.Dashboard{}, err
	}

	endpoint := fmt.Sprintf(singleDashboard, dashboardID)
	response, status, err := dashboardClient.client.Get(endpoint, headers, auth, nil)
	if err != nil {
		return dashboard.Dashboard{}, err
	}

	if err := dashboardClient.handleFailure(response, status); err != nil {
		return dashboard.Dashboard{}, err
	}

	dashboardFound := dashboard.Dashboard{}
	if err := json.Unmarshal(response, &dashboardFound); err != nil {
		return dashboard.Dashboard{}, err
	}

	return dashboardFound, nil
}

// Create create a dashboard
func (dashboardClient *DashboardClient) Create(dashboardToCreate dashboard.Dashboard) (string, error) {
	auth, err := dashboardClient.getAuth(dashboardTokenName)
	if err != nil {
		return "", err
	}

	body, err := json.Marshal(dashboardToCreate)
	if err != nil {
		return "", err
	}

	response, status, err := dashboardClient.client.Post(dashboards, headers, auth, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	if err := dashboardClient.handleFailure(response, status); err != nil {
		return "", err
	}

	success := dto.DashboardCreation{}
	if err := json.Unmarshal(response, &success); err != nil {
		return "", err
	}

	return success.DashboardID, nil
}
