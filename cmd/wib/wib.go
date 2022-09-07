package wib

import (
	"os"

	"github.com/Rickylss/wib/cmd/wib/cmd"
	log "github.com/sirupsen/logrus"
)

func Main() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "15:04:05",
	})
	if err := Run(); err != nil {
		os.Exit(1)
	}
}

func Run() error {
	return cmd.NewWinImgBuilderCmd().Execute()
}
