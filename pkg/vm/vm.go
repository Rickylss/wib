package vm

import "libvirt.org/go/libvirt"

type VM struct {
	Name string
	*SshContext
	*Image
	*libvirt.Domain
}

func (vm *VM) Run(cmd string) error {
	session, err := vm.SshContext.GetSession([]string{})
	if err != nil {
		return err
	}
	defer session.Close()

	return session.Run(cmd)
}

func (vm *VM) Output(cmd string) ([]byte, error) {
	session, err := vm.SshContext.GetSession([]string{})
	if err != nil {
		return nil, err
	}
	defer session.Close()

	return session.Output(cmd)
}

func (vm *VM) GetImageSize() error {
	_, err := vm.Domain.GetBlockInfo(vm.Image.Path, 0)
	if err != nil {
		return err
	}

	return nil
}
