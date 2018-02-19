package create

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
	"github.com/uniplaces/logfairy/dto/dashboard"
	dclient "github.com/uniplaces/logfairy/infrastructure/api/dashboard"
	"github.com/uniplaces/logfairy/infrastructure/api/widget"
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

func GetCommand(widgetClient widget.Client, dashboardClient dclient.Client) *cobra.Command {
	var dashboardWidgetID string
	var definitions string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "crate a widget",
		Long:  long,
		Run:   getRunDefinition(widgetClient, dashboardClient, definitions, dashboardWidgetID),
	}

	cmd.
		Flags().
		StringVarP(&dashboardWidgetID, "dashboard_id", "d", "", "id of dashboard containing the widget")

	if err := cmd.MarkFlagRequired("dashboard_id"); err != nil {
		log.Fatalln("no dashboard_id flag was found")
	}

	cmd.
		Flags().
		StringVarP(&definitions, "config", "c", "", "config file containing the widget definition")

	if err := cmd.MarkFlagRequired("config"); err != nil {
		log.Fatalln("no config flag was found")
	}

	return cmd
}

func getRunDefinition(
	widgetClient widget.Client,
	dashboardClient dclient.Client,
	definitions string,
	dashboardWidgetID string,
) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		definition, err := ioutil.ReadFile(definitions)
		if err != nil {
			log.Fatalln(err)
		}

		widgetToCreate := dashboard.Widget{}
		if err := json.Unmarshal(definition, &widgetToCreate); err != nil {
			log.Fatalln(err)
		}

		dashboard, err := dashboardClient.Get(dashboardWidgetID)
		if err != nil {
			log.Fatalln(err)
		}

		for _, widgetFound := range dashboard.Widgets {
			if widgetFound.Description == widgetToCreate.Description {
				log.Fatalln("widget already exists")
			}
		}

		widgetID, err := widgetClient.Create(widgetToCreate, dashboardWidgetID)
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("widget id: %s", widgetID)
	}
}
