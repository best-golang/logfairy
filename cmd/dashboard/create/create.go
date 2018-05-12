package create

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
	"github.com/uniplaces/logfairy/dto/dashboard"
	api "github.com/uniplaces/logfairy/infrastructure/api/dashboard"
)

const long = `dashboard create will try to create the dashboard defined in the config file.
	
The expected json is:
{
  "title": "bar",
  "description": "description for bar dashboard"
}`

var Definitions string

func GetCommand(client api.Dashboard) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "crate a dashboard",
		Long:  long,
		Run:   getRunDefinition(client),
	}

	cmd.
		Flags().
		StringVarP(&Definitions, "config", "c", "", "config file containing the dashboard definition")

	if err := cmd.MarkFlagRequired("config"); err != nil {
		log.Fatalln("no config flag was found")
	}

	return cmd
}

func getRunDefinition(client api.Dashboard) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		definition, err := ioutil.ReadFile(Definitions)
		if err != nil {
			log.Fatalln(err)
		}

		dashboardToCreate := dashboard.Dashboard{}
		if err := json.Unmarshal(definition, &dashboardToCreate); err != nil {
			log.Fatalln(err)
		}

		dashboards, err := client.List()
		if err != nil {
			log.Fatalln(err)
		}

		if _, exists := dashboards.GetByTitle(dashboardToCreate.Title); exists {
			log.Fatalln("dashboard already exists")
		}

		dashboardID, err := client.Create(dashboardToCreate)
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("dashboard id: %s", dashboardID)
	}
}
