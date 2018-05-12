package create

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
	"github.com/uniplaces/logfairy/dto/dashboard"
	dapi "github.com/uniplaces/logfairy/infrastructure/api/dashboard"
	wapi "github.com/uniplaces/logfairy/infrastructure/api/widget"
)

const long = `widget create will try to create the widget defined in the config file.
	
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
}`

var (
	DashboardWidgetID string
	Definitions       string
)

func GetCommand(widgetClient wapi.Widget, dashboardClient dapi.Dashboard) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "crate a widget",
		Long:  long,
		Run:   getRunDefinition(widgetClient, dashboardClient),
	}

	cmd.
		Flags().
		StringVarP(&DashboardWidgetID, "dashboard_id", "d", "", "id of dashboard containing the widget")

	if err := cmd.MarkFlagRequired("dashboard_id"); err != nil {
		log.Fatalln("no dashboard_id flag was found")
	}

	cmd.
		Flags().
		StringVarP(&Definitions, "config", "c", "", "config file containing the widget definition")

	if err := cmd.MarkFlagRequired("config"); err != nil {
		log.Fatalln("no config flag was found")
	}

	return cmd
}

func getRunDefinition(
	widgetClient wapi.Widget,
	dashboardClient dapi.Dashboard,
) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		definition, err := ioutil.ReadFile(Definitions)
		if err != nil {
			log.Fatalln(err)
		}

		widgetToCreate := dashboard.Widget{}
		if err := json.Unmarshal(definition, &widgetToCreate); err != nil {
			log.Fatalln(err)
		}

		dashboard, err := dashboardClient.Get(DashboardWidgetID)
		if err != nil {
			log.Fatalln(err)
		}

		if _, exists := dashboard.GetByDescription(widgetToCreate.Description); exists {
			log.Fatalln("widget already exists")
		}

		widgetID, err := widgetClient.Create(widgetToCreate, DashboardWidgetID)
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("widget id: %s", widgetID)
	}
}
