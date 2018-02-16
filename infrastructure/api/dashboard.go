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
	dashboards = "/api/dashboards"
	// /api/dashboards/<deashboard_id>
	singleDashboard    = "/api/dashboards/%s"
	dashboardTokenName = "dashboards"
)

type DashboardClient struct {
	Graylog
}

// New create an instance of Graylog deashboard api client
func NewDashboardClient(graylog Graylog) DashboardClient {
	return DashboardClient{Graylog: graylog}
}

func (dashboardClient *DashboardClient) List() (dashboard.Dashboards, error) {
	auth, err := dashboardClient.getAuth(dashboardTokenName)
	if err != nil {
		return dashboard.Dashboards{}, err
	}

	response, status, err := dashboardClient.client.Get(dashboards, headers, auth, nil)
	if err != nil {
		return dashboard.Dashboards{}, err
	}

	failure, err := dashboardClient.handleFailure(response, status)
	if err != nil {
		return dashboard.Dashboards{}, err
	}
	if failure != nil {
		return dashboard.Dashboards{}, errors.New(failure.Message)
	}

	list := dashboard.Dashboards{}
	if err := json.Unmarshal(response, &list); err != nil {
		return dashboard.Dashboards{}, err
	}

	return list, nil
}

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

	failure, err := dashboardClient.handleFailure(response, status)
	if err != nil {
		return dashboard.Dashboard{}, err
	}
	if failure != nil {
		return dashboard.Dashboard{}, errors.New(failure.Message)
	}

	dashboardFound := dashboard.Dashboard{}
	if err := json.Unmarshal(response, &dashboardFound); err != nil {
		return dashboard.Dashboard{}, err
	}

	return dashboardFound, nil
}

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

	failure, err := dashboardClient.handleFailure(response, status)
	if err != nil {
		return "", err
	}
	if failure != nil {
		return "", errors.New(failure.Message)
	}

	success := dto.DashboardCreation{}
	if err := json.Unmarshal(response, &success); err != nil {
		return "", err
	}

	return success.DashboardID, nil
}
