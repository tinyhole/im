# ap
接入点服务

## 功能说明：
提供TCP 长连接的接入服务


## 接口

## test setting

# test 
```shell script
export ETCDCTL_API=3
etcdctl put /conf/mua/im/ap/ '{"ApPort":8080,"LogFilePath":"../log/ap.log","LogLevel":"debug"}'
etcdctl get /conf/mua/im/ap --prefix
````