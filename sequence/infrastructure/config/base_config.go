package config

import (
	"github.com/ghodss/yaml"
	"io/ioutil"
)

type BaseConfig struct {
	ConfigCenterAddr   string `yaml:"ConfigCenterAddr"`
	ConfigPath         string `yaml:"ConfigPath"`
	RegistryCenterAddr string `yaml:"RegistryCenterAddr"`
	SrvName            string `yaml:"SrvName"`
}

func NewBaseConfig() (*BaseConfig, error) {
	baseConfig := &BaseConfig{}
	confData, err := ioutil.ReadFile("config_base.yaml")
	if err != nil {
		return nil, err
	}
	yaml.Unmarshal(confData, baseConfig)
	return baseConfig, err
}
