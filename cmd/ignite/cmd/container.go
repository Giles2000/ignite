package cmd

import (
	"github.com/luxas/ignite/pkg/errutils"
	"github.com/spf13/cobra"
	"io"
)

// NewContainerCmd runs the dhcp server and sets up routing inside Docker
func NewCmdContainer(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
			err := RunContainer(out, cmd, args)
			errutils.Check(err)
		},
	}

	//cmd.Flags().StringP("output", "o", "", "Output format; available options are 'yaml', 'json' and 'short'")
	return cmd
}

// RunBuild runs when the Container command is invoked
func RunContainer(out io.Writer, cmd *cobra.Command, args []string) error {

	return nil
}
