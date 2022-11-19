package cmd

import (
	"github.com/Rickylss/wib/cmd/wib/run"
	"github.com/Rickylss/wib/pkg/constants"
	"github.com/spf13/cobra"
)

func NewCreateCmd() *cobra.Command {
	flags := &run.CreateFlags{}
	cmd := &cobra.Command{
		Use:   "create [target_img]",
		Short: "create windows image",
		Long:  "create windows image",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			co, err := flags.NewCreateOptions(args)
			if err != nil {
				return err
			}

			return co.CreateImage()
		},
	}

	cmd.Flags().StringVarP(
		&flags.BaseImage,
		"base",
		"b",
		constants.DefaultBase,
		"base image to file",
	)
	cmd.MarkFlagRequired("base")

	cmd.Flags().StringVarP(
		&flags.ScriptsPath,
		"scripts",
		"s",
		constants.DefaultScriptsPath,
		"scripts path for win image.",
	)

	cmd.Flags().Int16VarP(
		&flags.VmStartTimeout,
		"timeout",
		"t",
		30,
		"timeout for vm start",
	)

	return cmd
}
