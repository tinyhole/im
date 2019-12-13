package config

import (
	microConfig "github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/etcd"
	"github.com/pkg/errors"
)

type Config struct {
	SrvName           string   `json:"srvName"`
	SrvID             int32    `json:"srvID"`
	MgoAddrs          []string `json:"mgoAddrs"`
	MgoUser           string   `json:"mgoUser"`
	MgoPassword       string   `json:"mgoPassword"`
	MgoReplicaSetName string   `json:"mgoReplicaSetName"`
	MgoPoolLimit      int      `json:"mgoPoolLimit"`
	LogFilePath       string   `json:"logFilePath"`
	LogLevel          string   `json:"logLevel"`
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
	return &conf, nil
}
