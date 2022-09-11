package virt

import (
	"github.com/google/uuid"
	"github.com/lithammer/dedent"
	log "github.com/sirupsen/logrus"
	vx "libvirt.org/libvirt-go-xml"
)

const (
	MemorySz = 4194304
	VmName   = "win10-test"
	VcpuN    = 4
	Emulator = "/usr/bin/qemu-system-x86_64"
)

func GetDefaultXML() (xmldoc string, err error) {
	domcfg := &vx.Domain{
		Type:  "kvm",
		Name:  VmName,
		UUID:  uuid.New().String(),
		Title: "wib",
		Metadata: &vx.DomainMetadata{
			XML: dedent.Dedent(`
            <libosinfo:libosinfo xmlns:libosinfo="http://libosinfo.org/xmlns/libvirt/domain/1.0">
              <libosinfo:os id="http://microsoft.com/win/10"/>
            </libosinfo:libosinfo>
		  `),
		},
		Memory: &vx.DomainMemory{
			Value: MemorySz,
			Unit:  "KiB",
		},
		CurrentMemory: &vx.DomainCurrentMemory{
			Value: MemorySz,
			Unit:  "KiB",
		},
		VCPU: &vx.DomainVCPU{
			Value:     VcpuN,
			Placement: "static",
		},
		Resource: &vx.DomainResource{
			Partition: "/machine",
		},
		OS: &vx.DomainOS{
			Type: &vx.DomainOSType{
				Type:    "hvm",
				Arch:    "x86_64",
				Machine: "q35",
			},
			BootDevices: []vx.DomainBootDevice{
				{Dev: "hd"},
			},
		},
		Features: DefaultFeatureListXML(),
		CPU: &vx.DomainCPU{
			Mode:       "host-passthrough",
			Check:      "none",
			Migratable: "on",
		},
		Clock:      DefaultClockXML(),
		OnPoweroff: "destroy",
		OnReboot:   "restart",
		OnCrash:    "destroy",
		PM: &vx.DomainPM{
			SuspendToMem: &vx.DomainPMPolicy{
				Enabled: "no",
			},
			SuspendToDisk: &vx.DomainPMPolicy{
				Enabled: "no",
			},
		},
		Devices: DefaultDeviceListXML(),
	}

	xmldoc, err = domcfg.Marshal()
	if err != nil {
		return
	}
	log.Debug(xmldoc)
	return
}

func DefaultFeatureListXML() *vx.DomainFeatureList {
	return &vx.DomainFeatureList{
		ACPI: &vx.DomainFeature{},
		APIC: &vx.DomainFeatureAPIC{},
		HyperV: &vx.DomainFeatureHyperV{
			Relaxed: &vx.DomainFeatureState{
				State: "on",
			},
			VAPIC: &vx.DomainFeatureState{
				State: "on",
			},
			Spinlocks: &vx.DomainFeatureHyperVSpinlocks{
				DomainFeatureState: vx.DomainFeatureState{
					State: "on",
				},
				Retries: 8191,
			},
		},
		VMPort: &vx.DomainFeatureState{State: "off"},
	}
}

func DefaultClockXML() *vx.DomainClock {
	return &vx.DomainClock{
		Offset: "localtime",
		Timer: []vx.DomainTimer{
			{
				Name:       "rtc",
				TickPolicy: "catchup",
			},
			{
				Name:       "pit",
				TickPolicy: "delay",
			},
			{
				Name:    "hpet",
				Present: "no",
			},
			{
				Name:    "hypervclock",
				Present: "yes",
			},
		},
	}
}

func DefaultDeviceListXML() *vx.DomainDeviceList {
	return &vx.DomainDeviceList{
		Emulator:    Emulator,
		Disks:       DefaultDiskListXML(),
		Controllers: DefaultControllerXML(),
		Interfaces:  DefaultInterfaceXML(),
		Serials:     DefaultSerialXML(),
		Consoles:    DefaultConsoleXML(),
		Channels:    DefaultChannelXML(),
		Inputs:      DefaultInputXML(),
		Graphics:    DefaultGraphicsXML(),
		Sounds:      DefaultSoundXML(),
		Audios:      DefaultAudioXML(),
		Videos:      DefaultVideoXML(),
		RedirDevs:   DefaultRedirdevXML(),
		MemBalloon:  MemBalloonXML(),
	}
}

func DefaultDiskListXML() []vx.DomainDisk {
	var controller uint = 0
	var bus uint = 0
	var target uint = 0
	var unit uint = 0
	return []vx.DomainDisk{
		{
			Device: "disk",
			Driver: &vx.DomainDiskDriver{
				Name:    "qemu",
				Type:    "qcow2",
				Discard: "unmap",
			},
			Source: &vx.DomainDiskSource{
				File: &vx.DomainDiskSourceFile{
					File: "/home/rickylss/repo/win10.qcow2",
				},
			},
			BackingStore: &vx.DomainDiskBackingStore{},
			Target: &vx.DomainDiskTarget{
				Dev: "sda",
				Bus: "sata",
			},
			Address: &vx.DomainAddress{
				Drive: &vx.DomainAddressDrive{
					Controller: &controller,
					Bus:        &bus,
					Target:     &target,
					Unit:       &unit,
				},
			},
		},
	}
}

func DefaultControllerXML() []vx.DomainController {
	var usbIndex uint = 0
	var usbPort uint = 15
	var usbDomain uint = 0x0000
	var usbBus uint = 0x02
	var usbSlot uint = 0x00
	var usbFunc uint = 0x0

	var pciRootIndex uint = 0
	var pciIndex uint = 1
	var pciCharssis uint = 1
	var pciPort uint = 0x10
	var pciDomain uint = 0x0000
	var pciBus uint = 0x00
	var pciSlot uint = 0x02
	var pciFunc uint = 0x0

	var sataIndex uint = 0
	var sataDomain uint = 0x0000
	var sataBus uint = 0x00
	var sataSlot uint = 0x1f
	var sataFunc uint = 0x2

	var vsIndex uint = 0
	var vsDomain uint = 0x0000
	var vsBus uint = 0x03
	var vsSlot uint = 0x00
	var vsFunc uint = 0x0
	return []vx.DomainController{
		{
			Type:  "usb",
			Index: &usbIndex,
			Model: "qemu-xhci",
			USB: &vx.DomainControllerUSB{
				Port: &usbPort,
			},
			Address: &vx.DomainAddress{
				PCI: &vx.DomainAddressPCI{
					Domain:   &usbDomain,
					Bus:      &usbBus,
					Slot:     &usbSlot,
					Function: &usbFunc,
				},
			},
		},
		{
			Type:  "pci",
			Index: &pciRootIndex,
			Model: "pcie-root",
		},
		{
			Type:  "pci",
			Index: &pciIndex,
			Model: "pcie-root-port",
			PCI: &vx.DomainControllerPCI{
				Model: &vx.DomainControllerPCIModel{Name: "pcie-root-port"},
				Target: &vx.DomainControllerPCITarget{
					Chassis: &pciCharssis,
					Port:    &pciPort,
				},
			},
			Address: &vx.DomainAddress{
				PCI: &vx.DomainAddressPCI{
					Domain:        &pciDomain,
					Bus:           &pciBus,
					Slot:          &pciSlot,
					Function:      &pciFunc,
					MultiFunction: "on",
				},
			},
		},
		{
			Type:  "sata",
			Index: &sataIndex,
			Address: &vx.DomainAddress{
				PCI: &vx.DomainAddressPCI{
					Domain:   &sataDomain,
					Bus:      &sataBus,
					Slot:     &sataSlot,
					Function: &sataFunc,
				},
			},
		},
		{
			Type:  "virtio-serial",
			Index: &vsIndex,
			Address: &vx.DomainAddress{
				PCI: &vx.DomainAddressPCI{
					Domain:   &vsDomain,
					Bus:      &vsBus,
					Slot:     &vsSlot,
					Function: &vsFunc,
				},
			},
		},
	}
}

func DefaultInterfaceXML() []vx.DomainInterface {
	var domain uint = 0x0000
	var bus uint = 0x01
	var slot uint = 0x00
	var function uint = 0x0
	return []vx.DomainInterface{
		{
			MAC: &vx.DomainInterfaceMAC{
				Address: "52:54:00:43:e3:36",
			},
			Source: &vx.DomainInterfaceSource{
				Network: &vx.DomainInterfaceSourceNetwork{
					Network: "default",
				},
			},
			Model: &vx.DomainInterfaceModel{
				Type: "e1000e",
			},
			Address: &vx.DomainAddress{
				PCI: &vx.DomainAddressPCI{
					Domain:   &domain,
					Bus:      &bus,
					Slot:     &slot,
					Function: &function,
				},
			},
		},
	}
}

func DefaultSerialXML() []vx.DomainSerial {
	var port uint = 0
	return []vx.DomainSerial{
		{
			Target: &vx.DomainSerialTarget{
				Type: "isa-serial",
				Port: &port,
				Model: &vx.DomainSerialTargetModel{
					Name: "isa-serial",
				},
			},
		},
	}
}

func DefaultConsoleXML() []vx.DomainConsole {
	var port uint = 0
	return []vx.DomainConsole{
		{
			Target: &vx.DomainConsoleTarget{
				Type: "serial",
				Port: &port,
			},
		},
	}
}

func DefaultChannelXML() []vx.DomainChannel {
	var controller uint = 0
	var bus uint = 0
	var port uint = 1
	return []vx.DomainChannel{
		{
			Target: &vx.DomainChannelTarget{
				VirtIO: &vx.DomainChannelTargetVirtIO{
					Name: "com.redhat.spice.0",
				},
			},
			Address: &vx.DomainAddress{
				VirtioSerial: &vx.DomainAddressVirtioSerial{
					Controller: &controller,
					Bus:        &bus,
					Port:       &port,
				},
			},
		},
	}
}

func DefaultInputXML() []vx.DomainInput {
	var bus uint = 0
	return []vx.DomainInput{
		{
			Type: "tablet",
			Bus:  "usb",
			Address: &vx.DomainAddress{
				USB: &vx.DomainAddressUSB{
					Bus:  &bus,
					Port: "1",
				},
			},
		},
		{
			Type: "mouse",
			Bus:  "ps2",
		},
		{
			Type: "keyboard",
			Bus:  "ps2",
		},
	}
}

func DefaultGraphicsXML() []vx.DomainGraphic {
	return []vx.DomainGraphic{
		{
			Spice: &vx.DomainGraphicSpice{
				AutoPort: "yes",
				Listen:   "127.0.0.1",
				Listeners: []vx.DomainGraphicListener{
					{
						Address: &vx.DomainGraphicListenerAddress{
							Address: "127.0.0.1",
						},
					},
				},
				Image: &vx.DomainGraphicSpiceImage{
					Compression: "off",
				},
			},
		},
	}
}

func DefaultSoundXML() []vx.DomainSound {
	var domain uint = 0x0000
	var bus uint = 0x00
	var slot uint = 0x1b
	var function uint = 0x0
	return []vx.DomainSound{
		{
			Model: "ich9",
			Address: &vx.DomainAddress{
				PCI: &vx.DomainAddressPCI{
					Domain:   &domain,
					Bus:      &bus,
					Slot:     &slot,
					Function: &function,
				},
			},
		},
	}
}

func DefaultAudioXML() []vx.DomainAudio {
	return []vx.DomainAudio{
		{
			ID:    1,
			SPICE: &vx.DomainAudioSPICE{},
		},
	}
}

func DefaultVideoXML() []vx.DomainVideo {
	var domain uint = 0x0000
	var bus uint = 0x00
	var slot uint = 0x01
	var function uint = 0x0
	return []vx.DomainVideo{
		{
			Model: vx.DomainVideoModel{
				Type:    "qxl",
				Ram:     65536,
				VRam:    65536,
				VGAMem:  16384,
				Heads:   1,
				Primary: "yes",
			},
			Address: &vx.DomainAddress{
				PCI: &vx.DomainAddressPCI{
					Domain:   &domain,
					Bus:      &bus,
					Slot:     &slot,
					Function: &function,
				},
			},
		},
	}
}

func DefaultRedirdevXML() []vx.DomainRedirDev {
	var bus uint = 0
	return []vx.DomainRedirDev{
		{
			Bus: "usb",
			Address: &vx.DomainAddress{
				USB: &vx.DomainAddressUSB{
					Bus:  &bus,
					Port: "2",
				},
			},
			Source: &vx.DomainChardevSource{
				SpiceVMC: &vx.DomainChardevSourceSpiceVMC{},
			},
		},
		{
			Bus: "usb",
			Address: &vx.DomainAddress{
				USB: &vx.DomainAddressUSB{
					Bus:  &bus,
					Port: "3",
				},
			},
			Source: &vx.DomainChardevSource{
				SpiceVMC: &vx.DomainChardevSourceSpiceVMC{},
			},
		},
	}
}

func MemBalloonXML() *vx.DomainMemBalloon {
	var domain uint = 0x0000
	var bus uint = 0x04
	var slot uint = 0x00
	var function uint = 0x0
	return &vx.DomainMemBalloon{
		Model: "virtio",
		Address: &vx.DomainAddress{
			PCI: &vx.DomainAddressPCI{
				Domain:   &domain,
				Bus:      &bus,
				Slot:     &slot,
				Function: &function,
			},
		},
	}
}
