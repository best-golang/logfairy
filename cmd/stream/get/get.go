package get

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/uniplaces/logfairy/helpers"
	"github.com/uniplaces/logfairy/infrastructure/api/stream"
)

var StreamID string

func GetCommand(client stream.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "list a single stream",
		Long:  "stream get look up for a stream with the given stream id and print out the json.",
		Run:   getRunDefinition(client),
	}

	cmd.
		Flags().
		StringVarP(&StreamID, "stream_id", "s", "", "id of stream to find")

	if err := cmd.MarkFlagRequired("stream_id"); err != nil {
		log.Fatalln("no stream_id flag was found")
	}

	return cmd
}

func getRunDefinition(client stream.Client) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		stream, err := client.Get(StreamID)
		if err != nil {
			log.Fatalln(err)
		}

		helpers.JSONPrettyPrint(stream)
	}
}
