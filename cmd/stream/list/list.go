package list

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/uniplaces/logfairy/helpers"
	"github.com/uniplaces/logfairy/infrastructure/api/stream"
)

func GetCommand(client stream.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "list streams",
		Long:  "stream list look up for all the streams it is allowed to reach.",
		Run:   getRunDefinition(client),
	}
}

func getRunDefinition(client stream.Client) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		streams, err := client.List()
		if err != nil {
			log.Fatalln(err)
		}

		helpers.JSONPrettyPrint(streams)
	}
}
