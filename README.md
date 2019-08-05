# gin-demo

使用Go与Gin的Demo项目

## 目录结构
```
cms
  server          
pkg
  router         |--路由层
  controller     |--控制器层,负责逻辑判断和返回结果
  service        |--服务层,主要是处理控制层传入的数据并进行业务理
  dao            |--数据访问层 是服务层获取数据的接口包
  model          |--数据模型层，定义在数据库中存储的实体结构
  database       |--数据仓库层,把数据库和redis和其他存储都放在这个包下
  util           |--工具包,主要是一些小工具的包      
```
## 启动方式

```sh
go build
go run cms/server/server.go
```

## 注意事项
仅供学习，误用于生产环境
