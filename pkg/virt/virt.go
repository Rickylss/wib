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

	net, _ := conn.LookupNetworkByName("default")
	leases, _ := net.GetDHCPLeases()
	for _, lease := range leases {
		if lease.Mac == "52:54:00:43:e3:36" {
			return lease.IPaddr, nil
		}
	}

	return "", nil
}
