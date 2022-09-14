package config

import (
	"log"

	"taego/lib/util"
	c "github.com/olebedev/config"
)

var (
	Config *c.Config
)

func init() {
	cfg, err := c.ParseYamlFile("etc/config.yaml")
	if err != nil {
		log.Fatalf("parse config from etc/config.yaml failed: %s", err)
	}

	mode := util.GetMode()
	if Config, err = cfg.Get(string(mode)); err != nil {
		log.Fatalf("get config from Config, mode=%s: %s", mode, err)
	}
}

func OpentraceSwitch() bool {
	return Config.UBool("opentrace.switch", true)
}
