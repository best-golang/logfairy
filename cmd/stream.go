package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/uniplaces/logfairy/dto/stream"
	"github.com/uniplaces/logfairy/infrastructure/api"
)

var streamID string

// streamsCmd represents the streams command
var streamsCmd = &cobra.Command{
	Use:   "stream",
	Short: "handle stream actions",
	Long:  ``,
}

// ListStreamsCmd represents the List command
var ListStreamsCmd = &cobra.Command{
	Use:   "list",
	Short: "list streams",
	Long:  `stream list look up for all the streams it is allowed to reach.`,
	Run: func(cmd *cobra.Command, args []string) {
		streams, err := streamClient.List()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		prettyPrint(streams)
	},
}

// GetStreamCmd represents the list single stream command
var GetStreamCmd = &cobra.Command{
	Use:   "get",
	Short: "list a single stream",
	Long:  `stream get look up for a stream with the given stream id and print out the json.`,
	Run: func(cmd *cobra.Command, args []string) {
		stream, err := streamClient.Get(streamID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		prettyPrint(stream)
	},
}

// CreateStreamCmd represents the create command
var CreateStreamCmd = &cobra.Command{
	Use:   "create",
	Short: "crate a stream",
	Long: `stream create will try to create the stream defined in the config file.
	
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
}`,
	Run: func(cmd *cobra.Command, args []string) {
		definition, err := ioutil.ReadFile(definitions)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		streamToCreate := stream.Stream{}
		if err := json.Unmarshal(definition, &streamToCreate); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		list, err := streamClient.List()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, streamFound := range list.Streams {
			if streamFound.Title == streamToCreate.Title {
				fmt.Println("stream already exists")
				os.Exit(1)
			}
		}

		streamID, err := streamClient.Create(streamToCreate)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("stream id: %s", streamID)
	},
}

func init() {
	streamClient = api.NewStreamClient(graylog)

	GetStreamCmd.Flags().StringVarP(&streamID, "stream_id", "s", "", "id of stream to find")
	GetStreamCmd.MarkFlagRequired("stream_id")

	CreateStreamCmd.Flags().StringVarP(&definitions, "config", "c", "", "config file containing the stream definition")
	CreateStreamCmd.MarkFlagRequired("config")

	streamsCmd.AddCommand(ListStreamsCmd)
	streamsCmd.AddCommand(GetStreamCmd)
	streamsCmd.AddCommand(CreateStreamCmd)

	rootCmd.AddCommand(streamsCmd)
}
