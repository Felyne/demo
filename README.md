
### swag生成restful API文档  

- [什么是OpenAPI？Swagger是什么？](什么是OpenAPI？Swagger是什么？)
- [安装和了解swag](https://github.com/swaggo/swag)
- [YApi教程](https://hellosean1025.github.io/yapi/documents/index.html)
- [docker部署YApi](https://github.com/fjc0k/docker-YApi)

项目里面用gin框架做了示范，其他类型的handler也是支持的  
在项目根目录下执行以下命令生成docs目录
```
swag init
```

### YApi管理项目的API文档

打开`localhost:40001`登录YApi  

1. 在个人空间下创建一个`test`分组，组长填自己的用户名好了

2. 在`test`分组下创建一个项目`demo`  

3. 在`数据管理`导入上面生成的docs/swagger.json文件，即可查看和测试api
