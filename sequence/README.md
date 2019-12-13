# sequence

消息序号服务

1. 每个收件箱有自己独立的序号


# config 
```shell script
export ETCDCTL_API=3
etcdctl put /conf/mua/im/sequence/ '{"SrvName":"sequence","SrvID":1,"MgoAddrs":["127.0.0.1:27017"],"LogFilePath":"../log/sequence.log","LogLevel":"debug"}'
etcdctl get /conf/mua/im/sequence --prefix
```

