package run

import (
	"os"

	"github.com/Rickylss/wib/pkg/image"

	log "github.com/sirupsen/logrus"
)

const (
	DefaultOutPut = "win.img"
	DefaultSize   = "40G"
)

type CreateFlags struct {
	OutPut  string
	Size    string
	Release bool
}

type CreateOptions struct {
	*CreateFlags
	BaseImage *image.BaseImage
}

func (cf *CreateFlags) NewCreateOptions(args []string) (co *CreateOptions, err error) {

	finfo, err := os.Stat(args[0])
	if err != nil {
		log.Errorf("os image file:%s do not exsit", args[0])
		return
	}

	co = &CreateOptions{
		CreateFlags: cf,
		BaseImage: &image.BaseImage{
			Name: finfo.Name(),
			Size: finfo.Size(),
		},
	}

	return
}

func (co *CreateOptions) CreateImage() error {
	return nil
}
