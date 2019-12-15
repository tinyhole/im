# relation
关系服务

# 功能
1. 个人关系
 - 互相follow 就成为好友
 - 单方面取消follow 就失去好友关系
2. 群关系

# 测试
## config
# config 
```shell script
export ETCDCTL_API=3
etcdctl put /conf/mua/im/relation/ '{"MgoAddrs":["127.0.0.1:27017"],"LogFilePath":"../log/relation.log","LogLevel":"debug"}'
etcdctl get /conf/mua/im/relation --prefix
```

## mongo

```json
{"_id":0,"name":"测试群","notice":"这是群公告","create_at":1576376177,"type":1}
{"src_uid":1,"group_id":0,"join_at":1576376177,"role":3,"is_validate":1,"create_at":1576376177}
{"src_uid":2, "group_id":0,"join_at":1576376177,"role":3,"is_validate":1,"create_at":1576376177}
```
