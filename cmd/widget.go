package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/uniplaces/logfairy/dto/dashboard"
	"github.com/uniplaces/logfairy/infrastructure/api"
)

var (
	widgetID          string
	dashboardWidgetID string
)

// widgetCmd represents the widgets command
var widgetCmd = &cobra.Command{
	Use:   "widget",
	Short: "handle widgets actions",
	Long:  ``,
}

// GetWidgetCmd represents the list single widget command
var GetWidgetCmd = &cobra.Command{
	Use:   "get",
	Short: "list a single widget",
	Long:  `widget get look up for a widget with the given widget id and print out the json.`,
	Run: func(cmd *cobra.Command, args []string) {
		widget, err := widgetClient.Get(widgetID, dashboardWidgetID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		prettyPrint(widget)
	},
}

// CreateWidgetCmd represents the create command
var CreateWidgetCmd = &cobra.Command{
	Use:   "create",
	Short: "crate a widget",
	Long: `widget create will try to create the widget defined in the config file.
	
The expected structure is:
{
  "cache_time": 1800,
  "description": "foo foo bar bar",
  "type": "STREAM_SEARCH_RESULT_COUNT",
  "config": {
    "timerange": {
      "type": "relative",
      "range": 86400
    },
    "lower_is_better": false,
    "stream_id": d87d7afd-89c5-4233-a20b-4476785f11cb,
    "trend": true,
    "query": "foo bar"
  }
}`,
	Run: func(cmd *cobra.Command, args []string) {
		definition, err := ioutil.ReadFile(definitions)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		widgetToCreate := dashboard.Widget{}
		if err := json.Unmarshal(definition, &widgetToCreate); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		dashboard, err := dashboardClient.Get(dashboardWidgetID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, widgetFound := range dashboard.Widgets {
			if widgetFound.Description == widgetToCreate.Description {
				fmt.Println("widget already exists")
				os.Exit(1)
			}
		}

		widgetID, err := widgetClient.Create(widgetToCreate, dashboardWidgetID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("widget id: %s", widgetID)
	},
}

func init() {
	initConfig()

	widgetClient = api.NewWidgetClient(graylog)
	dashboardClient = api.NewDashboardClient(graylog)

	GetWidgetCmd.Flags().StringVarP(&dashboardWidgetID, "dashboard_id", "d", "", "id of dashboard containing the widget")
	GetWidgetCmd.MarkFlagRequired("dashboard_id")
	GetWidgetCmd.Flags().StringVarP(&widgetID, "widget_id", "w", "", "id of widget to find")
	GetWidgetCmd.MarkFlagRequired("widget_id")

	CreateWidgetCmd.Flags().StringVarP(&dashboardWidgetID, "dashboard_id", "d", "", "id of dashboard containing the widget")
	CreateWidgetCmd.MarkFlagRequired("dashboard_id")
	CreateWidgetCmd.Flags().StringVarP(&definitions, "config", "c", "", "config file containing the widget definition")
	CreateWidgetCmd.MarkFlagRequired("config")

	widgetCmd.AddCommand(GetWidgetCmd)
	widgetCmd.AddCommand(CreateWidgetCmd)

	rootCmd.AddCommand(widgetCmd)
}
