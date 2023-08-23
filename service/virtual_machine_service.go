package service

import (
	"github.com/gin-gonic/gin"
	"github.com/luthermonson/go-proxmox"
)

func CloneVm(c *gin.Context, vmclone Clonevm) (newvm int, err error) {

	client := GetClient()
	node, _ := client.Node(vmclone.Node)
	vm, _ := node.VirtualMachine(100)
	newvm, _, err = vm.Clone(&vmclone.Vmcloneoption)
	if err != nil {
		return newvm, err
	}
	return newvm, nil

}

func DeleteVm(c *gin.Context, v Vmcyclelife) error {

	client := GetClient()
	node, _ := client.Node(v.Node)
	vm, _ := node.VirtualMachine(v.Vmid)
	if _, err := vm.Delete(); err != nil {
		return err
	}

	return nil
}

func ResizeDisk(c *gin.Context, resize Resize) error {

	client := GetClient()
	node, _ := client.Node(resize.Vm.Node)
	vm, _ := node.VirtualMachine(resize.Vm.Vmid)
	if err := vm.ResizeDisk(resize.Disk, resize.Size); err != nil {
		return err
	}
	return nil
}

func VmConfig(c *gin.Context, vmoption Configvm) error {
	client := GetClient()
	node, _ := client.Node(vmoption.Vm.Node)
	vm, _ := node.VirtualMachine(vmoption.Vm.Vmid)
	for _, v := range vmoption.VmOption {
		if _, err := vm.Config(v); err != nil {
			return err
		}

	}
	return nil
}
func VmStop(c *gin.Context, v Vmcyclelife) error {
	client := GetClient()
	node, _ := client.Node(v.Node)
	if vm, err := node.VirtualMachine(v.Vmid); err != nil {
		return err
	} else {
		vm.Stop()

	}

	return nil
}

func VmStart(c *gin.Context, v Vmcyclelife) error {
	client := GetClient()
	node, _ := client.Node(v.Node)
	if vm, err := node.VirtualMachine(v.Vmid); err != nil {
		return err
	} else {
		vm.Start()

	}
	return nil
}

func VmShutdown(c *gin.Context, v Vmcyclelife) error {
	client := GetClient()
	node, _ := client.Node(v.Node)
	if vm, err := node.VirtualMachine(v.Vmid); err != nil {
		return err
	} else {
		vm.Shutdown()
	}
	return nil
}

func VmReboot(c *gin.Context, v Vmcyclelife) error {
	client := GetClient()
	node, _ := client.Node(v.Node)
	if vm, err := node.VirtualMachine(v.Vmid); err != nil {
		return err
	} else {
		vm.Reboot()

	}
	return nil
}

// snapshot
func NewSnapshot(c *gin.Context, s Snapshotvm) error {
	client := GetClient()
	node, _ := client.Node(s.Vm.Node)
	vm, _ := node.VirtualMachine(s.Vm.Vmid)
	if _, err := vm.NewSnapshot(s.SnapshotName); err != nil {
		return err
	}
	return nil
}

func Snapshots(c *gin.Context, v Vmcyclelife) ([]*proxmox.Snapshot, error) {
	client := GetClient()
	node, _ := client.Node(v.Node)
	vm, _ := node.VirtualMachine(v.Vmid)
	if snapshots, err := vm.Snapshots(); err != nil {
		return nil, err
	} else {
		return snapshots, nil
	}

}

func SnapshotRollback(c *gin.Context, s Snapshotvm) error {
	client := GetClient()
	node, _ := client.Node(s.Vm.Node)
	vm, _ := node.VirtualMachine(s.Vm.Vmid)
	if _, err := vm.SnapshotRollback(s.SnapshotName); err != nil {
		return err
	}

	return nil
}
