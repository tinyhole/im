package config

import (
	"fmt"
	microConfig "github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/etcd"
	"github.com/pkg/errors"
)

type Config struct {
	SrvName           string   `json:"SrvName"`
	SrvID             int32    `json:"SrvID"`
	MgoAddrs          []string `json:"MgoAddrs"`
	MgoUser           string   `json:"MgoUser"`
	MgoPassword       string   `json:"MgoPassword"`
	MgoReplicaSetName string   `json:"MgoReplicaSetName"`
	MgoPoolLimit      int      `json:"MgoPoolLimit"`
	LogFilePath       string   `json:"LogFilePath"`
	LogLevel          string   `json:"LogLevel"`
}

func NewConfig(baseConf *BaseConfig) (*Config, error) {
	var (
		err error
	)
	fmt.Printf("%v", baseConf)
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
		fmt.Printf("microConfig error [%v]", err)
		return nil, errors.WithStack(err)
	}
	fmt.Printf("[%v]", microConfig.Get("srvName").String("aa"))
	fmt.Printf("%v", conf)
	return &conf, nil
}
