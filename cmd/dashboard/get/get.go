package get

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/uniplaces/logfairy/helpers"
	"github.com/uniplaces/logfairy/infrastructure/api/dashboard"
)

var DashboardID string

func GetCommand(client dashboard.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "list a single dashboard",
		Long:  "dashboard get look up for a dashboard with the given dashboard id and print out the json.",
		Run:   getRunDefinition(client),
	}

	cmd.
		Flags().
		StringVarP(&DashboardID, "dashboard_id", "d", "", "id of dashboard to find")

	if err := cmd.MarkFlagRequired("dashboard_id"); err != nil {
		log.Fatalln("no dashboard_id flag was found")
	}

	return cmd
}

func getRunDefinition(client dashboard.Client) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		dashboard, err := client.Get(DashboardID)
		if err != nil {
			log.Fatalln(err)
		}

		helpers.JSONPrettyPrint(dashboard)
	}
}
