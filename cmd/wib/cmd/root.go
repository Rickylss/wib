package cmd

import (
	"fmt"
	"os"

	"github.com/Rickylss/wib/cmd/wib/run"
	"github.com/Rickylss/wib/pkg/constants"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

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
		&cfgFile,
		"config",
		"c",
		"",
		`config file (default "$Home/.wib.yaml")`,
	)

	root.PersistentFlags().StringVarP(
		&flags.LogLevel,
		"loglevel",
		"l",
		constants.DefaultLogLevel.String(),
		"logrus log level [panic, fatal, error, warning, info, debug, trace]",
	)
	viper.BindPFlag("loglevel", root.PersistentFlags().Lookup("loglevel"))

	root.AddCommand(NewCreateCmd())
	root.AddCommand(NewSshCmd())
	root.AddCommand(NewResizeCmd())
	root.AddCommand(NewUploadCmd())

	return root
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".wib")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("read failed: %s", err)
		os.Exit(1)
	}
}
