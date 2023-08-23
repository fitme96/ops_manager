package service

import (
	"log"

	"github.com/go-ini/ini"
)

type Server struct {
	RunMode  string
	HttpPort string
}

var ServerSetting = &Server{}

type Database struct {
	User     string
	Password string
	Host     string
	Name     string
}

var DatabaseSetting = &Database{}

type Proxmox struct {
	Token  string
	Secret string
	Url    string
}

var ProxmoxSetting = &Proxmox{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("proxmox", ProxmoxSetting)

}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
