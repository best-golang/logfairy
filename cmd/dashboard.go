package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
	"github.com/uniplaces/logfairy/dto/dashboard"
	"github.com/uniplaces/logfairy/infrastructure/api"
)

var dashboardID string

// dashboardsCmd represents the dashboards command
var dashboardsCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "handle dashboards actions",
	Long:  ``,
}

// ListDashboardsCmd represents the List command
var ListDashboardsCmd = &cobra.Command{
	Use:   "list",
	Short: "list dashboards",
	Long:  `dashboard list look up for all the dashboards it is allowed to reach.`,
	Run: func(cmd *cobra.Command, args []string) {
		dashboards, err := dashboardClient.List()
		if err != nil {
			log.Fatalln(err)
		}

		prettyPrint(dashboards)
	},
}

// GetDashboardCmd represents the list single dashboard command
var GetDashboardCmd = &cobra.Command{
	Use:   "get",
	Short: "list a single dashboard",
	Long:  `dashboard get look up for a dashboard with the given dashboard id and print out the json.`,
	Run: func(cmd *cobra.Command, args []string) {
		dashboard, err := dashboardClient.Get(dashboardID)
		if err != nil {
			log.Fatalln(err)
		}

		prettyPrint(dashboard)
	},
}

// CreateDashboardCmd represents the create command
var CreateDashboardCmd = &cobra.Command{
	Use:   "create",
	Short: "crate a dashboard",
	Long: `dashboard create will try to create the dashboard defined in the config file.
	
The expected json is:
{
  "title": "bar",
  "description": "description for bar dashboard"
}`,
	Run: func(cmd *cobra.Command, args []string) {
		definition, err := ioutil.ReadFile(definitions)
		if err != nil {
			log.Fatalln(err)
		}

		dashboardToCreate := dashboard.Dashboard{}
		if err := json.Unmarshal(definition, &dashboardToCreate); err != nil {
			log.Fatalln(err)
		}

		list, err := dashboardClient.List()
		if err != nil {
			log.Fatalln(err)
		}

		for _, dashboardFound := range list.Dashboards {
			if dashboardFound.Title == dashboardToCreate.Title {
				log.Fatalln("dashboard already exists")
			}
		}

		dashboardID, err := dashboardClient.Create(dashboardToCreate)
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("dashboard id: %s", dashboardID)
	},
}

func init() {
	initConfig()

	dashboardClient = api.NewDashboardClient(graylog)

	GetDashboardCmd.
		Flags().
		StringVarP(&dashboardID, "dashboard_id", "d", "", "id of dashboard to find")
	GetDashboardCmd.MarkFlagRequired("dashboard_id")
	if err := GetDashboardCmd.MarkFlagRequired("dashboard_id"); err != nil {
		log.Fatalln("no dashboard_id flag was found")
	}

	CreateDashboardCmd.
		Flags().
		StringVarP(&definitions, "config", "c", "", "config file containing the dashboard definition")
	CreateDashboardCmd.MarkFlagRequired("config")
	if err := CreateDashboardCmd.MarkFlagRequired("config"); err != nil {
		log.Fatalln("no config flag was found")
	}

	dashboardsCmd.AddCommand(ListDashboardsCmd)
	dashboardsCmd.AddCommand(GetDashboardCmd)
	dashboardsCmd.AddCommand(CreateDashboardCmd)

	rootCmd.AddCommand(dashboardsCmd)
}
