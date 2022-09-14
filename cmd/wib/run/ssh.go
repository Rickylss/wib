package run

import (
	"strings"

	"github.com/Rickylss/wib/pkg/vm"
	log "github.com/sirupsen/logrus"
)

type SshFlags struct {
	Ip       string
	Port     int
	User     string
	Password string
	Key      string
}

type SshOptions struct {
	*SshFlags
	*vm.VM
	Cmd string
}

func (sf *SshFlags) NewSshOptions(args []string) (so *SshOptions, err error) {

	v := &vm.VM{
		SshContext: &vm.SshContext{
			User:     sf.User,
			Password: sf.Password,
			Port:     sf.Port,
			Ip:       sf.Ip,
		},
	}

	so = &SshOptions{
		SshFlags: sf,
		VM:       v,
		Cmd:      args[0],
	}

	return
}

func (so *SshOptions) Run() error {
	out, err := so.VM.Output(so.Cmd)
	if err != nil {
		return err
	}

	log.Infof(strings.TrimSpace(string(out)))

	return nil
}
