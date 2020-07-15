# go-dc

使用Go与Gin的Demo项目，采用依赖注入的方式进行开发，已完善单元测试

## 目录结构
```

cmd
  server.go      |--启动文件
configs
  config.yaml    |--服务配置文件
docs             |--常用文档 
pkg
  cache          |--缓存
  config         |--配置读取
  router         |--路由
  rpc            |--rpc服务
  controller     |--控制器：负责数据解析、逻辑判断和返回结果
  service        |--服务层：主要是处理控制层传入的数据并进行业务理
  dao            |--数据访问层：服务层获取数据的接口包
  model          |--数据模型层：定义在数据库中存储的实体结构
  database       |--数据仓库层：数据库等存储都放在这个包下
  util           |--工具包：主要是一些小工具的包      
```
## 启动方式

```sh
go run cmd/server.go
```

## 注意事项
仅供学习，误用于生产环境

## 更新日志

### 20200504
将mongo与redis的链接过程改成并发处理，mysql暂无处理

### 20200713
增加gin解析参数发生错误时返回的提示信息，以明确语义和位置

