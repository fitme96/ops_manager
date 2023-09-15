package service

import (
	"github.com/gin-gonic/gin"
	"github.com/luthermonson/go-proxmox"
)

// 生命周期
const (
	DELETE_VM   = 201
	STOP_VM     = 202
	START_VM    = 203
	SHUTDOWN_VM = 204
	REBOOT_VM   = 205
)

func vmaction(action uint, v Vmcyclelife) error {
	var err error
	client := GetClient()
	node, _ := client.Node(v.Node)
	vm, _ := node.VirtualMachine(v.Vmid)
	switch {
	case action == DELETE_VM:
		_, err = vm.Delete()
	case action == STOP_VM:
		_, err = vm.Stop()
	case action == START_VM:
		_, err = vm.Start()
	case action == SHUTDOWN_VM:
		_, err = vm.Shutdown()
	case action == REBOOT_VM:
		_, err = vm.Reboot()

	}
	if err != nil {
		return err
	}
	return nil
}

func DeleteVm(c *gin.Context, v Vmcyclelife) error {
	if err := vmaction(DELETE_VM, v); err != nil {
		return err
	}
	return nil
}

func VmStop(c *gin.Context, v Vmcyclelife) error {
	if err := vmaction(STOP_VM, v); err != nil {
		return err
	}
	return nil
}

func VmStart(c *gin.Context, v Vmcyclelife) error {
	if err := vmaction(START_VM, v); err != nil {
		return err
	}
	return nil
}

func VmShutdown(c *gin.Context, v Vmcyclelife) error {
	if err := vmaction(SHUTDOWN_VM, v); err != nil {
		return err
	}
	return nil
}

func VmReboot(c *gin.Context, v Vmcyclelife) error {
	if err := vmaction(REBOOT_VM, v); err != nil {
		return err
	}
	return nil
}

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

// snapshot
const (
	CREATE_SNAPSHOT   = 100
	ROLLBACK_SNAPSHOT = 101
)

func snapshotaction(action int, snap Snapshotvm) error {
	var err error
	client := GetClient()
	node, _ := client.Node(snap.Vm.Node)
	vm, _ := node.VirtualMachine(snap.Vm.Vmid)

	switch {
	case action == CREATE_SNAPSHOT:
		_, err = vm.NewSnapshot(snap.SnapshotName)

	case action == ROLLBACK_SNAPSHOT:
		_, err = vm.SnapshotRollback(snap.SnapshotName)

	}
	if err != nil {
		return err
	}

	return nil
}

func NewSnapshot(c *gin.Context, snap Snapshotvm) error {
	if err := snapshotaction(CREATE_SNAPSHOT, snap); err != nil {
		return err
	}

	return nil
}

func SnapshotRollback(c *gin.Context, snap Snapshotvm) error {
	if err := snapshotaction(ROLLBACK_SNAPSHOT, snap); err != nil {
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
