package config

import (
	microConfig "github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/etcd"
	"github.com/pkg/errors"
)

type Config struct {
	MgoAddrs          []string `json:"MgoAddrs"`
	MgoUser           string   `json:"MgoUser"`
	MgoPassword       string   `json:"MgoPassword"`
	MgoReplicaSetName string   `json:"MgoReplicaSetName"`
	MgoPoolLimit      int      `json:"MgoPoolLimit"`
	LogFilePath       string   `json:"LogFilePath"`
	LogLevel          string   `json:"LogLevel"`
	NSQLookupdAddr    []string `json:"NSQLookUpAddr"`
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
