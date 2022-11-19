package virt

import (
	"path"

	"github.com/Rickylss/wib/pkg/vm"
	"libvirt.org/go/libvirt"
	vx "libvirt.org/libvirt-go-xml"
)

type VirtManager struct {
	URI string
}

func NewVirtManager(uri string) *VirtManager {
	if uri == "" {
		uri = "qemu:///system"
	}

	return &VirtManager{
		URI: uri,
	}
}

func (v *VirtManager) StartVm(domxml *vx.Domain) (m *vm.VM, err error) {
	conn, err := libvirt.NewConnect(v.URI)
	if err != nil {
		return
	}
	defer conn.Close()

	xmldoc, err := domxml.Marshal()
	if err != nil {
		return
	}

	dom, err := conn.DomainCreateXML(xmldoc, libvirt.DOMAIN_NONE)
	if err != nil {
		return
	}

	name, err := dom.GetName()
	if err != nil {
		return
	}

	firstDisk := domxml.Devices.Disks[0]
	blkInfo, err := dom.GetBlockInfo(firstDisk.Source.File.File, 0)
	if err != nil {
		return
	}

	firstInterface := domxml.Devices.Interfaces[0]
	ip, err := v.GetInterfaceIp(firstInterface)
	if err != nil {
		return
	}

	m = &vm.VM{
		Name:   name,
		Domain: dom,
		Image: &vm.Image{
			Name: path.Base(firstDisk.Source.File.File),
			Path: firstDisk.Source.File.File,
			Size: int(blkInfo.Physical),
			Type: vm.ImageType(firstDisk.Driver.Type),
		},
		SshContext: &vm.SshContext{
			Ip: ip,
		},
	}

	return
}

func (v *VirtManager) StopVm(dom *libvirt.Domain) error {
	return dom.Destroy()
}

func (v *VirtManager) GetInterfaceIp(firstInterface vx.DomainInterface) (string, error) {
	network := firstInterface.Source.Network.Network
	mac := firstInterface.MAC.Address
	conn, err := libvirt.NewConnect(v.URI)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	net, err := conn.LookupNetworkByName(network)
	if err != nil {
		return "", err
	}
	leases, err := net.GetDHCPLeases()
	if err != nil {
		return "", err
	}
	for _, lease := range leases {
		if lease.Mac == mac {
			return lease.IPaddr, nil
		}
	}

	return "", nil
}
