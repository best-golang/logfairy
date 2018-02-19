package create

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
	"github.com/uniplaces/logfairy/dto/dashboard"
	dclient "github.com/uniplaces/logfairy/infrastructure/api/dashboard"
)

const long = `dashboard create will try to create the dashboard defined in the config file.
	
The expected json is:
{
  "title": "bar",
  "description": "description for bar dashboard"
}`

func GetCommand(client dclient.Client) *cobra.Command {
	var definitions string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "crate a dashboard",
		Long:  long,
		Run:   getRunDefinition(client, definitions),
	}

	cmd.
		Flags().
		StringVarP(&definitions, "config", "c", "", "config file containing the dashboard definition")

	if err := cmd.MarkFlagRequired("config"); err != nil {
		log.Fatalln("no config flag was found")
	}

	return cmd
}

func getRunDefinition(client dclient.Client, definitions string) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		definition, err := ioutil.ReadFile(definitions)
		if err != nil {
			log.Fatalln(err)
		}

		dashboardToCreate := dashboard.Dashboard{}
		if err := json.Unmarshal(definition, &dashboardToCreate); err != nil {
			log.Fatalln(err)
		}

		list, err := client.List()
		if err != nil {
			log.Fatalln(err)
		}

		for _, dashboardFound := range list.Dashboards {
			if dashboardFound.Title == dashboardToCreate.Title {
				log.Fatalln("dashboard already exists")
			}
		}

		dashboardID, err := client.Create(dashboardToCreate)
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("dashboard id: %s", dashboardID)
	}
}
