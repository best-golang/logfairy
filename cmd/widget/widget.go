package widget

import (
	"github.com/spf13/cobra"
	"github.com/uniplaces/logfairy/cmd/widget/create"
	"github.com/uniplaces/logfairy/cmd/widget/get"
	"github.com/uniplaces/logfairy/infrastructure/api/dashboard"
	"github.com/uniplaces/logfairy/infrastructure/api/widget"
)

func GetCommand(widgetClient widget.Client, dashboardClient dashboard.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "widget",
		Short: "handle widget actions",
		Long:  ``,
	}

	cmd.AddCommand(
		get.GetCommand(widgetClient),
		create.GetCommand(widgetClient, dashboardClient),
	)

	return cmd
}
