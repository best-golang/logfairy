package resume

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/uniplaces/logfairy/helpers"
	api "github.com/uniplaces/logfairy/infrastructure/api/stream"
)

var StreamID string

func GetCommand(client api.Stream) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resume",
		Short: "resume a stream",
		Long:  "stream resume tryes to start a stream for the given stream_id",
		Run:   getRunDefinition(client),
	}

	cmd.
		Flags().
		StringVarP(&StreamID, "stream_id", "s", "", "id of stream to resume")

	if err := cmd.MarkFlagRequired("stream_id"); err != nil {
		log.Fatalln("no stream_id flag was found")
	}

	return cmd
}

func getRunDefinition(client api.Stream) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		stream, err := client.Resume(StreamID)
		if err != nil {
			log.Fatalln(err)
		}

		helpers.JSONPrettyPrint(stream)
	}
}
