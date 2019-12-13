# logic
IM 逻辑服务
1. 身份认证
2. 消息投递
3. 会话服务


# test 
```shell script
export ETCDCTL_API=3
etcdctl put /conf/mua/im/logic/ '{"SrvName":"mua.im.logic","SrvID":1,"RedisAddr":"127.0.0.1:6379","LogFilePath":"../log/logic.log","LogLevel":"debug","NSQDAddr":"127.0.0.1:4150"}'
etcdctl get /conf/mua/im/logic --prefix
```