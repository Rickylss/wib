package run

import (
	"fmt"
	"time"

	"github.com/Rickylss/wib/pkg/virt"
	"github.com/Rickylss/wib/pkg/vm"
	"github.com/Rickylss/wib/pkg/xml"

	log "github.com/sirupsen/logrus"
)

type CreateFlags struct {
	BaseImage      string
	ScriptsPath    string
	VmStartTimeout int16
}

type CreateOptions struct {
	*CreateFlags
	*vm.VM
	TargetImage string
}

func (cf *CreateFlags) NewCreateOptions(args []string) (co *CreateOptions, err error) {
	if args[0] == "" {
		err = fmt.Errorf("please specific a target image")
		return
	}

	target := args[0]

	virtManager := virt.NewVirtManager("")

	domxml, err := xml.GetDefaultXML(target, cf.BaseImage)
	if err != nil {
		return
	}

	vm, err := virtManager.StartVm(domxml)
	if err != nil {
		return
	}

	vm.SshContext.User = "Admin"
	vm.SshContext.Password = "password"
	vm.SshContext.Port = 22
	vm.SshContext.Key = ""

	co = &CreateOptions{
		CreateFlags: cf,
		TargetImage: args[0],
		VM:          vm,
	}

	return
}

func (co *CreateOptions) CreateImage() error {
	// apply scripts under /etc/wib/scripts.d in order
	log.Infof("ip:%s, port:%s, passwd:%s", co.VM.Ip, co.VM.Port, co.VM.Password)

	sleepTime := co.VmStartTimeout / 5
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * time.Duration(sleepTime))
		out, err := co.VM.Output("systeminfo")
		if err != nil {
			continue
		}
		log.Infof(string(out))
		break
	}

	// stop vm destroy vm

	// convert snapshot image

	return nil
}
