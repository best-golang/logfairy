package dashboard

import (
	"github.com/spf13/cobra"
	"github.com/uniplaces/logfairy/cmd/dashboard/create"
	"github.com/uniplaces/logfairy/cmd/dashboard/get"
	"github.com/uniplaces/logfairy/cmd/dashboard/list"
	"github.com/uniplaces/logfairy/infrastructure/api/dashboard"
)

func GetCommand(client dashboard.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dashboard",
		Short: "handle dashboard actions",
		Long:  ``,
	}

	cmd.AddCommand(
		list.GetCommand(client),
		get.GetCommand(client),
		create.GetCommand(client),
	)

	return cmd
}
