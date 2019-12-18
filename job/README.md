# Job
消息投递服务
1. 将消息写到收件箱
2. 发送消息同步通知


# config 
```shell script
export ETCDCTL_API=3
etcdctl put /conf/mua/im/job/ '{"MgoAddrs":["127.0.0.1:27017"],"LogFilePath":"../log/job.log","LogLevel":"debug", "NSQLookupAddr":["127.0.0.1:4161"],"NSQDAddr":"127.0.0.1:4150"}'
etcdctl get /conf/mua/im/job --prefix
```
