package config

import (
	microConfig "github.com/micro/go-micro/config"
)

type BaseConfig struct {
	ConfigCenterAddr   string `yaml:"ConfigCenterAddr"`
	ConfigPath         string `yaml:"ConfigPath"`
	RegistryCenterAddr string `yaml:"RegistryCenterAddr"`
	SrvName            string `yaml:"SrvName"`
}

func NewBaseConfig() (*BaseConfig, error) {
	baseConfig := &BaseConfig{}
	err := microConfig.LoadFile("config_base.yaml")
	if err != nil {
		return nil, err
	}
	err = microConfig.Scan(baseConfig)
	return baseConfig, err
}
