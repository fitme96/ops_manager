package service

import (
	"crypto/tls"
	"net/http"

	"github.com/luthermonson/go-proxmox"
)

func GetClient() *proxmox.Client {
	insecureHTTPClient := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	tokenID := ProxmoxSetting.Token
	secret := ProxmoxSetting.Secret

	vmclient := proxmox.NewClient(ProxmoxSetting.Url,
		proxmox.WithClient(&insecureHTTPClient),
		proxmox.WithAPIToken(tokenID, secret),
	)
	return vmclient
}
