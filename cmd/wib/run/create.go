package run

import (
	"os"

	"github.com/Rickylss/wib/pkg/constants"
	"github.com/Rickylss/wib/pkg/image"
	"github.com/Rickylss/wib/pkg/virt"

	log "github.com/sirupsen/logrus"
)

const (
	DefaultBase        = "win.img"
	DefaultSize        = "40G"
	DefaultScriptsPath = constants.WorkDir + "scripts.d"
)

type CreateFlags struct {
	BaseImage   string
	Size        string
	Release     bool
	ScriptsPath string
}

type CreateOptions struct {
	*CreateFlags
	Image image.Image
}

func (cf *CreateFlags) NewCreateOptions(args []string) (co *CreateOptions, err error) {

	finfo, err := os.Stat(args[0])
	if err != nil {
		log.Errorf("os image file:%s do not exsit", args[0])
		return
	}

	co = &CreateOptions{
		CreateFlags: cf,
		Image: &image.Qcow2Image{
			Name: finfo.Name(),
			Size: finfo.Size(),
		},
	}

	return
}

func (co *CreateOptions) CreateImage() error {

	vm := virt.NewVirtManager()
	dom, err := vm.StartVm()
	if err != nil {
		return err
	}

	log.Debug("do something")

	log.Debug("then destroy")
	dom.Destroy()
	// check image file

	// create vm from xml

	// make external snapshot for vm by libvirt

	// get vm ip address

	// connect vm by ip:port passwd

	// apply scripts under /etc/wib/scripts.d in order

	// stop vm destroy vm

	// convert snapshot image

	return nil
}
