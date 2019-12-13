package config

import (
	"fmt"
	microConfig "github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/etcd"
	"github.com/pkg/errors"
)

type Config struct {
	RedisAddr     string `json:"RedisAddr"`
	RedisPassword string `json:"RedisPassword"`
	NSQDAddr      string `json:"NSQDAddr"`
	LogFilePath   string `json:"LogFilePath"`
	LogLevel      string `json:"LogLevel"`
}

func NewConfig(baseConf *BaseConfig) (*Config, error) {
	var (
		err error
	)
	etcdSource := etcd.NewSource(etcd.WithAddress(baseConf.ConfigCenterAddr),
		etcd.WithPrefix(baseConf.ConfigPath),
		etcd.StripPrefix(true),
	)
	err = microConfig.Load(etcdSource)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	conf := Config{}
	err = microConfig.Scan(&conf)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	fmt.Printf("conf [%v]", conf)
	return &conf, nil
}
