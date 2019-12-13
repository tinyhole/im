package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type BaseConfig struct {
	ConfigCenterAddr   string `yaml:"ConfigCenterAddr"`
	ConfigPath         string `yaml:"ConfigPath"`
	RegistryCenterAddr string `yaml:"RegistryCenterAddr"`
	SrvID              uint32 `yaml:"SrvID"`
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
