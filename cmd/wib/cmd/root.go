package cmd

import (
	"github.com/Rickylss/wib/cmd/wib/run"
	"github.com/Rickylss/wib/pkg/constants"
	"github.com/spf13/cobra"
)

func NewWinImgBuilderCmd() *cobra.Command {
	flags := &run.Flags{}
	root := &cobra.Command{
		Use:   "wib",
		Short: "windows vm image builder",
		Long: "        ._____.      \n" +
			"__  _  _|__\\_ |__     \n" +
			"\\ \\/ \\/ /  || __ \\ \n" +
			" \\     /|  || \\_\\ \\\n" +
			"  \\/\\_/ |__||___  /  \n" +
			"                \\/  \n\n" +
			"windows vm image builder",
		Version:      constants.Version,
		SilenceUsage: true,
		Args:         cobra.NoArgs,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return flags.SetLogLevel()
		},
	}

	root.PersistentFlags().StringVarP(
		&flags.LogLevel,
		"loglevel",
		"l",
		run.DefaultLogLevel.String(),
		"logrus log level [panic, fatal, error, warning, info, debug, trace]",
	)

	root.AddCommand(NewCreateCmd())

	return root
}
