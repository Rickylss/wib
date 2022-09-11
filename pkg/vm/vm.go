package vm

import "github.com/Rickylss/wib/pkg/image"

type VmSsh struct {
	Ip      string
	User    string
	Passwd  string
	SshPort int
	SshKey  string
}

type Vm struct {
	Name string
	*VmSsh
	image.Image
}
