# IM 服务

# 目录说明
- ap im 接入点服务组件
- idl 公共协议目录
- job im 消息同步服务组件
- logic im 消息投递.会话状态，auth 服务组件
- relation im 关系服务组件
- sequence im 消息序号服务组件


## idl 

### 工具编译

[idl-go](https://github.com/tinyhole/idl)

### 环境变量配置
```shell script
    export THIDL=/Users/apple/workbench/src/github.com/tinyhole/idl
```
### 输出目录

idl-go 工具依赖当前目录下的idl.yaml工具，输出目录为当前目录下的idl目录

确保idl-go工具在环境变量中，然后在项目根目录执行如下命令即可
```shell script
idl-go
```


