package virt

import (
	"fmt"

	"libvirt.org/go/libvirt"
)

type VirtManager struct {
	URI string
}

func NewVirtManager() *VirtManager {
	return &VirtManager{
		URI: "qemu:///system",
	}
}

func (v *VirtManager) StartVm() (dom *libvirt.Domain, err error) {
	conn, err := libvirt.NewConnect(v.URI)
	if err != nil {
		return
	}

	domcfg, err := GetDefaultXML()
	if err != nil {
		return
	}

	dom, err = conn.DomainCreateXML(domcfg, libvirt.DOMAIN_NONE)
	if err != nil {
		return
	}

	name, _ := dom.GetName()

	fmt.Printf("domain:%s is running\n", name)

	return
}

func (v *VirtManager) StopVm(dom *libvirt.Domain) error {
	return dom.Destroy()
}

func (v *VirtManager) GetIpByMac(network string, mac string) (string, error) {
	conn, err := libvirt.NewConnect(v.URI)
	if err != nil {
		return "", err
	}

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
