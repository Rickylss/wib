package cmd

import (
	"github.com/Rickylss/wib/cmd/wib/run"
	"github.com/spf13/cobra"
)

func NewSshCmd() *cobra.Command {
	flags := &run.SshFlags{}
	cmd := &cobra.Command{
		Use:   "ssh [target_img]",
		Short: "ssh windows image",
		Long:  "ssh windows image",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			so, err := flags.NewSshOptions(args)
			if err != nil {
				return err
			}

			return so.Run()
		},
	}

	cmd.Flags().StringVarP(
		&flags.Ip,
		"ip",
		"i",
		"",
		"ip of vm",
	)

	cmd.Flags().IntVarP(
		&flags.Port,
		"port",
		"P",
		22,
		"port of vm",
	)

	cmd.Flags().StringVarP(
		&flags.User,
		"user",
		"u",
		"",
		"user of vm",
	)

	cmd.Flags().StringVarP(
		&flags.Password,
		"password",
		"p",
		"",
		"password of vm",
	)

	cmd.Flags().StringVarP(
		&flags.Key,
		"key",
		"k",
		"",
		"key of vm",
	)

	return cmd
}
