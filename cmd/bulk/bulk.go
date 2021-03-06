package bulk

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
	"github.com/uniplaces/logfairy/dto/bulk"
	"github.com/uniplaces/logfairy/dto/dashboard"
	"github.com/uniplaces/logfairy/dto/stream"
	dapi "github.com/uniplaces/logfairy/infrastructure/api/dashboard"
	sapi "github.com/uniplaces/logfairy/infrastructure/api/stream"
	wapi "github.com/uniplaces/logfairy/infrastructure/api/widget"
)

const (
	bulkLong = `bulk will try to create all the definitions in the config file.
	
	The expected json is:
	{
		"streams": {
			"foo_hash": {
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
		},
		"dashboards": [
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
							"stream_id": "foo_hash",
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
	the position definied for the value stream_id.`
)

var Definitions string

func GetCommand(
	streamClient sapi.Stream,
	dashboardClient dapi.Dashboard,
	widgetClient wapi.Widget,
) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bulk",
		Short: "bulk allows to create a set of streams, dashboards and widgets",
		Long:  bulkLong,
		Run:   getRunDefinition(streamClient, dashboardClient, widgetClient),
	}

	cmd.
		Flags().
		StringVarP(&Definitions, "config", "c", "", "config file containing the definitions")

	if err := cmd.MarkFlagRequired("config"); err != nil {
		log.Fatalln("no config flag was found")
	}

	return cmd
}

func getRunDefinition(
	streamClient sapi.Stream,
	dashboardClient dapi.Dashboard,
	widgetClient wapi.Widget,
) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		definition, err := ioutil.ReadFile(Definitions)
		if err != nil {
			log.Fatalln(err)
		}

		bulkObject := bulk.Bulk{}
		if err := json.Unmarshal(definition, &bulkObject); err != nil {
			log.Fatalln(err)
		}

		updatedStream := make(map[string]stream.Stream, len(bulkObject.Streams))
		for index, streamToCreate := range bulkObject.Streams {
			if err := createStream(streamClient, &streamToCreate); err != nil {
				log.Fatalln(err)
			}

			updatedStream[index] = streamToCreate
		}

		for _, dashboardToCreate := range bulkObject.Dashboards {
			if err := createDashboard(dashboardClient, widgetClient, &dashboardToCreate, updatedStream); err != nil {
				log.Fatalln(err)
			}
		}
	}
}

func createStream(streamClient sapi.Stream, streamToCreate *stream.Stream) error {
	streams, err := streamClient.List()
	if err != nil {
		log.Fatalln(err)
	}

	if foundStream, exists := streams.GetByTitle(streamToCreate.Title); exists {
		streamToCreate.ID = foundStream.ID

		return nil
	}

	streamID, err := streamClient.Create(*streamToCreate)
	if err != nil {
		return err
	}

	streamToCreate.ID = &streamID

	return nil
}

func createDashboard(
	dashboardClient dapi.Dashboard,
	widgetClient wapi.Widget,
	dashboardToCreate *dashboard.Dashboard,
	streams map[string]stream.Stream,
) error {
	dashboards, err := dashboardClient.List()
	if err != nil {
		log.Fatalln(err)
	}

	widgets := extractWidgets(dashboardToCreate)
	dashboardID, err := getDashboardID(dashboardClient, dashboards, *dashboardToCreate)
	if err != nil {
		return err
	}

	for _, widgetToCreate := range widgets {
		if err := createWidget(dashboardClient, widgetClient, widgetToCreate, dashboardID, streams); err != nil {
			return err
		}
	}

	return nil
}

func getDashboardID(
	dashboardClient dapi.Dashboard,
	dashboards dashboard.Dashboards,
	dashboardToCreate dashboard.Dashboard,
) (string, error) {
	foundDashboard, exists := dashboards.GetByTitle(dashboardToCreate.Title)
	if !exists {
		return dashboardClient.Create(dashboardToCreate)
	}

	return *foundDashboard.ID, nil
}

func createWidget(
	dashboardClient dapi.Dashboard,
	widgetClient wapi.Widget,
	widgetToCreate dashboard.Widget,
	dashboardID string,
	streams map[string]stream.Stream,
) error {
	dashboard, err := dashboardClient.Get(dashboardID)
	if err != nil {
		log.Fatalln(err)
	}

	if _, exists := dashboard.GetByDescription(widgetToCreate.Description); exists {
		return nil
	}

	widgetToCreate.Config.StreamID = *streams[widgetToCreate.Config.StreamID].ID
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
