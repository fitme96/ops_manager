package service

import "github.com/luthermonson/go-proxmox"

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

type Resize struct {
	Vm   Vmcyclelife
	Disk string `json:"disk"`
	Size string `json:"size"`
}

type Vmcyclelife struct {
	Node string `json:"node"`
	Vmid int    `json:"vmid"`
}

type Clonevm struct {
	Node          string                             `json:"node"`
	Vmcloneoption proxmox.VirtualMachineCloneOptions `json:"vmcloneoption"`
}

type Configvm struct {
	Vm       Vmcyclelife
	VmOption []proxmox.VirtualMachineOption
}

type Snapshotvm struct {
	Vm           Vmcyclelife
	SnapshotName string
}

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
