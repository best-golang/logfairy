package cmd

import (
	"github.com/spf13/cobra"
	"github.com/uniplaces/logfairy/cmd/bulk"
	"github.com/uniplaces/logfairy/cmd/dashboard"
	"github.com/uniplaces/logfairy/cmd/stream"
	"github.com/uniplaces/logfairy/cmd/widget"
)

func GetCommand(dependencies Dependencies) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logfairy",
		Short: "logfairy to standardize interaction with graylog",
		Long:  "logfairy goal is to standardize the way streams, dashboard and widget are created",
	}

	cmd.AddCommand(
		stream.GetCommand(dependencies.StreamClient),
		widget.GetCommand(dependencies.WidgetClient, dependencies.DashboardClient),
		dashboard.GetCommand(dependencies.DashboardClient),
		bulk.GetCommand(dependencies.StreamClient, dependencies.DashboardClient, dependencies.WidgetClient),
	)

	return cmd
}
