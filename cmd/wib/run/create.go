package run

import (
	"os"
	"time"

	"github.com/Rickylss/wib/pkg/virt"
	"github.com/Rickylss/wib/pkg/vm"
	"github.com/Rickylss/wib/pkg/xml"

	log "github.com/sirupsen/logrus"
)

type CreateFlags struct {
	BaseImage   string
	ScriptsPath string
}

type CreateOptions struct {
	*CreateFlags
	*vm.VM
}

func (cf *CreateFlags) NewCreateOptions(args []string) (co *CreateOptions, err error) {
	if _, err = os.Stat(args[0]); err != nil {
		log.Errorf("os image file:%s do not exsit", args[0])
		return
	}

	virtManager := virt.NewVirtManager("")

	domxml, err := xml.GetDefaultXML()
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
		VM:          vm,
	}

	return
}

func (co *CreateOptions) CreateImage() error {
	// apply scripts under /etc/wib/scripts.d in order
	time.Sleep(time.Second * 10)
	log.Infof("ip:%s, port:%s, passwd:%s", co.VM.Ip, co.VM.Port, co.VM.Password)
	out, err := co.VM.Output("systeminfo")
	if err != nil {
		return err
	}
	log.Infof(string(out))

	// stop vm destroy vm

	// convert snapshot image

	return nil
}
