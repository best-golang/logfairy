package update

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
	"github.com/uniplaces/logfairy/dto/stream"
	api "github.com/uniplaces/logfairy/infrastructure/api/stream"
)

const long = `stream update will try to update the stream defined in the config file.
		
The expected json is:
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
}`

var Definitions string
var StreamID string

func GetCommand(client api.Stream) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "update a stream",
		Long:  long,
		Run:   getRunDefinition(client),
	}

	cmd.
		Flags().
		StringVarP(&Definitions, "config", "c", "", "config file containing the stream definition")

	if err := cmd.MarkFlagRequired("config"); err != nil {
		log.Fatalln("no config flag was found")
	}

	cmd.
		Flags().
		StringVarP(&StreamID, "stream_id", "s", "", "id of stream to update")

	if err := cmd.MarkFlagRequired("stream_id"); err != nil {
		log.Fatalln("no id flag was found")
	}

	return cmd
}

func getRunDefinition(client api.Stream) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		definition, err := ioutil.ReadFile(Definitions)
		if err != nil {
			log.Fatalln(err)
		}

		streamToUpdate := stream.Stream{}
		if err := json.Unmarshal(definition, &streamToUpdate); err != nil {
			log.Fatalln(err)
		}

		streamID, err := client.Update(StreamID, streamToUpdate)
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("stream id: %s", streamID)
	}
}
