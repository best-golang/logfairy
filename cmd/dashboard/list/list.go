package list

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/uniplaces/logfairy/helpers"
	"github.com/uniplaces/logfairy/infrastructure/api/dashboard"
)

func GetCommand(client dashboard.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "list dashboards",
		Long:  "dashboard list look up for all the dashboards it is allowed to reach.",
		Run:   getRunDefinition(client),
	}
}

func getRunDefinition(client dashboard.Client) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		dashboards, err := client.List()
		if err != nil {
			log.Fatalln(err)
		}

		helpers.JSONPrettyPrint(dashboards)
	}
}
