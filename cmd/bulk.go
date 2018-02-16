package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/uniplaces/logfairy/dto/bulk"
	"github.com/uniplaces/logfairy/dto/dashboard"
	"github.com/uniplaces/logfairy/dto/stream"
	"github.com/uniplaces/logfairy/infrastructure/api"
)

// bulkCmd represents the bulk command
var bulkCmd = &cobra.Command{
	Use:   "bulk",
	Short: "bulk allows to create a set of streams, dashboards and widgets",
	Long: `bulk will try to create all the definitions in the config file.
	
The expected json is:
{
  "streams": [
    {
      "title": "foo",
      "description": "description for foo",
      "matching_type": "AND",
      "rules": [
        {
          "field": "_env",
          "description": "",
          "type": 1,
          "inverted": false,
          "value": "prod"
        },
        {
          "field": "_app-id",
          "description": "",
          "type": 1,
          "inverted": false,
          "value": "foo"
        }
      ],
      "remove_matches_from_default_stream": false,
      "index_set_id": "5b0bfb3bgg857f3b700b58g5"
    }
  ],
  "dashboard": [
    {
      "title": "bar",
      "description": "description for bar dashboard",
      "widgets": [
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
            "stream_id": "0",
            "trend": true,
            "query": "foo bar"
          }
        }
      ]
    }
  ]
}

Please notice that when creating dashboards graylog doesn't accept widgets, 
behind the scene the widgets are extracted from the dashboard structure and 
build it one by one replacesing the stream_id value with the stream id in 
the position definied for the value stream_id.`,
	Run: func(cmd *cobra.Command, args []string) {
		definition, err := ioutil.ReadFile(definitions)
		if err != nil {
			log.Fatalln(err)
		}

		bulkObject := bulk.Bulk{}
		if err := json.Unmarshal(definition, &bulkObject); err != nil {
			log.Fatalln(err)
		}

		updatedStream := make([]stream.Stream, len(bulkObject.Streams))
		for index, streamToCreate := range bulkObject.Streams {
			if err := createStream(&streamToCreate); err != nil {
				log.Fatalln(err)
			}

			updatedStream[index] = streamToCreate
		}

		for _, dashboardToCreate := range bulkObject.Dashboards {
			if err := createDashboard(&dashboardToCreate, updatedStream); err != nil {
				log.Fatalln(err)
			}
		}
	},
}

func init() {
	widgetClient = api.NewWidgetClient(graylog)
	dashboardClient = api.NewDashboardClient(graylog)
	streamClient = api.NewStreamClient(graylog)

	bulkCmd.
		Flags().
		StringVarP(&definitions, "config", "c", "", "config file containing the definitions")
	if err := bulkCmd.MarkFlagRequired("config"); err != nil {
		log.Fatalln("no config flag was found")
	}

	rootCmd.AddCommand(bulkCmd)
}

func createStream(streamToCreate *stream.Stream) error {
	streamID, err := streamClient.Create(*streamToCreate)
	if err != nil {
		return err
	}

	streamToCreate.ID = &streamID

	return nil
}

func createDashboard(dashboardToCreate *dashboard.Dashboard, streams []stream.Stream) error {
	widgets := extractWidgets(dashboardToCreate)
	dashboardID, err := dashboardClient.Create(*dashboardToCreate)
	if err != nil {
		return err
	}

	for _, widgetToCreate := range widgets {
		if err := createWidget(widgetToCreate, dashboardID, streams); err != nil {
			return err
		}
	}

	return nil
}

func createWidget(widgetToCreate dashboard.Widget, dashboardID string, streams []stream.Stream) error {
	streamPosition, err := strconv.Atoi(widgetToCreate.Config.StreamID)
	if err != nil {
		return err
	}

	widgetToCreate.Config.StreamID = *streams[streamPosition].ID
	widgetToCreate.Config.StreamID, err = widgetClient.Create(widgetToCreate, dashboardID)

	return err
}

func extractWidgets(dashboardToCreate *dashboard.Dashboard) []dashboard.Widget {
	if dashboardToCreate.Widgets == nil {
		return []dashboard.Widget{}
	}

	widgets := dashboardToCreate.Widgets
	dashboardToCreate.Widgets = nil

	return widgets
}
