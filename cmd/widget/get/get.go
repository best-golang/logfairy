package get

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/uniplaces/logfairy/helpers"
	"github.com/uniplaces/logfairy/infrastructure/api/widget"
)

var (
	WidgetID    string
	DashboardID string
)

func GetCommand(client widget.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "list a single widget",
		Long:  "widget get look up for a widget with the given widget id and print out the json.",
		Run:   getRunDefinition(client),
	}

	cmd.
		Flags().
		StringVarP(&WidgetID, "widget_id", "w", "", "id of widget to find")

	if err := cmd.MarkFlagRequired("widget_id"); err != nil {
		log.Fatalln("no widget_id flag was found")
	}

	cmd.
		Flags().
		StringVarP(&DashboardID, "dashboard_id", "d", "", "id of dashboard containing the widget")

	if err := cmd.MarkFlagRequired("dashboard_id"); err != nil {
		log.Fatalln("no dashboard_id flag was found")
	}

	return cmd
}

func getRunDefinition(client widget.Client) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		widget, err := client.Get(WidgetID, DashboardID)
		if err != nil {
			log.Fatalln(err)
		}

		helpers.JSONPrettyPrint(widget)
	}
}
