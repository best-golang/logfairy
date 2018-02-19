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

func GetCommand(client sclient.Client) *cobra.Command {
	var definitions string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "crate a stream",
		Long:  long,
		Run:   getRunDefinition(client, definitions),
	}

	cmd.
		Flags().
		StringVarP(&definitions, "config", "c", "", "config file containing the stream definition")

	if err := cmd.MarkFlagRequired("config"); err != nil {
		log.Fatalln("no config flag was found")
	}

	return cmd
}

func getRunDefinition(client sclient.Client, definitions string) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		definition, err := ioutil.ReadFile(definitions)
		if err != nil {
			log.Fatalln(err)
		}

		streamToCreate := stream.Stream{}
		if err := json.Unmarshal(definition, &streamToCreate); err != nil {
			log.Fatalln(err)
		}

		list, err := client.List()
		if err != nil {
			log.Fatalln(err)
		}

		for _, streamFound := range list.Streams {
			if streamFound.Title == streamToCreate.Title {
				log.Fatalln("stream already exists")
			}
		}

		streamID, err := client.Create(streamToCreate)
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("stream id: %s", streamID)
	}
}
