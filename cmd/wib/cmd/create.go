package cmd

import (
	"github.com/Rickylss/wib/cmd/wib/run"
	"github.com/spf13/cobra"
)

func NewCreateCmd() *cobra.Command {
	flags := &run.CreateFlags{}
	cmd := &cobra.Command{
		Use:   "create [base_img]",
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
		&flags.OutPut,
		"output",
		"o",
		run.DefaultOutPut,
		"output image to file",
	)

	cmd.Flags().StringVarP(
		&flags.Size,
		"size",
		"s",
		run.DefaultSize,
		"create image size",
	)

	cmd.Flags().BoolVar(
		&flags.Release,
		"release",
		false,
		"create release image",
	)

	return cmd
}
