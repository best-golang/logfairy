package get

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/uniplaces/logfairy/helpers"
	"github.com/uniplaces/logfairy/infrastructure/api/widget"
)

func GetCommand(client widget.Client) *cobra.Command {
	var widgetID string
	var dashboardID string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "list a single widget",
		Long:  "widget get look up for a widget with the given widget id and print out the json.",
		Run:   getRunDefinition(client, widgetID, dashboardID),
	}

	cmd.
		Flags().
		StringVarP(&widgetID, "widget_id", "s", "", "id of widget to find")

	if err := cmd.MarkFlagRequired("widget_id"); err != nil {
		log.Fatalln("no widget_id flag was found")
	}

	cmd.
		Flags().
		StringVarP(&dashboardID, "dashboard_id", "d", "", "id of dashboard containing the widget")

	if err := cmd.MarkFlagRequired("dashboard_id"); err != nil {
		log.Fatalln("no dashboard_id flag was found")
	}

	return cmd
}

func getRunDefinition(
	client widget.Client,
	widgetID string,
	dashboardID string,
) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		widget, err := client.Get(widgetID, dashboardID)
		if err != nil {
			log.Fatalln(err)
		}

		helpers.JSONPrettyPrint(widget)
	}
}
