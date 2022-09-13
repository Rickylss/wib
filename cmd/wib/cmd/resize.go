package cmd

import "github.com/spf13/cobra"

func NewResizeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resize [image] [size]",
		Short: "resize windows image",
		Long:  "resize windows image",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	return cmd
}
