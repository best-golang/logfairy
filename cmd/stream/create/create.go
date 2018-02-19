package create

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
	"github.com/uniplaces/logfairy/dto/stream"
	sclient "github.com/uniplaces/logfairy/infrastructure/api/stream"
)

const long = `stream create will try to create the stream defined in the config file.
		
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

func GetCommand(client sclient.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "crate a stream",
		Long:  long,
		Run:   getRunDefinition(client),
	}

	cmd.
		Flags().
		StringVarP(&Definitions, "config", "c", "", "config file containing the stream definition")

	if err := cmd.MarkFlagRequired("config"); err != nil {
		log.Fatalln("no config flag was found")
	}

	return cmd
}

func getRunDefinition(client sclient.Client) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		definition, err := ioutil.ReadFile(Definitions)
		if err != nil {
			log.Fatalln(err)
		}

		streamToCreate := stream.Stream{}
		if err := json.Unmarshal(definition, &streamToCreate); err != nil {
			log.Fatalln(err)
		}

		streams, err := client.List()
		if err != nil {
			log.Fatalln(err)
		}

		if _, exists := streams.GetByTitle(streamToCreate.Title); exists {
			log.Fatalln("stream already exists")
		}

		streamID, err := client.Create(streamToCreate)
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("stream id: %s", streamID)
	}
}
