package run

import (
	"github.com/Rickylss/wib/pkg/constants"
	log "github.com/sirupsen/logrus"
)

type Flags struct {
	LogLevel string
}

func (flags *Flags) SetLogLevel() error {
	level := constants.DefaultLogLevel
	parsed, err := log.ParseLevel(flags.LogLevel)
	if err != nil {
		log.Warnf("Invalid log level '%s', defaulting to '%s'", flags.LogLevel, level)
	} else {
		level = parsed
	}
	log.SetLevel(level)
	return nil
}
