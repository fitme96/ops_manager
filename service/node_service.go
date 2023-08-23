package service

import (
	"github.com/gin-gonic/gin"
	"github.com/luthermonson/go-proxmox"
)

func PveVersion(c *gin.Context) (*proxmox.Version, error) {
	client := GetClient()
	if version, err := client.Version(); err != nil {
		return nil, err
	} else {
		return version, nil

	}
}

func Nodelist(c *gin.Context) ([]string, error) {
	var nodelist []string
	client := GetClient()
	if ns, err := client.Nodes(); err != nil {
		return nil, err
	} else {
		for _, v := range ns {
			nodelist = append(nodelist, v.Node)
		}
		return nodelist, nil
	}

}
func VmDetail(c *gin.Context, vm Vmcyclelife) (*proxmox.VirtualMachine, error) {

	client := GetClient()
	node, _ := client.Node(vm.Node)
	if vmdetail, err := node.VirtualMachine(vm.Vmid); err != nil {
		return nil, err
	} else {
		return vmdetail, nil
	}

}

func VmList(c *gin.Context, nodename string) ([]string, error) {
	var vmlist []string
	client := GetClient()
	node, _ := client.Node(nodename)

	if vmlistdetail, err := node.VirtualMachines(); err != nil {
		return nil, err
	} else {
		for _, v := range vmlistdetail {
			vmlist = append(vmlist, v.Name)
		}
		return vmlist, nil
	}

}
