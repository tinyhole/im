package config

import (
	microConfig "github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/etcd"
	"github.com/pkg/errors"
)

type Config struct {
	MgoAddrs          []string `json:"mgoAddrs"`
	MgoUser           string   `json:"mgoUser"`
	MgoPassword       string   `json:"mgoPassword"`
	MgoReplicaSetName string   `json:"mgoReplicaSetName"`
	MgoPoolLimit      int      `json:"mgoPoolLimit"`
	LogFilePath       string   `json:"LogFilePath"`
	LogLevel          string   `json:"LogLevel"`
	NSQLookUpAddr     string   `json:"NSQLookUpAddr"`
	NSQDAddr          string   `json:"NSQDAddr"`
	NSQChannel        string   `json:"NSQChannel"`
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
