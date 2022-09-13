package constants

import (
	log "github.com/sirupsen/logrus"
)

const (
	Version = "v0.1"
	WorkDir = "/etc/wib"
)

const (
	DefaultLogLevel    = log.WarnLevel
	DefaultBase        = "win.img"
	DefaultScriptsPath = WorkDir + "scripts.d"
)
