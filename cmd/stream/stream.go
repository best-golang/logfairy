package stream

import (
	"github.com/spf13/cobra"
	"github.com/uniplaces/logfairy/cmd/stream/create"
	"github.com/uniplaces/logfairy/cmd/stream/get"
	"github.com/uniplaces/logfairy/cmd/stream/list"
	"github.com/uniplaces/logfairy/infrastructure/api/stream"
)

func GetCommand(client stream.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stream",
		Short: "handle stream actions",
		Long:  "",
	}

	cmd.AddCommand(
		list.GetCommand(client),
		get.GetCommand(client),
		create.GetCommand(client),
	)

	return cmd
}
