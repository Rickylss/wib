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
	_, err := vm.StartVm()
	if err != nil {
		return err
	}

	log.Debug("do something")

	ip, err := vm.GetIpByMac("default", "52:54:00:43:e3:36")
	if err != nil {
		return err
	}

	log.Infof("ip:%s", ip)

	//dom.Destroy()

	// connect vm by ip:port passwd

	// apply scripts under /etc/wib/scripts.d in order

	// stop vm destroy vm

	// convert snapshot image

	return nil
}
