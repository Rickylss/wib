package cmd

import "github.com/spf13/cobra"

func NewUploadCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upload [image] [target]",
		Short: "upload windows image",
		Long:  "upload windows image",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	return cmd
}
