# Wings
As the operation and maintenance personnel to control and monitor servers, switches and other equipment control
自动注册

# register etcd

rpc 注册服务 服务自带一个api接口在值里面，提供监控数据采集
api 单独部署的服务 自带一个监控数据采集接口和访问接口

# 说明

该包只提供模块注册和api注册功能 其余功能请引包使用

# 系统

## 监控

### 系统信息|配置
用ansible实时获取

### 自定义监控
用golang+http api实现，自己定制化

## 管理

### 依赖环境部署
ansible-playbook + role进行管理和部署 